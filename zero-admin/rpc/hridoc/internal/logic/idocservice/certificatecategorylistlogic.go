package idocservicelogic

import (
	"context"
	"zero-admin/rpc/proto/idoc"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CertificateCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCertificateCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CertificateCategoryListLogic {
	return &CertificateCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CertificateCategoryList 证书类别列表
func (l *CertificateCategoryListLogic) CertificateCategoryList(in *idoc.CertificateCategoryListReq) (*idoc.CertificateCategoryListResp, error) {
	// todo: add your logic here and delete this line

	return &idoc.CertificateCategoryListResp{}, nil
}
