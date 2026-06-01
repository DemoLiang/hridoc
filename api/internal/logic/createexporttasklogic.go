// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
