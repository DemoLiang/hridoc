package roleservicelogic

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/proto/sys"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// RoleUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:57
*/
type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleUpdate 更新角色(id为1的是系统预留超级管理员角色,不能更新)
func (l *RoleUpdateLogic) RoleUpdate(in *sys.RoleUpdateReq) (*sys.RoleUpdateResp, error) {

	//id为1的是系统预留超级管理员角色,不用关联
	if in.Id == 1 {
		return &sys.RoleUpdateResp{}, nil
	}

	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		logc.Errorf(l.ctx, "查询角色信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	sysRole := &sysmodel.SysRole{
		Id:         in.Id,
		Name:       in.Name,
		Remark:     sql.NullString{String: in.Remark, Valid: true},
		CreateBy:   role.CreateBy,
		CreateTime: role.CreateTime,
		UpdateBy:   sql.NullString{String: in.LastUpdateBy, Valid: true},
		DelFlag:    0,
		Status:     in.Status,
	}
	if err1 := l.svcCtx.RoleModel.Update(l.ctx, sysRole); err1 != nil {
		logc.Errorf(l.ctx, "更新角色信息失败,参数:%+v,异常:%s", in, err1.Error())
		return nil, err1
	}

	return &sys.RoleUpdateResp{}, nil
}
