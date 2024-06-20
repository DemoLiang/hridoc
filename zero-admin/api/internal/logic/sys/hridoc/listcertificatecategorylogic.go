package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCertificateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCertificateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCertificateCategoryLogic {
	return &ListCertificateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCertificateCategoryLogic) ListCertificateCategory(req *types.ListCertificateCategoryReq) (resp *types.ListCertificateCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
