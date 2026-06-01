package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CertCategoryModel = (*customCertCategoryModel)(nil)

type (
	// CertCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCertCategoryModel.
	CertCategoryModel interface {
		certCategoryModel
	}

	customCertCategoryModel struct {
		*defaultCertCategoryModel
	}
)

// NewCertCategoryModel returns a model for the database table.
func NewCertCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CertCategoryModel {
	return &customCertCategoryModel{
		defaultCertCategoryModel: newCertCategoryModel(conn, c, opts...),
	}
}
