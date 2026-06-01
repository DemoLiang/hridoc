// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/model"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
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

func (l *CreateExportTaskLogic) CreateExportTask(req *types.ExportTaskReq) (resp *types.ExportTaskResp, err error) {
	if len(req.UserIds) == 0 || len(req.CategoryCodes) == 0 {
		return &types.ExportTaskResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidParam, Message: "参数错误"},
		}, nil
	}

	task := &model.ExportTask{
		TaskName: sql.NullString{String: fmt.Sprintf("导出任务_%s", time.Now().Format("20060102150405")), Valid: true},
		Status:   1,
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

	go l.processExport(taskId, req)

	return &types.ExportTaskResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data:     types.ExportTaskData{TaskId: taskId},
	}, nil
}

func (l *CreateExportTaskLogic) processExport(taskId int64, req *types.ExportTaskReq) {
	ctx := context.Background()

	users, certCount, missCount, err := l.buildPreviewData(ctx, req)
	if err != nil {
		l.updateTaskFailed(ctx, taskId, err.Error())
		return
	}

	f := excelize.NewFile()
	sheet := "导出结果"
	_ = f.SetSheetName("Sheet1", sheet)

	headers := []string{"姓名", "身份证号"}
	for _, code := range req.CategoryCodes {
		name := l.getCategoryName(ctx, code)
		headers = append(headers, name)
	}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		_ = f.SetCellValue(sheet, cell, h)
	}

	for rIdx, u := range users {
		rowNum := rIdx + 2
		_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", rowNum), u.UserName)
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", rowNum), l.getUserIdCard(ctx, u.UserId))

		for cIdx, pc := range u.Categories {
			col := string(rune('C' + cIdx))
			if pc.HasCert {
				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, rowNum), pc.CertName)
			} else {
				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, rowNum), "缺证")
			}
		}
	}

	for i := range headers {
		col := string(rune('A' + i))
		_ = f.SetColWidth(sheet, col, col, 18)
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		l.updateTaskFailed(ctx, taskId, "生成Excel失败")
		return
	}

	objectName := fmt.Sprintf("exports/task_%d_%s.xlsx", taskId, time.Now().Format("20060102150405"))
	err = l.svcCtx.MinIO.PutObject(ctx, objectName, bytes.NewReader(buf.Bytes()), int64(buf.Len()),
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	if err != nil {
		logx.Errorf("upload export file failed: %v", err)
		l.updateTaskFailed(ctx, taskId, "上传导出文件失败")
		return
	}

	url := l.svcCtx.MinIO.ObjectURL(objectName)
	now := time.Now()
	_, err = l.svcCtx.DB.ExecCtx(ctx,
		`update export_task set task_name=?, user_count=?, cert_count=?, miss_count=?, file_url=?, status=?, completed_at=? where id=?`,
		fmt.Sprintf("导出任务_%s", time.Now().Format("20060102150405")),
		int64(len(users)), certCount, missCount, url, 2, now, taskId,
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

func (l *CreateExportTaskLogic) buildPreviewData(ctx context.Context, req *types.ExportTaskReq) ([]types.PreviewUser, int64, int64, error) {
	userPlaceholders := make([]string, len(req.UserIds))
	userArgs := make([]any, len(req.UserIds))
	for i, id := range req.UserIds {
		userPlaceholders[i] = "?"
		userArgs[i] = id
	}

	catPlaceholders := make([]string, len(req.CategoryCodes))
	catArgs := make([]any, len(req.CategoryCodes))
	for i, code := range req.CategoryCodes {
		catPlaceholders[i] = "?"
		catArgs[i] = code
	}

	userQuery := fmt.Sprintf("select id, name from user where id in (%s) and status = 1", strings.Join(userPlaceholders, ","))
	var userRows []struct {
		Id   int64  `db:"id"`
		Name string `db:"name"`
	}
	if err := l.svcCtx.DB.QueryRowsPartialCtx(ctx, &userRows, userQuery, userArgs...); err != nil {
		return nil, 0, 0, err
	}

	userMap := make(map[int64]string, len(userRows))
	for _, u := range userRows {
		userMap[u.Id] = u.Name
	}

	catQuery := fmt.Sprintf("select id, code, name from cert_category where code in (%s) and status = 1", strings.Join(catPlaceholders, ","))
	var catRows []struct {
		Id   int64  `db:"id"`
		Code string `db:"code"`
		Name string `db:"name"`
	}
	if err := l.svcCtx.DB.QueryRowsPartialCtx(ctx, &catRows, catQuery, catArgs...); err != nil {
		return nil, 0, 0, err
	}

	catIdMap := make(map[int64]struct{ Code, Name string }, len(catRows))
	catCodeMap := make(map[string]struct{ Id int64; Name string }, len(catRows))
	for _, c := range catRows {
		catIdMap[c.Id] = struct{ Code, Name string }{Code: c.Code, Name: c.Name}
		catCodeMap[c.Code] = struct{ Id int64; Name string }{Id: c.Id, Name: c.Name}
	}

	certPlaceholders := make([]string, len(req.UserIds))
	certArgs := make([]any, len(req.UserIds))
	for i, id := range req.UserIds {
		certPlaceholders[i] = "?"
		certArgs[i] = id
	}
	certQuery := fmt.Sprintf(`select c.id, c.user_id, c.category_id, c.name as cert_name
		from certificate c
		where c.user_id in (%s) and c.status = 1
		order by c.user_id, c.category_id, c.id desc`, strings.Join(certPlaceholders, ","))
	var certRows []struct {
		Id         int64  `db:"id"`
		UserId     int64  `db:"user_id"`
		CategoryId int64  `db:"category_id"`
		CertName   string `db:"cert_name"`
	}
	if err := l.svcCtx.DB.QueryRowsPartialCtx(ctx, &certRows, certQuery, certArgs...); err != nil {
		return nil, 0, 0, err
	}

	certMap := make(map[int64]map[string]struct {
		CertId   int64
		CertName string
	})
	for _, r := range certRows {
		if _, ok := certMap[r.UserId]; !ok {
			certMap[r.UserId] = make(map[string]struct{ CertId int64; CertName string })
		}
		if catInfo, ok := catIdMap[r.CategoryId]; ok {
			if _, exists := certMap[r.UserId][catInfo.Code]; !exists {
				certMap[r.UserId][catInfo.Code] = struct {
					CertId   int64
					CertName string
				}{CertId: r.Id, CertName: r.CertName}
			}
		}
	}

	var users []types.PreviewUser
	var certCount, missCount int64
	for _, uid := range req.UserIds {
		name, ok := userMap[uid]
		if !ok {
			continue
		}
		pu := types.PreviewUser{
			UserId:   uid,
			UserName: name,
		}
		for _, code := range req.CategoryCodes {
			catInfo, ok := catCodeMap[code]
			if !ok {
				continue
			}
			pc := types.PreviewCategory{
				CategoryCode: code,
				CategoryName: catInfo.Name,
			}
			if certs, ok := certMap[uid]; ok {
				if c, ok := certs[code]; ok {
					pc.HasCert = true
					pc.CertId = c.CertId
					pc.CertName = c.CertName
					certCount++
				} else {
					missCount++
				}
			} else {
				missCount++
			}
			pu.Categories = append(pu.Categories, pc)
		}
		users = append(users, pu)
	}

	return users, certCount, missCount, nil
}

func (l *CreateExportTaskLogic) getCategoryName(ctx context.Context, code string) string {
	var row struct {
		Name string `db:"name"`
	}
	_ = l.svcCtx.DB.QueryRowCtx(ctx, &row, "select name from cert_category where code = ? limit 1", code)
	return row.Name
}

func (l *CreateExportTaskLogic) getUserIdCard(ctx context.Context, userId int64) string {
	var row struct {
		IdCard string `db:"id_card"`
	}
	_ = l.svcCtx.DB.QueryRowCtx(ctx, &row, "select id_card from user where id = ? limit 1", userId)
	return row.IdCard
}
