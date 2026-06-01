// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCertLogic {
	return &DeleteCertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCertLogic) DeleteCert(req *types.DeleteReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
