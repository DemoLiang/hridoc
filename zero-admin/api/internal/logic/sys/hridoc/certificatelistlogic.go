package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCertificateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateListLogic {
	return &CertificateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CertificateListLogic) CertificateList(req *types.ListCertificateReq) (resp *types.ListCertificateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
