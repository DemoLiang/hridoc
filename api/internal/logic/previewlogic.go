// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"

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

func (l *PreviewLogic) Preview(req *types.PreviewReq) (resp *types.PreviewResp, err error) {
	// todo: add your logic here and delete this line

	return
}
