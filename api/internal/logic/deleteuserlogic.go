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

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteReq) (resp *types.BaseResp, err error) {
	err = l.svcCtx.UserModel.Delete(l.ctx, req.Id)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.BaseResp{Code: errorx.ErrUserNotFound, Message: "用户不存在"}, nil
		}
		logx.Errorf("delete user failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	return &types.BaseResp{Code: 0, Message: "success"}, nil
}
