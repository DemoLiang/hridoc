package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/proto/sys"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRoleMenuLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:43
*/
type UpdateRoleMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateRoleMenuLogic {
	return UpdateRoleMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRoleMenu 更新角色与菜单的关联
func (l *UpdateRoleMenuLogic) UpdateRoleMenu(req types.UpdateRoleMenuReq) (*types.UpdateRoleMenuResp, error) {
	_, err := l.svcCtx.RoleService.UpdateMenuRole(l.ctx, &sys.UpdateMenuRoleReq{
		RoleId:   req.RoleId,
		MenuIds:  req.MenuIds,
		CreateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新角色菜单失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新角色菜单失败")
	}

	return &types.UpdateRoleMenuResp{
		Code:    "000000",
		Message: "更新角色菜单成功",
	}, nil
}
