// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/model"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type ImportExcelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportExcelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportExcelLogic {
	return &ImportExcelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportExcelLogic) ImportExcel(req *types.ExcelImportReq) (resp *types.ExcelImportResp, err error) {
	if len(req.File) == 0 {
		return &types.ExcelImportResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidParam, Message: "文件不能为空"},
		}, nil
	}

	f, err := excelize.OpenReader(bytes.NewReader(req.File))
	if err != nil {
		logx.Errorf("open excel failed: %v", err)
		return &types.ExcelImportResp{
			BaseResp: types.BaseResp{Code: errorx.ErrExcelFormat, Message: "无法解析 Excel 文件"},
		}, nil
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		logx.Errorf("read excel rows failed: %v", err)
		return &types.ExcelImportResp{
			BaseResp: types.BaseResp{Code: errorx.ErrExcelFormat, Message: "读取 Excel 内容失败"},
		}, nil
	}

	if len(rows) < 2 {
		return &types.ExcelImportResp{
			BaseResp: types.BaseResp{Code: errorx.ErrExcelFormat, Message: "Excel 文件内容为空"},
		}, nil
	}

	maxRows := l.svcCtx.Config.Excel.MaxRows
	if maxRows > 0 && len(rows)-1 > maxRows {
		return &types.ExcelImportResp{
			BaseResp: types.BaseResp{Code: errorx.ErrExcelValidation, Message: fmt.Sprintf("数据行数超过上限 %d", maxRows)},
		}, nil
	}

	var success, failed int64
	var failList []types.ExcelFailItem

	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 2 {
			failed++
			failList = append(failList, types.ExcelFailItem{
				Row:    int64(i + 1),
				Reason: "数据列不足，至少需要姓名和身份证号",
				Data:   strings.Join(row, ","),
			})
			continue
		}

		name := strings.TrimSpace(row[0])
		idCard := strings.TrimSpace(row[1])
		if name == "" || idCard == "" {
			failed++
			failList = append(failList, types.ExcelFailItem{
				Row:    int64(i + 1),
				Reason: "姓名或身份证号不能为空",
				Data:   strings.Join(row, ","),
			})
			continue
		}

		existing, err := l.svcCtx.UserModel.FindOneByIdCard(l.ctx, idCard)
		if err == nil && existing != nil {
			existing.Name = name
			if len(row) > 2 && strings.TrimSpace(row[2]) != "" {
			}
			if err := l.svcCtx.UserModel.Update(l.ctx, existing); err != nil {
				logx.Errorf("update user failed: %v", err)
				failed++
				failList = append(failList, types.ExcelFailItem{
					Row:    int64(i + 1),
					Reason: "更新用户信息失败",
					Data:   strings.Join(row, ","),
				})
				continue
			}
			success++
			continue
		}

		user := &model.User{
			Name:     name,
			IdCard:   idCard,
			Role:     2,
			Status:   1,
			Password: "$2a$10$N9qo8uLOickgx2ZMRZoMy.MqrqU6L1LRCXy7JLZP6JvVYfP1RRKZK",
		}
		if len(row) > 3 && strings.TrimSpace(row[3]) != "" {
			user.Education = sql.NullString{String: strings.TrimSpace(row[3]), Valid: true}
		}
		_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
		if err != nil {
			logx.Errorf("insert user failed: %v", err)
			failed++
			failList = append(failList, types.ExcelFailItem{
				Row:    int64(i + 1),
				Reason: "创建用户失败: " + err.Error(),
				Data:   strings.Join(row, ","),
			})
			continue
		}
		success++
	}

	return &types.ExcelImportResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.ExcelImportData{
			Total:    int64(len(rows) - 1),
			Success:  success,
			Failed:   failed,
			FailList: failList,
		},
	}, nil
}
