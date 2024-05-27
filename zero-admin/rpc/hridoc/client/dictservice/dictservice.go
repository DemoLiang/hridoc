// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package dictservice

import (
	"context"

	"zero-admin/rpc/proto/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ConfigAddReq           = sys.ConfigAddReq
	ConfigAddResp          = sys.ConfigAddResp
	ConfigDeleteReq        = sys.ConfigDeleteReq
	ConfigDeleteResp       = sys.ConfigDeleteResp
	ConfigListData         = sys.ConfigListData
	ConfigListReq          = sys.ConfigListReq
	ConfigListResp         = sys.ConfigListResp
	ConfigUpdateReq        = sys.ConfigUpdateReq
	ConfigUpdateResp       = sys.ConfigUpdateResp
	DeptAddReq             = sys.DeptAddReq
	DeptAddResp            = sys.DeptAddResp
	DeptDeleteReq          = sys.DeptDeleteReq
	DeptDeleteResp         = sys.DeptDeleteResp
	DeptListData           = sys.DeptListData
	DeptListReq            = sys.DeptListReq
	DeptListResp           = sys.DeptListResp
	DeptUpdateReq          = sys.DeptUpdateReq
	DeptUpdateResp         = sys.DeptUpdateResp
	DictAddReq             = sys.DictAddReq
	DictAddResp            = sys.DictAddResp
	DictDeleteReq          = sys.DictDeleteReq
	DictDeleteResp         = sys.DictDeleteResp
	DictListData           = sys.DictListData
	DictListReq            = sys.DictListReq
	DictListResp           = sys.DictListResp
	DictUpdateReq          = sys.DictUpdateReq
	DictUpdateResp         = sys.DictUpdateResp
	InfoReq                = sys.InfoReq
	InfoResp               = sys.InfoResp
	JobAddReq              = sys.JobAddReq
	JobAddResp             = sys.JobAddResp
	JobDeleteReq           = sys.JobDeleteReq
	JobDeleteResp          = sys.JobDeleteResp
	JobListData            = sys.JobListData
	JobListReq             = sys.JobListReq
	JobListResp            = sys.JobListResp
	JobUpdateReq           = sys.JobUpdateReq
	JobUpdateResp          = sys.JobUpdateResp
	LoginLogAddReq         = sys.LoginLogAddReq
	LoginLogAddResp        = sys.LoginLogAddResp
	LoginLogDeleteReq      = sys.LoginLogDeleteReq
	LoginLogDeleteResp     = sys.LoginLogDeleteResp
	LoginLogListData       = sys.LoginLogListData
	LoginLogListReq        = sys.LoginLogListReq
	LoginLogListResp       = sys.LoginLogListResp
	LoginReq               = sys.LoginReq
	LoginResp              = sys.LoginResp
	MenuAddReq             = sys.MenuAddReq
	MenuAddResp            = sys.MenuAddResp
	MenuDeleteReq          = sys.MenuDeleteReq
	MenuDeleteResp         = sys.MenuDeleteResp
	MenuListData           = sys.MenuListData
	MenuListReq            = sys.MenuListReq
	MenuListResp           = sys.MenuListResp
	MenuListTree           = sys.MenuListTree
	MenuUpdateReq          = sys.MenuUpdateReq
	MenuUpdateResp         = sys.MenuUpdateResp
	QueryMenuByRoleIdReq   = sys.QueryMenuByRoleIdReq
	QueryMenuByRoleIdResp  = sys.QueryMenuByRoleIdResp
	ReSetPasswordReq       = sys.ReSetPasswordReq
	ReSetPasswordResp      = sys.ReSetPasswordResp
	RoleAddReq             = sys.RoleAddReq
	RoleAddResp            = sys.RoleAddResp
	RoleDeleteReq          = sys.RoleDeleteReq
	RoleDeleteResp         = sys.RoleDeleteResp
	RoleListData           = sys.RoleListData
	RoleListReq            = sys.RoleListReq
	RoleListResp           = sys.RoleListResp
	RoleUpdateReq          = sys.RoleUpdateReq
	RoleUpdateResp         = sys.RoleUpdateResp
	StatisticsLoginLogReq  = sys.StatisticsLoginLogReq
	StatisticsLoginLogResp = sys.StatisticsLoginLogResp
	SysLogAddReq           = sys.SysLogAddReq
	SysLogAddResp          = sys.SysLogAddResp
	SysLogDeleteReq        = sys.SysLogDeleteReq
	SysLogDeleteResp       = sys.SysLogDeleteResp
	SysLogListData         = sys.SysLogListData
	SysLogListReq          = sys.SysLogListReq
	SysLogListResp         = sys.SysLogListResp
	UpdateConfigRoleReq    = sys.UpdateConfigRoleReq
	UpdateConfigRoleResp   = sys.UpdateConfigRoleResp
	UpdateMenuRoleReq      = sys.UpdateMenuRoleReq
	UpdateMenuRoleResp     = sys.UpdateMenuRoleResp
	UserAddReq             = sys.UserAddReq
	UserAddResp            = sys.UserAddResp
	UserDeleteReq          = sys.UserDeleteReq
	UserDeleteResp         = sys.UserDeleteResp
	UserListData           = sys.UserListData
	UserListReq            = sys.UserListReq
	UserListResp           = sys.UserListResp
	UserStatusReq          = sys.UserStatusReq
	UserStatusResp         = sys.UserStatusResp
	UserUpdateReq          = sys.UserUpdateReq
	UserUpdateResp         = sys.UserUpdateResp

	DictService interface {
		DictAdd(ctx context.Context, in *DictAddReq, opts ...grpc.CallOption) (*DictAddResp, error)
		DictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error)
		DictUpdate(ctx context.Context, in *DictUpdateReq, opts ...grpc.CallOption) (*DictUpdateResp, error)
		DictDelete(ctx context.Context, in *DictDeleteReq, opts ...grpc.CallOption) (*DictDeleteResp, error)
	}

	defaultDictService struct {
		cli zrpc.Client
	}
)

func NewDictService(cli zrpc.Client) DictService {
	return &defaultDictService{
		cli: cli,
	}
}

func (m *defaultDictService) DictAdd(ctx context.Context, in *DictAddReq, opts ...grpc.CallOption) (*DictAddResp, error) {
	client := sys.NewDictServiceClient(m.cli.Conn())
	return client.DictAdd(ctx, in, opts...)
}

func (m *defaultDictService) DictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error) {
	client := sys.NewDictServiceClient(m.cli.Conn())
	return client.DictList(ctx, in, opts...)
}

func (m *defaultDictService) DictUpdate(ctx context.Context, in *DictUpdateReq, opts ...grpc.CallOption) (*DictUpdateResp, error) {
	client := sys.NewDictServiceClient(m.cli.Conn())
	return client.DictUpdate(ctx, in, opts...)
}

func (m *defaultDictService) DictDelete(ctx context.Context, in *DictDeleteReq, opts ...grpc.CallOption) (*DictDeleteResp, error) {
	client := sys.NewDictServiceClient(m.cli.Conn())
	return client.DictDelete(ctx, in, opts...)
}
