// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"zero-admin/rpc/hridoc/internal/logic/jobservice"
	"zero-admin/rpc/hridoc/internal/svc"
	"zero-admin/rpc/proto/sys"
)

type JobServiceServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedJobServiceServer
}

func NewJobServiceServer(svcCtx *svc.ServiceContext) *JobServiceServer {
	return &JobServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *JobServiceServer) JobAdd(ctx context.Context, in *sys.JobAddReq) (*sys.JobAddResp, error) {
	l := jobservicelogic.NewJobAddLogic(ctx, s.svcCtx)
	return l.JobAdd(in)
}

func (s *JobServiceServer) JobList(ctx context.Context, in *sys.JobListReq) (*sys.JobListResp, error) {
	l := jobservicelogic.NewJobListLogic(ctx, s.svcCtx)
	return l.JobList(in)
}

func (s *JobServiceServer) JobUpdate(ctx context.Context, in *sys.JobUpdateReq) (*sys.JobUpdateResp, error) {
	l := jobservicelogic.NewJobUpdateLogic(ctx, s.svcCtx)
	return l.JobUpdate(in)
}

func (s *JobServiceServer) JobDelete(ctx context.Context, in *sys.JobDeleteReq) (*sys.JobDeleteResp, error) {
	l := jobservicelogic.NewJobDeleteLogic(ctx, s.svcCtx)
	return l.JobDelete(in)
}