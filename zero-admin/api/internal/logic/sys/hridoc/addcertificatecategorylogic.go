package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCertificateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCertificateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCertificateCategoryLogic {
	return &AddCertificateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCertificateCategoryLogic) AddCertificateCategory(req *types.AddCertificateCategoryReq) (resp *types.AddCertificateCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
