// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package server

import (
	"context"

	"zero-admin/rpc/ums/internal/logic/membermembertagrelationservice"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"
)

type MemberMemberTagRelationServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedMemberMemberTagRelationServiceServer
}

func NewMemberMemberTagRelationServiceServer(svcCtx *svc.ServiceContext) *MemberMemberTagRelationServiceServer {
	return &MemberMemberTagRelationServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *MemberMemberTagRelationServiceServer) MemberMemberTagRelationAdd(ctx context.Context, in *umsclient.MemberMemberTagRelationAddReq) (*umsclient.MemberMemberTagRelationAddResp, error) {
	l := membermembertagrelationservicelogic.NewMemberMemberTagRelationAddLogic(ctx, s.svcCtx)
	return l.MemberMemberTagRelationAdd(in)
}

func (s *MemberMemberTagRelationServiceServer) MemberMemberTagRelationList(ctx context.Context, in *umsclient.MemberMemberTagRelationListReq) (*umsclient.MemberMemberTagRelationListResp, error) {
	l := membermembertagrelationservicelogic.NewMemberMemberTagRelationListLogic(ctx, s.svcCtx)
	return l.MemberMemberTagRelationList(in)
}

func (s *MemberMemberTagRelationServiceServer) MemberMemberTagRelationUpdate(ctx context.Context, in *umsclient.MemberMemberTagRelationUpdateReq) (*umsclient.MemberMemberTagRelationUpdateResp, error) {
	l := membermembertagrelationservicelogic.NewMemberMemberTagRelationUpdateLogic(ctx, s.svcCtx)
	return l.MemberMemberTagRelationUpdate(in)
}

func (s *MemberMemberTagRelationServiceServer) MemberMemberTagRelationDelete(ctx context.Context, in *umsclient.MemberMemberTagRelationDeleteReq) (*umsclient.MemberMemberTagRelationDeleteResp, error) {
	l := membermembertagrelationservicelogic.NewMemberMemberTagRelationDeleteLogic(ctx, s.svcCtx)
	return l.MemberMemberTagRelationDelete(in)
}
