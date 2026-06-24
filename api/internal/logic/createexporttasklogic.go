// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"archive/zip"
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/model"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/DemoLiang/hridoc/api/pkg/watermark"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateExportTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateExportTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExportTaskLogic {
	return &CreateExportTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateExportTaskLogic) CreateExportTask(req *types.ExportTaskReq, excelBytes []byte) (resp *types.ExportTaskResp, err error) {
	if len(req.CategoryCodes) == 0 {
		return &types.ExportTaskResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidParam, Message: "参数错误"},
		}, nil
	}

	task := &model.ExportTask{
		TaskName:  sql.NullString{String: fmt.Sprintf("导出任务_%s", time.Now().Format("20060102150405")), Valid: true},
		Status:    1,
		CreatedBy: 0,
		CreatedAt: time.Now(),
	}

	result, err := l.svcCtx.ExportTaskModel.Insert(l.ctx, task)
	if err != nil {
		logx.Errorf("insert export task failed: %v", err)
		return &types.ExportTaskResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "创建任务失败"},
		}, nil
	}

	taskId, _ := result.LastInsertId()

	go l.processExport(taskId, excelBytes, req.UserIds, req)

	return &types.ExportTaskResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data:     types.ExportTaskData{TaskId: taskId},
	}, nil
}

func (l *CreateExportTaskLogic) processExport(taskId int64, excelBytes []byte, userIds []int64, req *types.ExportTaskReq) {
	ctx := context.Background()

	var matchedUsers []types.ExcelPreviewUser
	var certInfos map[int64][]certInfo
	var missCount int
	var parseErr error

	if len(excelBytes) > 0 {
		matchedUsers, certInfos, missCount, parseErr = l.parseExcelAndMatchUsers(ctx, excelBytes, req.CategoryCodes)
	} else if len(userIds) > 0 {
		matchedUsers, certInfos, missCount, parseErr = l.buildUsersFromIds(ctx, userIds, req.CategoryCodes)
	} else {
		parseErr = fmt.Errorf("缺少名单数据")
	}

	if parseErr != nil {
		l.updateTaskFailed(ctx, taskId, parseErr.Error())
		return
	}

	if len(matchedUsers) == 0 {
		l.updateTaskFailed(ctx, taskId, "未找到可导出的证件文件")
		return
	}

	// 2. 创建临时目录
	tempDir := fmt.Sprintf("./tmp/exports/task_%d", taskId)
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		logx.Errorf("create temp dir failed: %v", err)
		l.updateTaskFailed(ctx, taskId, "创建临时目录失败")
		return
	}
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()

	// 3. 下载证件文件、打水印、保存到临时目录
	var totalCertCount int
	for _, user := range matchedUsers {
		userCerts := certInfos[user.UserId]
		if len(userCerts) == 0 {
			continue
		}

		userDir := filepath.Join(tempDir, sanitizeFileName(user.UserName))
		if err := os.MkdirAll(userDir, 0755); err != nil {
			logx.Errorf("create user dir failed: %v", err)
			continue
		}

		for _, cert := range userCerts {
			if cert.FileUrl == "" {
				continue
			}

			// 从 MinIO 下载文件
			objectName := extractMinioObjectName(cert.FileUrl, l.svcCtx.Config.MinIO.Bucket)
				if objectName == "" {
					logx.Errorf("cannot extract object name from fileUrl=%s", cert.FileUrl)
					continue
				}
				obj, err := l.svcCtx.MinIO.GetObject(ctx, objectName)
			if err != nil {
				logx.Errorf("get object from minio failed, fileUrl=%s: %v", cert.FileUrl, err)
				continue
			}

			data, err := io.ReadAll(obj)
			obj.Close()
			if err != nil {
				logx.Errorf("read object data failed, fileUrl=%s: %v", cert.FileUrl, err)
				continue
			}

			// 打水印
			opts := &watermark.Options{
				Text:     req.WatermarkText,
				Mode:     req.WatermarkMode,
				Color:    req.WatermarkColor,
				Opacity:  req.WatermarkOpacity,
				FontSize: req.WatermarkFontSize,
			}
			processedData, fileType, err := watermark.ApplyIfNeeded(data, cert.FileType, opts)
			if err != nil {
				logx.Errorf("apply watermark failed, certId=%d: %v", cert.CertId, err)
				processedData = data
				fileType = cert.FileType
			}

			// 确定文件扩展名
			ext := getFileExt(fileType)
			fileName := fmt.Sprintf("%s-%s%s", sanitizeFileName(cert.CatName), sanitizeFileName(cert.CertName), ext)
			filePath := filepath.Join(userDir, fileName)

			if err := os.WriteFile(filePath, processedData, 0644); err != nil {
				logx.Errorf("write file failed: %v", err)
				continue
			}

			totalCertCount++
		}
	}

	if totalCertCount == 0 {
		l.updateTaskFailed(ctx, taskId, "未找到可导出的证件文件")
		return
	}

	// 4. 打包 ZIP
	zipBuf := &bytes.Buffer{}
	zipWriter := zip.NewWriter(zipBuf)

	err := filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(tempDir, path)
		if err != nil {
			return err
		}
		relPath = filepath.ToSlash(relPath)

		fileWriter, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		fileData, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		_, err = fileWriter.Write(fileData)
		return err
	})
	if err != nil {
		logx.Errorf("walk temp dir failed: %v", err)
		l.updateTaskFailed(ctx, taskId, "打包 ZIP 失败")
		return
	}

	if err := zipWriter.Close(); err != nil {
		logx.Errorf("close zip writer failed: %v", err)
		l.updateTaskFailed(ctx, taskId, "打包 ZIP 失败")
		return
	}

	// 5. 上传 ZIP 到 MinIO
	objectName := fmt.Sprintf("exports/task_%d_%s.zip", taskId, time.Now().Format("20060102150405"))
	err = l.svcCtx.MinIO.PutObject(ctx, objectName, bytes.NewReader(zipBuf.Bytes()), int64(zipBuf.Len()), "application/zip")
	if err != nil {
		logx.Errorf("upload zip to minio failed: %v", err)
		l.updateTaskFailed(ctx, taskId, "上传导出文件失败")
		return
	}

	// 6. 更新任务状态为完成
	url := l.svcCtx.MinIO.ObjectURL(objectName)
	now := time.Now()
	_, err = l.svcCtx.DB.ExecCtx(ctx,
		`update export_task set task_name=?, user_count=?, cert_count=?, miss_count=?, file_url=?, status=?, completed_at=? where id=?`,
		fmt.Sprintf("导出任务_%s", time.Now().Format("20060102150405")),
		int64(len(matchedUsers)), int64(totalCertCount), int64(missCount), url, 2, now, taskId,
	)
	if err != nil {
		logx.Errorf("update export task success status failed: %v", err)
	}
}

func (l *CreateExportTaskLogic) updateTaskFailed(ctx context.Context, taskId int64, reason string) {
	_, err := l.svcCtx.DB.ExecCtx(ctx,
		`update export_task set status=?, fail_reason=? where id=?`,
		3, reason, taskId,
	)
	if err != nil {
		logx.Errorf("update export task failed status failed: %v", err)
	}
}

// parseExcelAndMatchUsers 解析 Excel 并匹配用户，返回匹配到的用户列表、证件信息、missCount
func (l *CreateExportTaskLogic) parseExcelAndMatchUsers(ctx context.Context, excelBytes []byte, categoryCodes []string) ([]types.ExcelPreviewUser, map[int64][]certInfo, int, error) {
	f, err := excelize.OpenReader(strings.NewReader(string(excelBytes)))
	if err != nil {
		return nil, nil, 0, fmt.Errorf("Excel 文件格式错误")
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		sheetName = "Sheet1"
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("Excel 文件读取失败")
	}

	if len(rows) < 2 {
		return nil, nil, 0, fmt.Errorf("Excel 数据为空")
	}

	// 查询证件类型信息
	catPlaceholders := make([]string, len(categoryCodes))
	catArgs := make([]any, len(categoryCodes))
	for i, code := range categoryCodes {
		catPlaceholders[i] = "?"
		catArgs[i] = code
	}

	catQuery := fmt.Sprintf("select id, code, name from cert_category where code in (%s) and status = 1", strings.Join(catPlaceholders, ","))
	var catRows []struct {
		Id   int64  `db:"id"`
		Code string `db:"code"`
		Name string `db:"name"`
	}
	if err := l.svcCtx.DB.QueryRowsPartialCtx(ctx, &catRows, catQuery, catArgs...); err != nil {
		return nil, nil, 0, fmt.Errorf("查询证件类型失败")
	}

	catCodeMap := make(map[string]struct{ Id int64; Name string }, len(catRows))
	for _, c := range catRows {
		catCodeMap[c.Code] = struct{ Id int64; Name string }{Id: c.Id, Name: c.Name}
	}

	var matchedUsers []types.ExcelPreviewUser
	certInfos := make(map[int64][]certInfo)
	var missCount int

	for rowIdx := 1; rowIdx < len(rows); rowIdx++ {
		row := rows[rowIdx]
		if len(row) < 2 {
			continue
		}

		name := strings.TrimSpace(row[0])
		idCard := strings.TrimSpace(row[1])

		if name == "" && idCard == "" {
			continue
		}

		// 步骤 1: 尝试身份证号精确匹配
		var matchedUser *model.User
		if idCard != "" {
			user, err := l.svcCtx.UserModel.FindOneByIdCard(ctx, idCard)
			if err == nil && user != nil && user.Status == 1 {
				matchedUser = user
			}
		}

		// 步骤 2: 身份证号未匹配，尝试姓名匹配
		if matchedUser == nil && name != "" {
			var nameUsers []struct {
				Id   int64  `db:"id"`
				Name string `db:"name"`
			}
			queryErr := l.svcCtx.DB.QueryRowsPartialCtx(ctx, &nameUsers,
				"select id, name from user where name = ? and status = 1", name)
			if queryErr == nil && len(nameUsers) == 1 {
				matchedUser = &model.User{
					Id:   nameUsers[0].Id,
					Name: nameUsers[0].Name,
				}
			}
		}

		if matchedUser == nil {
			continue
		}

		// 步骤 3: 查询该用户的证件（需要 file_url 和 file_type）
		certCatPlaceholders := make([]string, len(categoryCodes))
		certCatArgs := make([]any, len(categoryCodes))
		for i, code := range categoryCodes {
			certCatPlaceholders[i] = "?"
			certCatArgs[i] = code
		}
		certCatArgs = append(certCatArgs, matchedUser.Id)

		certQuery := fmt.Sprintf(`select c.id, c.user_id, c.name as cert_name, c.file_url, c.file_type, cc.code, cc.name as cat_name
			from certificate c join cert_category cc ON c.category_id = cc.id
			where cc.code in (%s) and c.user_id = ? and c.status = 1
			order by c.user_id, cc.code, c.id desc`, strings.Join(certCatPlaceholders, ","))

		var certRows []struct {
			Id       int64  `db:"id"`
			UserId   int64  `db:"user_id"`
			CertName string `db:"cert_name"`
			FileUrl  string `db:"file_url"`
			FileType string `db:"file_type"`
			Code     string `db:"code"`
			CatName  string `db:"cat_name"`
		}
		if err := l.svcCtx.DB.QueryRowsPartialCtx(ctx, &certRows, certQuery, certCatArgs...); err != nil {
			logx.Errorf("query certs failed: %v", err)
		}

		// 去重：每个 categoryCode 只取第一个
		seenCodes := make(map[string]bool)
		var userCerts []certInfo
		for _, r := range certRows {
			if seenCodes[r.Code] {
				continue
			}
			seenCodes[r.Code] = true
			userCerts = append(userCerts, certInfo{
				CertId:   r.Id,
				CertName: r.CertName,
				FileUrl:  r.FileUrl,
				FileType: r.FileType,
				CatName:  r.CatName,
			})
		}

		if len(userCerts) == 0 {
			missCount++
			continue
		}

		matchedUsers = append(matchedUsers, types.ExcelPreviewUser{
			UserId:   matchedUser.Id,
			UserName: matchedUser.Name,
		})
		certInfos[matchedUser.Id] = userCerts
	}

	return matchedUsers, certInfos, missCount, nil
}

// buildUsersFromIds 根据用户ID直接查询用户和证件信息
func (l *CreateExportTaskLogic) buildUsersFromIds(ctx context.Context, userIds []int64, categoryCodes []string) ([]types.ExcelPreviewUser, map[int64][]certInfo, int, error) {
	if len(categoryCodes) == 0 {
		return nil, nil, 0, fmt.Errorf("请选择证件类型")
	}

	placeholders := make([]string, len(userIds))
	args := make([]any, len(userIds))
	for i, id := range userIds {
		placeholders[i] = "?"
		args[i] = id
	}

	var users []struct {
		Id   int64  `db:"id"`
		Name string `db:"name"`
	}
	query := fmt.Sprintf("select id, name from user where id in (%s) and status = 1", strings.Join(placeholders, ","))
	if err := l.svcCtx.DB.QueryRowsPartialCtx(ctx, &users, query, args...); err != nil {
		return nil, nil, 0, fmt.Errorf("查询用户失败")
	}

	catPlaceholders := make([]string, len(categoryCodes))
	catArgs := make([]any, len(categoryCodes))
	for i, code := range categoryCodes {
		catPlaceholders[i] = "?"
		catArgs[i] = code
	}
	var catRows []struct {
		Id   int64  `db:"id"`
		Code string `db:"code"`
		Name string `db:"name"`
	}
	catQuery := fmt.Sprintf("select id, code, name from cert_category where code in (%s) and status = 1", strings.Join(catPlaceholders, ","))
	if err := l.svcCtx.DB.QueryRowsPartialCtx(ctx, &catRows, catQuery, catArgs...); err != nil {
		return nil, nil, 0, fmt.Errorf("查询证件类型失败")
	}
	catCodeMap := make(map[string]struct{ Id int64; Name string }, len(catRows))
	for _, c := range catRows {
		catCodeMap[c.Code] = struct{ Id int64; Name string }{Id: c.Id, Name: c.Name}
	}

	var matchedUsers []types.ExcelPreviewUser
	certInfos := make(map[int64][]certInfo)
	var missCount int

	for _, u := range users {
		certCatPlaceholders := make([]string, len(categoryCodes))
		certCatArgs := make([]any, len(categoryCodes))
		for i, code := range categoryCodes {
			certCatPlaceholders[i] = "?"
			certCatArgs[i] = code
		}
		certCatArgs = append(certCatArgs, u.Id)

		certQuery := fmt.Sprintf(`select c.id, c.user_id, c.name as cert_name, c.file_url, c.file_type, cc.code, cc.name as cat_name
			from certificate c join cert_category cc ON c.category_id = cc.id
			where cc.code in (%s) and c.user_id = ? and c.status = 1
			order by c.user_id, cc.code, c.id desc`, strings.Join(certCatPlaceholders, ","))

		var certRows []struct {
			Id       int64  `db:"id"`
			UserId   int64  `db:"user_id"`
			CertName string `db:"cert_name"`
			FileUrl  string `db:"file_url"`
			FileType string `db:"file_type"`
			Code     string `db:"code"`
			CatName  string `db:"cat_name"`
		}
		if err := l.svcCtx.DB.QueryRowsPartialCtx(ctx, &certRows, certQuery, certCatArgs...); err != nil {
			logx.Errorf("query certs failed: %v", err)
		}

		seenCodes := make(map[string]bool)
		var userCerts []certInfo
		for _, cr := range certRows {
			if seenCodes[cr.Code] {
				continue
			}
			seenCodes[cr.Code] = true
			userCerts = append(userCerts, certInfo{
				CertId:   cr.Id,
				CertName: cr.CertName,
				FileUrl:  cr.FileUrl,
				FileType: cr.FileType,
				CatName:  cr.CatName,
			})
		}

		if len(userCerts) == 0 {
			missCount++
			continue
		}

		matchedUsers = append(matchedUsers, types.ExcelPreviewUser{
			UserId:   u.Id,
			UserName: u.Name,
		})
		certInfos[u.Id] = userCerts
	}

	return matchedUsers, certInfos, missCount, nil
}

type certInfo struct {
	CertId   int64
	CertName string
	FileUrl  string
	FileType string
	CatName  string
}

func extractMinioObjectName(fileUrl, bucket string) string {
	parsed := strings.Split(fileUrl, "?")[0]
	parts := strings.Split(parsed, "/")
	if len(parts) < 2 {
		return ""
	}
	// Try exact bucket match
	bucketIndex := -1
	for i := 0; i < len(parts); i++ {
		if parts[i] == bucket {
			bucketIndex = i
			break
		}
	}
	if bucketIndex != -1 && bucketIndex+1 < len(parts) {
		return strings.Join(parts[bucketIndex+1:], "/")
	}
	// Fallback: last two path segments
	return parts[len(parts)-2] + "/" + parts[len(parts)-1]
}

func sanitizeFileName(name string) string {
	// 移除文件名中的非法字符
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
	)
	return replacer.Replace(name)
}

func getFileExt(fileType string) string {
	switch fileType {
	case "jpeg", "jpg":
		return ".jpg"
	case "png":
		return ".png"
	case "pdf":
		return ".pdf"
	case "image":
		return ".png"
	default:
		return ""
	}
}
