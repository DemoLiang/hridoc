package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateCategoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCertificateCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateCategoryDeleteLogic {
	return &CertificateCategoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CertificateCategoryDeleteLogic) CertificateCategoryDelete(req *types.DeleteCertificateCategoryReq) (resp *types.DeleteCertificateCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
