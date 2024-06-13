package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCertificateUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateUpdateLogic {
	return &CertificateUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CertificateUpdateLogic) CertificateUpdate(req *types.UpdateCertificateReq) (resp *types.UpdateCertificateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
