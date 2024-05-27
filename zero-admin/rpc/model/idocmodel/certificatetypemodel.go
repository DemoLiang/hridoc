package hridoc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CertificateTypeModel = (*customCertificateTypeModel)(nil)

type (
	// CertificateTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCertificateTypeModel.
	CertificateTypeModel interface {
		certificateTypeModel
	}

	customCertificateTypeModel struct {
		*defaultCertificateTypeModel
	}
)

// NewCertificateTypeModel returns a model for the database table.
func NewCertificateTypeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CertificateTypeModel {
	return &customCertificateTypeModel{
		defaultCertificateTypeModel: newCertificateTypeModel(conn, c, opts...),
	}
}
