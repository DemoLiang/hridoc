// Code generated by goctl. DO NOT EDIT.
// Source: cms.proto

package server

import (
	"context"

	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/logic/prefrenceareaproductrelationservice"
	"zero-admin/rpc/cms/internal/svc"
)

type PrefrenceAreaProductRelationServiceServer struct {
	svcCtx *svc.ServiceContext
	cmsclient.UnimplementedPrefrenceAreaProductRelationServiceServer
}

func NewPrefrenceAreaProductRelationServiceServer(svcCtx *svc.ServiceContext) *PrefrenceAreaProductRelationServiceServer {
	return &PrefrenceAreaProductRelationServiceServer{
		svcCtx: svcCtx,
	}
}

// 优选商品关联
func (s *PrefrenceAreaProductRelationServiceServer) PrefrenceAreaProductRelationAdd(ctx context.Context, in *cmsclient.PrefrenceAreaProductRelationAddReq) (*cmsclient.PrefrenceAreaProductRelationAddResp, error) {
	l := prefrenceareaproductrelationservicelogic.NewPrefrenceAreaProductRelationAddLogic(ctx, s.svcCtx)
	return l.PrefrenceAreaProductRelationAdd(in)
}

func (s *PrefrenceAreaProductRelationServiceServer) PrefrenceAreaProductRelationDelete(ctx context.Context, in *cmsclient.PrefrenceAreaProductRelationDeleteReq) (*cmsclient.PrefrenceAreaProductRelationDeleteResp, error) {
	l := prefrenceareaproductrelationservicelogic.NewPrefrenceAreaProductRelationDeleteLogic(ctx, s.svcCtx)
	return l.PrefrenceAreaProductRelationDelete(in)
}

func (s *PrefrenceAreaProductRelationServiceServer) PrefrenceAreaProductRelationUpdate(ctx context.Context, in *cmsclient.PrefrenceAreaProductRelationUpdateReq) (*cmsclient.PrefrenceAreaProductRelationUpdateResp, error) {
	l := prefrenceareaproductrelationservicelogic.NewPrefrenceAreaProductRelationUpdateLogic(ctx, s.svcCtx)
	return l.PrefrenceAreaProductRelationUpdate(in)
}

func (s *PrefrenceAreaProductRelationServiceServer) PrefrenceAreaProductRelationList(ctx context.Context, in *cmsclient.PrefrenceAreaProductRelationListReq) (*cmsclient.PrefrenceAreaProductRelationListResp, error) {
	l := prefrenceareaproductrelationservicelogic.NewPrefrenceAreaProductRelationListLogic(ctx, s.svcCtx)
	return l.PrefrenceAreaProductRelationList(in)
}
