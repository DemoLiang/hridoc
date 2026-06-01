package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/DemoLiang/hridoc/api/internal/config"
	"github.com/DemoLiang/hridoc/api/model"
)

type ServiceContext struct {
	Config            config.Config
	DB                sqlx.SqlConn
	UserModel         model.UserModel
	CertCategoryModel model.CertCategoryModel
	CertificateModel  model.CertificateModel
	ExportTaskModel   model.ExportTaskModel
	OperationLogModel model.OperationLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMysql(c.Mysql.DataSource)
	cacheConf := cache.CacheConf{
		{
			RedisConf: c.CacheRedis,
			Weight:    100,
		},
	}
	return &ServiceContext{
		Config:            c,
		DB:                db,
		UserModel:         model.NewUserModel(db, cacheConf),
		CertCategoryModel: model.NewCertCategoryModel(db, cacheConf),
		CertificateModel:  model.NewCertificateModel(db, cacheConf),
		ExportTaskModel:   model.NewExportTaskModel(db, cacheConf),
		OperationLogModel: model.NewOperationLogModel(db, cacheConf),
	}
}
