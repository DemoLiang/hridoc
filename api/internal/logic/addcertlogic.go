// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCertLogic {
	return &AddCertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCertLogic) AddCert(req *types.AddCertReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
