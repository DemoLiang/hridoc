package main

import (
	"flag"
	"fmt"
	"zero-admin/rpc/hridoc/internal/config"
	configserviceServer "zero-admin/rpc/hridoc/internal/server/configservice"
	deptserviceServer "zero-admin/rpc/hridoc/internal/server/deptservice"
	dictserviceServer "zero-admin/rpc/hridoc/internal/server/dictservice"
	idocserviceServer "zero-admin/rpc/hridoc/internal/server/idocservice"
	jobserviceServer "zero-admin/rpc/hridoc/internal/server/jobservice"
	loginlogserviceServer "zero-admin/rpc/hridoc/internal/server/loginlogservice"
	menuserviceServer "zero-admin/rpc/hridoc/internal/server/menuservice"
	roleserviceServer "zero-admin/rpc/hridoc/internal/server/roleservice"
	syslogserviceServer "zero-admin/rpc/hridoc/internal/server/syslogservice"
	userserviceServer "zero-admin/rpc/hridoc/internal/server/userservice"
	"zero-admin/rpc/hridoc/internal/svc"
	"zero-admin/rpc/proto/idoc"
	"zero-admin/rpc/proto/sys"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "rpc/sys/etc/sys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sys.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))
		sys.RegisterRoleServiceServer(grpcServer, roleserviceServer.NewRoleServiceServer(ctx))
		sys.RegisterMenuServiceServer(grpcServer, menuserviceServer.NewMenuServiceServer(ctx))
		sys.RegisterDictServiceServer(grpcServer, dictserviceServer.NewDictServiceServer(ctx))
		sys.RegisterDeptServiceServer(grpcServer, deptserviceServer.NewDeptServiceServer(ctx))
		sys.RegisterLoginLogServiceServer(grpcServer, loginlogserviceServer.NewLoginLogServiceServer(ctx))
		sys.RegisterSysLogServiceServer(grpcServer, syslogserviceServer.NewSysLogServiceServer(ctx))
		sys.RegisterConfigServiceServer(grpcServer, configserviceServer.NewConfigServiceServer(ctx))
		sys.RegisterJobServiceServer(grpcServer, jobserviceServer.NewJobServiceServer(ctx))
		idoc.RegisterIdocServiceServer(grpcServer, idocserviceServer.NewIdocServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
