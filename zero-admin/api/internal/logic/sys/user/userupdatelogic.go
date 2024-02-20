package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UserUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 14:05
*/
type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserUpdateLogic {
	return UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserUpdate 更新用户信息
func (l *UserUpdateLogic) UserUpdate(req types.UpdateUserReq) (*types.UpdateUserResp, error) {
	userUpdateReq := sysclient.UserUpdateReq{
		Id:           req.Id,
		Email:        req.Email,
		Mobile:       req.Mobile,
		Name:         req.Name,
		NickName:     req.NickName,
		DeptId:       req.DeptId,
		LastUpdateBy: l.ctx.Value("userName").(string),
		RoleId:       req.RoleId,
		Status:       req.Status,
		JobId:        req.JobId,
	}

	if _, err := l.svcCtx.UserService.UserUpdate(l.ctx, &userUpdateReq); err != nil {
		logc.Errorf(l.ctx, "更新用户信息失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新用户失败")
	}

	return &types.UpdateUserResp{
		Code:    "000000",
		Message: "更新用户成功",
	}, nil
}
