// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/model"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type PreviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreviewLogic {
	return &PreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PreviewLogic) Preview(req *types.PreviewReq, excelBytes []byte) (resp *types.ExcelPreviewResp, err error) {
	if len(req.CategoryCodes) == 0 {
		return &types.ExcelPreviewResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidParam, Message: "请选择证件类型"},
		}, nil
	}

	// 解析 Excel
	f, err := excelize.OpenReader(strings.NewReader(string(excelBytes)))
	if err != nil {
		logx.Errorf("open excel failed: %v", err)
		return &types.ExcelPreviewResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidParam, Message: "Excel 文件格式错误"},
		}, nil
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		sheetName = "Sheet1"
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		logx.Errorf("get excel rows failed: %v", err)
		return &types.ExcelPreviewResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidParam, Message: "Excel 文件读取失败"},
		}, nil
	}

	if len(rows) < 2 {
		return &types.ExcelPreviewResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidParam, Message: "Excel 数据为空，至少需要一行表头和一行数据"},
		}, nil
	}

	// 查询证件类型信息
	catPlaceholders := make([]string, len(req.CategoryCodes))
	catArgs := make([]any, len(req.CategoryCodes))
	for i, code := range req.CategoryCodes {
		catPlaceholders[i] = "?"
		catArgs[i] = code
	}

	catQuery := fmt.Sprintf("select id, code, name from cert_category where code in (%s) and status = 1", strings.Join(catPlaceholders, ","))
	var catRows []struct {
		Id   int64  `db:"id"`
		Code string `db:"code"`
		Name string `db:"name"`
	}
	if err = l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &catRows, catQuery, catArgs...); err != nil {
		logx.Errorf("query categories failed: %v", err)
		return &types.ExcelPreviewResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	catCodeMap := make(map[string]struct{ Id int64; Name string }, len(catRows))
	for _, c := range catRows {
		catCodeMap[c.Code] = struct{ Id int64; Name string }{Id: c.Id, Name: c.Name}
	}

	// 处理每一行数据
	var users []types.ExcelPreviewUser
	var matchedCount, missCount, unmatchedCount int

	for rowIdx := 1; rowIdx < len(rows); rowIdx++ {
		row := rows[rowIdx]
		if len(row) < 2 {
			continue
		}

		name := strings.TrimSpace(row[0])
		idCard := strings.TrimSpace(row[1])

		// 跳过空行
		if name == "" && idCard == "" {
			continue
		}

		previewUser := types.ExcelPreviewUser{
			RowIndex: rowIdx + 1, // Excel 行号从 2 开始
			Name:     name,
			IdCard:   idCard,
		}

		// 步骤 1: 尝试身份证号精确匹配
		var matchedUser *model.User
		var matchStatus int // 0-未匹配 1-身份证号匹配 2-姓名匹配 3-重名无法匹配

		if idCard != "" {
			user, err := l.svcCtx.UserModel.FindOneByIdCard(l.ctx, idCard)
			if err == nil && user != nil && user.Status == 1 {
				matchedUser = user
				matchStatus = 1
			} else if err != nil && err != model.ErrNotFound {
				logx.Errorf("FindOneByIdCard failed: %v", err)
			}
		}

		// 步骤 2: 身份证号未匹配，尝试姓名匹配
		if matchedUser == nil && name != "" {
			var nameUsers []struct {
				Id   int64  `db:"id"`
				Name string `db:"name"`
			}
			queryErr := l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &nameUsers,
				"select id, name from user where name = ? and status = 1", name)
			if queryErr != nil {
				logx.Errorf("query user by name failed: %v", queryErr)
			}

			if len(nameUsers) == 0 {
				matchStatus = 0
				previewUser.MissReason = "未找到匹配人员"
			} else if len(nameUsers) == 1 {
				matchedUser = &model.User{
					Id:   nameUsers[0].Id,
					Name: nameUsers[0].Name,
				}
				matchStatus = 2
			} else {
				matchStatus = 3
				previewUser.MissReason = "存在重名人员，请补充身份证号"
			}
		}

		previewUser.MatchStatus = matchStatus

		if matchedUser == nil {
			unmatchedCount++
			users = append(users, previewUser)
			continue
		}

		previewUser.UserId = matchedUser.Id
		previewUser.UserName = matchedUser.Name

		// 步骤 3: 查询该用户的证件
		certCatPlaceholders := make([]string, len(req.CategoryCodes))
		certCatArgs := make([]any, len(req.CategoryCodes))
		for i, code := range req.CategoryCodes {
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
		if err = l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &certRows, certQuery, certCatArgs...); err != nil {
			logx.Errorf("query certs failed: %v", err)
		}

		// 构建每个 categoryCode 的匹配情况
		certMap := make(map[string]struct {
			CertId   int64
			CertName string
			FileUrl  string
			FileType string
			CatName  string
		})
		for _, r := range certRows {
			if _, exists := certMap[r.Code]; !exists {
				certMap[r.Code] = struct {
					CertId   int64
					CertName string
					FileUrl  string
					FileType string
					CatName  string
				}{
					CertId:   r.Id,
					CertName: r.CertName,
					FileUrl:  r.FileUrl,
					FileType: r.FileType,
					CatName:  r.CatName,
				}
			}
		}

		hasAnyCert := false
		for _, code := range req.CategoryCodes {
			catInfo, ok := catCodeMap[code]
			if !ok {
				continue
			}
			pc := types.PreviewCategory{
				CategoryCode: code,
				CategoryName: catInfo.Name,
			}
			if cert, ok := certMap[code]; ok {
				pc.HasCert = true
				pc.CertId = cert.CertId
				pc.CertName = cert.CertName
				hasAnyCert = true
			}
			previewUser.Categories = append(previewUser.Categories, pc)
		}

		if hasAnyCert {
			matchedCount++
		} else {
			missCount++
			previewUser.MissReason = "匹配到人员但未找到指定类型证件"
		}

		users = append(users, previewUser)
	}

	return &types.ExcelPreviewResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.ExcelPreviewData{
			TotalCount:     len(users),
			MatchedCount:   matchedCount,
			MissCount:      missCount,
			UnmatchedCount: unmatchedCount,
			Users:          users,
		},
	}, nil
}
