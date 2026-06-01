// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCertListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCertListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCertListLogic {
	return &GetCertListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCertListLogic) GetCertList(req *types.CertListReq) (resp *types.CertListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
