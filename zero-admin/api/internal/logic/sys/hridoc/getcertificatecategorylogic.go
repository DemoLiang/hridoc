package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCertificateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCertificateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCertificateCategoryLogic {
	return &GetCertificateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCertificateCategoryLogic) GetCertificateCategory(req *types.GetCertificateCategoryReq) (resp *types.GetCertificateCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
