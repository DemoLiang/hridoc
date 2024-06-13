package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAddCertificateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAddCertificateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddCertificateLogic {
	return &AddAddCertificateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAddCertificateLogic) AddAddCertificate(req *types.AddCertificateReq) (resp *types.AddCertificateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
