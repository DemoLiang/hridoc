package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCertificateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateDeleteLogic {
	return &CertificateDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CertificateDeleteLogic) CertificateDelete(req *types.DeleteCertificateReq) (resp *types.DeleteCertificateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
