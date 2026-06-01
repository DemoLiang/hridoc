// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCertDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCertDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCertDetailLogic {
	return &GetCertDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCertDetailLogic) GetCertDetail(req *types.CertDetailReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
