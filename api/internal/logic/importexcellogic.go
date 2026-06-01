// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
