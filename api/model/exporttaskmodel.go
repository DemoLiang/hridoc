package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ExportTaskModel = (*customExportTaskModel)(nil)

type (
	// ExportTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customExportTaskModel.
	ExportTaskModel interface {
		exportTaskModel
	}

	customExportTaskModel struct {
		*defaultExportTaskModel
	}
)

// NewExportTaskModel returns a model for the database table.
func NewExportTaskModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ExportTaskModel {
	return &customExportTaskModel{
		defaultExportTaskModel: newExportTaskModel(conn, c, opts...),
	}
}
