package couponhistoryservicelogic

import (
	"context"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCouponStatusLogic
/*
Author: LiuFeiHua
Date: 2023/12/7 10:39
*/
type UpdateCouponStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCouponStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponStatusLogic {
	return &UpdateCouponStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCouponStatus 更新优惠券状态
//1.更新用户优惠券状态
//2.更新优惠券数量
func (l *UpdateCouponStatusLogic) UpdateCouponStatus(in *smsclient.UpdateCouponStatusReq) (*smsclient.UpdateCouponStatusResp, error) {
	//1.更新用户优惠券状态
	err := l.svcCtx.SmsCouponHistoryModel.UpdateCouponStatus(l.ctx, in.CouponId, in.MemberId, in.UseStatus)
	if err != nil {
		return nil, err
	}

	//2.更新优惠券数量
	err = l.svcCtx.SmsCouponModel.UpdateUseCount(l.ctx, in.UseStatus == 0, in.CouponId)
	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateCouponStatusResp{}, nil
}
