package idocservicelogic

import (
	"context"
	"zero-admin/rpc/proto/idoc"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateCategoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCertificateCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateCategoryUpdateLogic {
	return &CertificateCategoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CertificateCategoryUpdate 证书类别更新
func (l *CertificateCategoryUpdateLogic) CertificateCategoryUpdate(in *idoc.CertificateCategoryUpdateReq) (*idoc.CertificateCategoryUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &idoc.CertificateCategoryUpdateResp{}, nil
}
