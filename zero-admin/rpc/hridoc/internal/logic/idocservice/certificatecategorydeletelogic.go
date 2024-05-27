package idocservicelogic

import (
	"context"
	"zero-admin/rpc/proto/idoc"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateCategoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCertificateCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateCategoryDeleteLogic {
	return &CertificateCategoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CertificateCategoryDelete 证书类别删除
func (l *CertificateCategoryDeleteLogic) CertificateCategoryDelete(in *idoc.CertificateCategoryDeleteReq) (*idoc.CertificateCategoryDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &idoc.CertificateCategoryDeleteResp{}, nil
}
