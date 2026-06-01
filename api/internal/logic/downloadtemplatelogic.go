// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadTemplateLogic {
	return &DownloadTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadTemplateLogic) DownloadTemplate() (resp *types.TemplateResp, err error) {
	f := excelize.NewFile()
	sheet := "导入模板"
	_ = f.SetSheetName("Sheet1", sheet)

	headers := []string{"姓名", "身份证号", "证件类型(可选)", "证件等级(可选)"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		_ = f.SetCellValue(sheet, cell, h)
	}

	_ = f.SetColWidth(sheet, "A", "A", 15)
	_ = f.SetColWidth(sheet, "B", "B", 22)
	_ = f.SetColWidth(sheet, "C", "C", 22)
	_ = f.SetColWidth(sheet, "D", "D", 15)

	_ = f.SetCellValue(sheet, "A2", "张三")
	_ = f.SetCellValue(sheet, "B2", "11010119900101xxxx")
	_ = f.SetCellValue(sheet, "C2", "软考证书")
	_ = f.SetCellValue(sheet, "D2", "高级")

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		logx.Errorf("write excel template failed: %v", err)
		return &types.TemplateResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "生成模板失败"},
		}, nil
	}

	timestamp := time.Now().Format("20060102150405")
	objectName := fmt.Sprintf("templates/import_template_%s.xlsx", timestamp)
	err = l.svcCtx.MinIO.PutObject(l.ctx, objectName, bytes.NewReader(buf.Bytes()), int64(buf.Len()),
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	if err != nil {
		logx.Errorf("upload template to minio failed: %v", err)
		return &types.TemplateResp{
			BaseResp: types.BaseResp{Code: errorx.ErrUploadFailed, Message: "上传模板失败"},
		}, nil
	}

	url := l.svcCtx.MinIO.ObjectURL(objectName)
	return &types.TemplateResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data:     types.TemplateData{Url: url},
	}, nil
}
