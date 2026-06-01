package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/DemoLiang/hridoc/api/internal/config"
	"github.com/DemoLiang/hridoc/api/model"
	"github.com/DemoLiang/hridoc/api/pkg/minio"
)

type ServiceContext struct {
	Config            config.Config
	DB                sqlx.SqlConn
	MinIO             *minio.Client
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
	minioClient, err := minio.NewClient(c.MinIO)
	if err != nil {
		// MinIO 连接失败不应阻止服务启动，记录错误但继续
		// 实际生产代码中可以改为 panic 或更严格的处理
		minioClient = nil
	}

	return &ServiceContext{
		Config:            c,
		DB:                db,
		MinIO:             minioClient,
		UserModel:         model.NewUserModel(db, cacheConf),
		CertCategoryModel: model.NewCertCategoryModel(db, cacheConf),
		CertificateModel:  model.NewCertificateModel(db, cacheConf),
		ExportTaskModel:   model.NewExportTaskModel(db, cacheConf),
		OperationLogModel: model.NewOperationLogModel(db, cacheConf),
	}
}
