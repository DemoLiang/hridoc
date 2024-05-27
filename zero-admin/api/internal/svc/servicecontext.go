package svc

import (
	"zero-admin/api/internal/config"
	"zero-admin/api/internal/middleware"
	"zero-admin/rpc/hridoc/client/configservice"
	"zero-admin/rpc/hridoc/client/deptservice"
	"zero-admin/rpc/hridoc/client/dictservice"
	"zero-admin/rpc/hridoc/client/jobservice"
	"zero-admin/rpc/hridoc/client/loginlogservice"
	"zero-admin/rpc/hridoc/client/menuservice"
	"zero-admin/rpc/hridoc/client/roleservice"
	"zero-admin/rpc/hridoc/client/syslogservice"
	"zero-admin/rpc/hridoc/client/userservice"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	AddLog   rest.Middleware

	//系统相关
	ConfigService   configservice.ConfigService
	DeptService     deptservice.DeptService
	DictService     dictservice.DictService
	JobService      jobservice.JobService
	LoginLogService loginlogservice.LoginLogService
	SysLogService   syslogservice.SysLogService
	MenuService     menuservice.MenuService
	RoleService     roleservice.RoleService
	UserService     userservice.UserService

	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	sysClient := zrpc.MustNewClient(c.SysRpc)
	logService := syslogservice.NewSysLogService(sysClient)
	return &ServiceContext{
		Config: c,

		ConfigService:   configservice.NewConfigService(sysClient),
		DeptService:     deptservice.NewDeptService(sysClient),
		DictService:     dictservice.NewDictService(sysClient),
		JobService:      jobservice.NewJobService(sysClient),
		LoginLogService: loginlogservice.NewLoginLogService(sysClient),
		SysLogService:   logService,
		MenuService:     menuservice.NewMenuService(sysClient),
		RoleService:     roleservice.NewRoleService(sysClient),
		UserService:     userservice.NewUserService(sysClient),
		CheckUrl:        middleware.NewCheckUrlMiddleware(newRedis).Handle,
		AddLog:          middleware.NewAddLogMiddleware(logService).Handle,
		Redis:           newRedis,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
