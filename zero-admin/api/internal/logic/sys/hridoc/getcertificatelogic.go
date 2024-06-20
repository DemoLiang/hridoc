package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCertificateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCertificateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCertificateLogic {
	return &GetCertificateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCertificateLogic) GetCertificate(req *types.GetCertificateReq) (resp *types.GetCertificateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
