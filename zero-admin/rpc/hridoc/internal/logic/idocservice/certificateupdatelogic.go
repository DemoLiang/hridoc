package idocservicelogic

import (
	"context"
	"zero-admin/rpc/proto/idoc"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCertificateUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateUpdateLogic {
	return &CertificateUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CertificateUpdate 证书更新
func (l *CertificateUpdateLogic) CertificateUpdate(in *idoc.CertificateUpdateReq) (*idoc.CertificateUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &idoc.CertificateUpdateResp{}, nil
}
