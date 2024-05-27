package idocservicelogic

import (
	"context"
	"zero-admin/rpc/proto/idoc"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateCategoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCertificateCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateCategoryAddLogic {
	return &CertificateCategoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CertificateCategoryAdd 证书类别增加
func (l *CertificateCategoryAddLogic) CertificateCategoryAdd(in *idoc.CertificateCategoryReq) (*idoc.CertificateCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &idoc.CertificateCategoryResp{}, nil
}
