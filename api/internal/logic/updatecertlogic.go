// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCertLogic {
	return &UpdateCertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCertLogic) UpdateCert(req *types.UpdateCertReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
