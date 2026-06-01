// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
