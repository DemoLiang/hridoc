// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"zero-admin/rpc/hridoc/internal/logic/syslogservice"
	"zero-admin/rpc/hridoc/internal/svc"
	"zero-admin/rpc/proto/sys"
)

type SysLogServiceServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedSysLogServiceServer
}

func NewSysLogServiceServer(svcCtx *svc.ServiceContext) *SysLogServiceServer {
	return &SysLogServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *SysLogServiceServer) SysLogAdd(ctx context.Context, in *sys.SysLogAddReq) (*sys.SysLogAddResp, error) {
	l := syslogservicelogic.NewSysLogAddLogic(ctx, s.svcCtx)
	return l.SysLogAdd(in)
}

func (s *SysLogServiceServer) SysLogList(ctx context.Context, in *sys.SysLogListReq) (*sys.SysLogListResp, error) {
	l := syslogservicelogic.NewSysLogListLogic(ctx, s.svcCtx)
	return l.SysLogList(in)
}

func (s *SysLogServiceServer) SysLogDelete(ctx context.Context, in *sys.SysLogDeleteReq) (*sys.SysLogDeleteResp, error) {
	l := syslogservicelogic.NewSysLogDeleteLogic(ctx, s.svcCtx)
	return l.SysLogDelete(in)
}
