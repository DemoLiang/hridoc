package hridoc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CertificateModel = (*customCertificateModel)(nil)

type (
	// CertificateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCertificateModel.
	CertificateModel interface {
		certificateModel
	}

	customCertificateModel struct {
		*defaultCertificateModel
	}
)

// NewCertificateModel returns a model for the database table.
func NewCertificateModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CertificateModel {
	return &customCertificateModel{
		defaultCertificateModel: newCertificateModel(conn, c, opts...),
	}
}
