package idocservicelogic

import (
	"context"
	"zero-admin/rpc/proto/idoc"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCertificateAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateAddLogic {
	return &CertificateAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CertificateAdd 证书添加
func (l *CertificateAddLogic) CertificateAdd(in *idoc.CertificateAddReq) (*idoc.CertificateAddResp, error) {
	// todo: add your logic here and delete this line

	return &idoc.CertificateAddResp{}, nil
}
