// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.BaseResp, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.BaseResp{Code: errorx.ErrUserNotFound, Message: "用户不存在"}, nil
		}
		logx.Errorf("find user failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	user.Name = req.Name
	user.Phone = sql.NullString{String: req.Phone, Valid: req.Phone != ""}
	user.Email = sql.NullString{String: req.Email, Valid: req.Email != ""}
	user.IdCard = req.IdCard
	user.Education = sql.NullString{String: req.Education, Valid: req.Education != ""}
	user.Role = req.Role
	user.Status = req.Status

	if req.Password != "" {
		hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			logx.Errorf("hash password failed: %v", err)
			return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
		}
		user.Password = string(hashedPwd)
	}

	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		logx.Errorf("update user failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	return &types.BaseResp{Code: 0, Message: "success"}, nil
}
