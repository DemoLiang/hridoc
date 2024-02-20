// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package server

import (
	"context"

	"zero-admin/rpc/pms/internal/logic/productvertifyrecordservice"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"
)

type ProductVertifyRecordServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedProductVertifyRecordServiceServer
}

func NewProductVertifyRecordServiceServer(svcCtx *svc.ServiceContext) *ProductVertifyRecordServiceServer {
	return &ProductVertifyRecordServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductVertifyRecordServiceServer) ProductVertifyRecordAdd(ctx context.Context, in *pmsclient.ProductVertifyRecordAddReq) (*pmsclient.ProductVertifyRecordAddResp, error) {
	l := productvertifyrecordservicelogic.NewProductVertifyRecordAddLogic(ctx, s.svcCtx)
	return l.ProductVertifyRecordAdd(in)
}

func (s *ProductVertifyRecordServiceServer) ProductVertifyRecordList(ctx context.Context, in *pmsclient.ProductVertifyRecordListReq) (*pmsclient.ProductVertifyRecordListResp, error) {
	l := productvertifyrecordservicelogic.NewProductVertifyRecordListLogic(ctx, s.svcCtx)
	return l.ProductVertifyRecordList(in)
}

func (s *ProductVertifyRecordServiceServer) ProductVertifyRecordUpdate(ctx context.Context, in *pmsclient.ProductVertifyRecordUpdateReq) (*pmsclient.ProductVertifyRecordUpdateResp, error) {
	l := productvertifyrecordservicelogic.NewProductVertifyRecordUpdateLogic(ctx, s.svcCtx)
	return l.ProductVertifyRecordUpdate(in)
}

func (s *ProductVertifyRecordServiceServer) ProductVertifyRecordDelete(ctx context.Context, in *pmsclient.ProductVertifyRecordDeleteReq) (*pmsclient.ProductVertifyRecordDeleteResp, error) {
	l := productvertifyrecordservicelogic.NewProductVertifyRecordDeleteLogic(ctx, s.svcCtx)
	return l.ProductVertifyRecordDelete(in)
}
