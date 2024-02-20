// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package server

import (
	"context"

	"zero-admin/rpc/sms/internal/logic/couponhistoryservice"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"
)

type CouponHistoryServiceServer struct {
	svcCtx *svc.ServiceContext
	smsclient.UnimplementedCouponHistoryServiceServer
}

func NewCouponHistoryServiceServer(svcCtx *svc.ServiceContext) *CouponHistoryServiceServer {
	return &CouponHistoryServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CouponHistoryServiceServer) CouponHistoryAdd(ctx context.Context, in *smsclient.CouponHistoryAddReq) (*smsclient.CouponHistoryAddResp, error) {
	l := couponhistoryservicelogic.NewCouponHistoryAddLogic(ctx, s.svcCtx)
	return l.CouponHistoryAdd(in)
}

func (s *CouponHistoryServiceServer) CouponHistoryList(ctx context.Context, in *smsclient.CouponHistoryListReq) (*smsclient.CouponHistoryListResp, error) {
	l := couponhistoryservicelogic.NewCouponHistoryListLogic(ctx, s.svcCtx)
	return l.CouponHistoryList(in)
}

func (s *CouponHistoryServiceServer) CouponHistoryUpdate(ctx context.Context, in *smsclient.CouponHistoryUpdateReq) (*smsclient.CouponHistoryUpdateResp, error) {
	l := couponhistoryservicelogic.NewCouponHistoryUpdateLogic(ctx, s.svcCtx)
	return l.CouponHistoryUpdate(in)
}

func (s *CouponHistoryServiceServer) CouponHistoryDelete(ctx context.Context, in *smsclient.CouponHistoryDeleteReq) (*smsclient.CouponHistoryDeleteResp, error) {
	l := couponhistoryservicelogic.NewCouponHistoryDeleteLogic(ctx, s.svcCtx)
	return l.CouponHistoryDelete(in)
}

// 登录时获取用户还没有使用的获取优惠券数量
func (s *CouponHistoryServiceServer) CouponCount(ctx context.Context, in *smsclient.CouponCountReq) (*smsclient.CouponCountResp, error) {
	l := couponhistoryservicelogic.NewCouponCountLogic(ctx, s.svcCtx)
	return l.CouponCount(in)
}

// 获取会员优惠券
func (s *CouponHistoryServiceServer) QueryMemberCouponList(ctx context.Context, in *smsclient.QueryMemberCouponListReq) (*smsclient.QueryMemberCouponListResp, error) {
	l := couponhistoryservicelogic.NewQueryMemberCouponListLogic(ctx, s.svcCtx)
	return l.QueryMemberCouponList(in)
}

// 更新优惠券状态
func (s *CouponHistoryServiceServer) UpdateCouponStatus(ctx context.Context, in *smsclient.UpdateCouponStatusReq) (*smsclient.UpdateCouponStatusResp, error) {
	l := couponhistoryservicelogic.NewUpdateCouponStatusLogic(ctx, s.svcCtx)
	return l.UpdateCouponStatus(in)
}

// 获取该用户所有优惠券(包括商品和优惠券,商品分类和优惠券的关联关糸)
func (s *CouponHistoryServiceServer) QueryCouponHistoryDetailList(ctx context.Context, in *smsclient.CouponHistoryDetailListReq) (*smsclient.CouponHistoryDetailListResp, error) {
	l := couponhistoryservicelogic.NewQueryCouponHistoryDetailListLogic(ctx, s.svcCtx)
	return l.QueryCouponHistoryDetailList(in)
}
