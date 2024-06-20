package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCertificateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCertificateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCertificateLogic {
	return &AddCertificateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCertificateLogic) AddCertificate(req *types.AddCertificateReq) (resp *types.AddCertificateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
