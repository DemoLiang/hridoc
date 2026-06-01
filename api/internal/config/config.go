package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type MinIO struct {
	Endpoint        string
	AccessKey       string
	SecretKey       string
	Bucket          string
	UseSSL          bool
	PresignedExpiry int
}

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	CacheRedis redis.RedisConf
	JwtAuth    struct {
		AccessSecret string
		AccessExpire int64
	}
	MinIO MinIO
	Watermark struct {
		FontPath    string
		DefaultText string
	}
	Upload struct {
		MaxFileSize  int64
		AllowedTypes []string
	}
	Excel struct {
		MaxRows int
	}
	Export struct {
		TempDir     string
		CleanupDays int
	}
}
