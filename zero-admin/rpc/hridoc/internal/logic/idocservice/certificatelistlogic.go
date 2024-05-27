package idocservicelogic

import (
	"context"
	"zero-admin/rpc/proto/idoc"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCertificateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateListLogic {
	return &CertificateListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CertificateList 证书列表
func (l *CertificateListLogic) CertificateList(in *idoc.CertificateListReq) (*idoc.CertificateListResp, error) {
	// todo: add your logic here and delete this line

	return &idoc.CertificateListResp{}, nil
}
