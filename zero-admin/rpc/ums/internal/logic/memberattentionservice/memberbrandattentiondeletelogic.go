package memberattentionservicelogic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberBrandAttentionDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/4 16:53
*/
type MemberBrandAttentionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberBrandAttentionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberBrandAttentionDeleteLogic {
	return &MemberBrandAttentionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberBrandAttentionDelete 取消关注
func (l *MemberBrandAttentionDeleteLogic) MemberBrandAttentionDelete(in *umsclient.MemberBrandAttentionDeleteReq) (*umsclient.MemberBrandAttentionDeleteResp, error) {
	err := l.svcCtx.UmsMemberBrandAttentionModel.DeleteByIdAndMemberId(l.ctx, in.Id, in.MemberId)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberBrandAttentionDeleteResp{}, nil
}
