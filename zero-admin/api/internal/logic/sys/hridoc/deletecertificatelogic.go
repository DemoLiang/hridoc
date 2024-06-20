package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCertificateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCertificateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCertificateLogic {
	return &DeleteCertificateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCertificateLogic) DeleteCertificate(req *types.DeleteCertificateReq) (resp *types.DeleteCertificateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
