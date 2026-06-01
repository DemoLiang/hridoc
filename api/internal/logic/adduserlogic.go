// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/model"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.UserModel.FindOneByName(l.ctx, req.Name)
	if err == nil {
		return &types.BaseResp{Code: errorx.ErrUserExists, Message: "用户已存在"}, nil
	}
	if !errorx.IsNotFound(err) {
		logx.Errorf("query user failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logx.Errorf("hash password failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Name:      req.Name,
		Phone:     sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Email:     sql.NullString{String: req.Email, Valid: req.Email != ""},
		IdCard:    req.IdCard,
		Education: sql.NullString{String: req.Education, Valid: req.Education != ""},
		Role:      req.Role,
		Status:    1,
		Password:  string(hashedPwd),
	})
	if err != nil {
		logx.Errorf("insert user failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	return &types.BaseResp{Code: 0, Message: "success"}, nil
}
