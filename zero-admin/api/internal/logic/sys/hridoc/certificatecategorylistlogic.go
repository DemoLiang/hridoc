package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCertificateCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateCategoryListLogic {
	return &CertificateCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CertificateCategoryListLogic) CertificateCategoryList(req *types.ListCertificateCategoryReq) (resp *types.ListCertificateCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
