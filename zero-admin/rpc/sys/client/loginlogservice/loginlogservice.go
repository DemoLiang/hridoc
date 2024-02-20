// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package loginlogservice

import (
	"context"

	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ConfigAddReq           = sysclient.ConfigAddReq
	ConfigAddResp          = sysclient.ConfigAddResp
	ConfigDeleteReq        = sysclient.ConfigDeleteReq
	ConfigDeleteResp       = sysclient.ConfigDeleteResp
	ConfigListData         = sysclient.ConfigListData
	ConfigListReq          = sysclient.ConfigListReq
	ConfigListResp         = sysclient.ConfigListResp
	ConfigUpdateReq        = sysclient.ConfigUpdateReq
	ConfigUpdateResp       = sysclient.ConfigUpdateResp
	DeptAddReq             = sysclient.DeptAddReq
	DeptAddResp            = sysclient.DeptAddResp
	DeptDeleteReq          = sysclient.DeptDeleteReq
	DeptDeleteResp         = sysclient.DeptDeleteResp
	DeptListData           = sysclient.DeptListData
	DeptListReq            = sysclient.DeptListReq
	DeptListResp           = sysclient.DeptListResp
	DeptUpdateReq          = sysclient.DeptUpdateReq
	DeptUpdateResp         = sysclient.DeptUpdateResp
	DictAddReq             = sysclient.DictAddReq
	DictAddResp            = sysclient.DictAddResp
	DictDeleteReq          = sysclient.DictDeleteReq
	DictDeleteResp         = sysclient.DictDeleteResp
	DictListData           = sysclient.DictListData
	DictListReq            = sysclient.DictListReq
	DictListResp           = sysclient.DictListResp
	DictUpdateReq          = sysclient.DictUpdateReq
	DictUpdateResp         = sysclient.DictUpdateResp
	InfoReq                = sysclient.InfoReq
	InfoResp               = sysclient.InfoResp
	JobAddReq              = sysclient.JobAddReq
	JobAddResp             = sysclient.JobAddResp
	JobDeleteReq           = sysclient.JobDeleteReq
	JobDeleteResp          = sysclient.JobDeleteResp
	JobListData            = sysclient.JobListData
	JobListReq             = sysclient.JobListReq
	JobListResp            = sysclient.JobListResp
	JobUpdateReq           = sysclient.JobUpdateReq
	JobUpdateResp          = sysclient.JobUpdateResp
	LoginLogAddReq         = sysclient.LoginLogAddReq
	LoginLogAddResp        = sysclient.LoginLogAddResp
	LoginLogDeleteReq      = sysclient.LoginLogDeleteReq
	LoginLogDeleteResp     = sysclient.LoginLogDeleteResp
	LoginLogListData       = sysclient.LoginLogListData
	LoginLogListReq        = sysclient.LoginLogListReq
	LoginLogListResp       = sysclient.LoginLogListResp
	LoginReq               = sysclient.LoginReq
	LoginResp              = sysclient.LoginResp
	MenuAddReq             = sysclient.MenuAddReq
	MenuAddResp            = sysclient.MenuAddResp
	MenuDeleteReq          = sysclient.MenuDeleteReq
	MenuDeleteResp         = sysclient.MenuDeleteResp
	MenuListData           = sysclient.MenuListData
	MenuListReq            = sysclient.MenuListReq
	MenuListResp           = sysclient.MenuListResp
	MenuListTree           = sysclient.MenuListTree
	MenuUpdateReq          = sysclient.MenuUpdateReq
	MenuUpdateResp         = sysclient.MenuUpdateResp
	QueryMenuByRoleIdReq   = sysclient.QueryMenuByRoleIdReq
	QueryMenuByRoleIdResp  = sysclient.QueryMenuByRoleIdResp
	ReSetPasswordReq       = sysclient.ReSetPasswordReq
	ReSetPasswordResp      = sysclient.ReSetPasswordResp
	RoleAddReq             = sysclient.RoleAddReq
	RoleAddResp            = sysclient.RoleAddResp
	RoleDeleteReq          = sysclient.RoleDeleteReq
	RoleDeleteResp         = sysclient.RoleDeleteResp
	RoleListData           = sysclient.RoleListData
	RoleListReq            = sysclient.RoleListReq
	RoleListResp           = sysclient.RoleListResp
	RoleUpdateReq          = sysclient.RoleUpdateReq
	RoleUpdateResp         = sysclient.RoleUpdateResp
	StatisticsLoginLogReq  = sysclient.StatisticsLoginLogReq
	StatisticsLoginLogResp = sysclient.StatisticsLoginLogResp
	SysLogAddReq           = sysclient.SysLogAddReq
	SysLogAddResp          = sysclient.SysLogAddResp
	SysLogDeleteReq        = sysclient.SysLogDeleteReq
	SysLogDeleteResp       = sysclient.SysLogDeleteResp
	SysLogListData         = sysclient.SysLogListData
	SysLogListReq          = sysclient.SysLogListReq
	SysLogListResp         = sysclient.SysLogListResp
	UpdateConfigRoleReq    = sysclient.UpdateConfigRoleReq
	UpdateConfigRoleResp   = sysclient.UpdateConfigRoleResp
	UpdateMenuRoleReq      = sysclient.UpdateMenuRoleReq
	UpdateMenuRoleResp     = sysclient.UpdateMenuRoleResp
	UserAddReq             = sysclient.UserAddReq
	UserAddResp            = sysclient.UserAddResp
	UserDeleteReq          = sysclient.UserDeleteReq
	UserDeleteResp         = sysclient.UserDeleteResp
	UserListData           = sysclient.UserListData
	UserListReq            = sysclient.UserListReq
	UserListResp           = sysclient.UserListResp
	UserStatusReq          = sysclient.UserStatusReq
	UserStatusResp         = sysclient.UserStatusResp
	UserUpdateReq          = sysclient.UserUpdateReq
	UserUpdateResp         = sysclient.UserUpdateResp

	LoginLogService interface {
		LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error)
		LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error)
		LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error)
		// 统计后台用户登录---(查询当天登录人数（根据IP,统计当前周登录人数（根据IP）,统计当前月登录人数（根据IP）)
		StatisticsLoginLog(ctx context.Context, in *StatisticsLoginLogReq, opts ...grpc.CallOption) (*StatisticsLoginLogResp, error)
	}

	defaultLoginLogService struct {
		cli zrpc.Client
	}
)

func NewLoginLogService(cli zrpc.Client) LoginLogService {
	return &defaultLoginLogService{
		cli: cli,
	}
}

func (m *defaultLoginLogService) LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error) {
	client := sysclient.NewLoginLogServiceClient(m.cli.Conn())
	return client.LoginLogAdd(ctx, in, opts...)
}

func (m *defaultLoginLogService) LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error) {
	client := sysclient.NewLoginLogServiceClient(m.cli.Conn())
	return client.LoginLogList(ctx, in, opts...)
}

func (m *defaultLoginLogService) LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error) {
	client := sysclient.NewLoginLogServiceClient(m.cli.Conn())
	return client.LoginLogDelete(ctx, in, opts...)
}

// 统计后台用户登录---(查询当天登录人数（根据IP,统计当前周登录人数（根据IP）,统计当前月登录人数（根据IP）)
func (m *defaultLoginLogService) StatisticsLoginLog(ctx context.Context, in *StatisticsLoginLogReq, opts ...grpc.CallOption) (*StatisticsLoginLogResp, error) {
	client := sysclient.NewLoginLogServiceClient(m.cli.Conn())
	return client.StatisticsLoginLog(ctx, in, opts...)
}
