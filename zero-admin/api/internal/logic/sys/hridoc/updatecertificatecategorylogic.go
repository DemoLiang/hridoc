package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCertificateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCertificateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCertificateCategoryLogic {
	return &UpdateCertificateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCertificateCategoryLogic) UpdateCertificateCategory(req *types.UpdateCertificateCategoryReq) (resp *types.UpdateCertificateCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
