// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskDownloadLogic {
	return &GetTaskDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskDownloadLogic) GetTaskDownload(req *types.TaskStatusReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
