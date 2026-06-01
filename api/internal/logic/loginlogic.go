package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/auth"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.Username)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.LoginResp{
				BaseResp: types.BaseResp{Code: errorx.ErrUserNotFound, Message: "用户名或密码错误"},
			}, nil
		}
		return &types.LoginResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	if user.Status != 1 {
		return &types.LoginResp{
			BaseResp: types.BaseResp{Code: errorx.ErrUserNotFound, Message: "用户已被禁用"},
		}, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return &types.LoginResp{
			BaseResp: types.BaseResp{Code: errorx.ErrPassword, Message: "用户名或密码错误"},
		}, nil
	}

	jwtUtil := auth.NewJWT(l.svcCtx.Config.JwtAuth.AccessSecret, l.svcCtx.Config.JwtAuth.AccessExpire)
	token, err := jwtUtil.GenerateToken(user.Id, user.Name, int(user.Role))
	if err != nil {
		logx.Errorf("generate token failed: %v", err)
		return &types.LoginResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	return &types.LoginResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.LoginData{
			Token:  token,
			Expire: l.svcCtx.Config.JwtAuth.AccessExpire,
		},
	}, nil
}
