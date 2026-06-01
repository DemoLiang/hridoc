// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"

	"github.com/DemoLiang/hridoc/api/internal/middleware"
	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserInfoResp, err error) {
	userId := middleware.GetUserId(l.ctx)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.UserInfoResp{
				BaseResp: types.BaseResp{Code: errorx.ErrUserNotFound, Message: "用户不存在"},
			}, nil
		}
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

func nullString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}
