package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCertificateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCertificateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCertificateCategoryLogic {
	return &DeleteCertificateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCertificateCategoryLogic) DeleteCertificateCategory(req *types.DeleteCertificateCategoryReq) (resp *types.DeleteCertificateCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
