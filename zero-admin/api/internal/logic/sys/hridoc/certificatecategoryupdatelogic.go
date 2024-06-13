package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateCategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCertificateCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateCategoryUpdateLogic {
	return &CertificateCategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CertificateCategoryUpdateLogic) CertificateCategoryUpdate(req *types.UpdateCertificateCategoryReq) (resp *types.UpdateCertificateCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
