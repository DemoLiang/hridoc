package memberreceiveaddressservicelogic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberReceiveAddressQueryDetailLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 11:25
*/
type MemberReceiveAddressQueryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressQueryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressQueryDetailLogic {
	return &MemberReceiveAddressQueryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReceiveAddressQueryDetailLogic) MemberReceiveAddressQueryDetail(in *umsclient.MemberReceiveAddressQueryDetailReq) (*umsclient.MemberReceiveAddressQueryDetailResp, error) {
	address, err := l.svcCtx.UmsMemberReceiveAddressModel.FindByIdAndMemberId(l.ctx, in.AddressID, in.UserId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询会员地址信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	resp := &umsclient.MemberReceiveAddressQueryDetailResp{
		Id:            address.Id,
		MemberId:      address.MemberId,
		Name:          address.Name,
		PhoneNumber:   address.PhoneNumber,
		DefaultStatus: address.DefaultStatus,
		PostCode:      address.PostCode,
		Province:      address.Province,
		City:          address.City,
		Region:        address.Region,
		DetailAddress: address.DetailAddress,
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(resp)
	logc.Infof(l.ctx, "查询会员地址信息,参数：%s,响应：%s", reqStr, listStr)

	return resp, nil
}
