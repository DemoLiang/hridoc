package hridoc

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCertificateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCertificateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCertificateLogic {
	return &UpdateCertificateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCertificateLogic) UpdateCertificate(req *types.UpdateCertificateReq) (resp *types.UpdateCertificateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
