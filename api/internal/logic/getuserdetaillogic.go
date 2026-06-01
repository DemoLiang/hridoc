// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailLogic) GetUserDetail(req *types.UserDetailReq) (resp *types.UserInfoResp, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.UserInfoResp{
				BaseResp: types.BaseResp{Code: errorx.ErrUserNotFound, Message: "用户不存在"},
			}, nil
		}
		logx.Errorf("get user detail failed: %v", err)
		return &types.UserInfoResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	return &types.UserInfoResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.UserInfo{
			Id:        user.Id,
			Name:      user.Name,
			Phone:     nullString(user.Phone),
			Email:     nullString(user.Email),
			IdCard:    user.IdCard,
			Education: nullString(user.Education),
			Role:      user.Role,
			Status:    user.Status,
		},
	}, nil
}
