package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCertificateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCertificateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCertificateLogic {
	return &ListCertificateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCertificateLogic) ListCertificate(req *types.ListCertificateReq) (resp *types.ListCertificateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
