// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"zero-admin/rpc/sys/internal/logic/userservice"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) Login(ctx context.Context, in *sysclient.LoginReq) (*sysclient.LoginResp, error) {
	l := userservicelogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServiceServer) UserInfo(ctx context.Context, in *sysclient.InfoReq) (*sysclient.InfoResp, error) {
	l := userservicelogic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UserServiceServer) UserAdd(ctx context.Context, in *sysclient.UserAddReq) (*sysclient.UserAddResp, error) {
	l := userservicelogic.NewUserAddLogic(ctx, s.svcCtx)
	return l.UserAdd(in)
}

func (s *UserServiceServer) UserList(ctx context.Context, in *sysclient.UserListReq) (*sysclient.UserListResp, error) {
	l := userservicelogic.NewUserListLogic(ctx, s.svcCtx)
	return l.UserList(in)
}

func (s *UserServiceServer) UserUpdate(ctx context.Context, in *sysclient.UserUpdateReq) (*sysclient.UserUpdateResp, error) {
	l := userservicelogic.NewUserUpdateLogic(ctx, s.svcCtx)
	return l.UserUpdate(in)
}

func (s *UserServiceServer) UserDelete(ctx context.Context, in *sysclient.UserDeleteReq) (*sysclient.UserDeleteResp, error) {
	l := userservicelogic.NewUserDeleteLogic(ctx, s.svcCtx)
	return l.UserDelete(in)
}

func (s *UserServiceServer) ReSetPassword(ctx context.Context, in *sysclient.ReSetPasswordReq) (*sysclient.ReSetPasswordResp, error) {
	l := userservicelogic.NewReSetPasswordLogic(ctx, s.svcCtx)
	return l.ReSetPassword(in)
}

func (s *UserServiceServer) UpdateUserStatus(ctx context.Context, in *sysclient.UserStatusReq) (*sysclient.UserStatusResp, error) {
	l := userservicelogic.NewUpdateUserStatusLogic(ctx, s.svcCtx)
	return l.UpdateUserStatus(in)
}
