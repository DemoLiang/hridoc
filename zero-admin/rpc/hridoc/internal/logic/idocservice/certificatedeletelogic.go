package idocservicelogic

import (
	"context"
	"zero-admin/rpc/proto/idoc"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCertificateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateDeleteLogic {
	return &CertificateDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CertificateDelete 证书删除
func (l *CertificateDeleteLogic) CertificateDelete(in *idoc.CertificateDeleteReq) (*idoc.CertificateDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &idoc.CertificateDeleteResp{}, nil
}
