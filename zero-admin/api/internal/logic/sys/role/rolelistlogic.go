package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// RoleListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:39
*/
type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleListLogic {
	return RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RoleList 角色列表
func (l *RoleListLogic) RoleList(req types.ListRoleReq) (*types.ListRoleResp, error) {
	l.Logger.Infof("==========================req:%+v", req)
	resp, err := l.svcCtx.RoleService.RoleList(l.ctx, &sysclient.RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		Status:   req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询角色列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询角色失败")
	}
	logc.Infof(l.ctx, "---------------:%+v", resp)

	var list []*types.ListRoleData

	for _, role := range resp.List {
		list = append(list, &types.ListRoleData{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreateTime,
			LastUpdateBy:   role.LastUpdateBy,
			LastUpdateTime: role.LastUpdateTime,
			DelFlag:        role.DelFlag,
			Label:          role.Name,
			Value:          strconv.FormatInt(role.Id, 10),
			Status:         role.Status,
		})
	}
	logc.Infof(l.ctx, "list:%+v", list)

	return &types.ListRoleResp{
		Code:     "000000",
		Message:  "查询角色成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
