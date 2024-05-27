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

// RoleAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:54
*/
type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleAdd 新增角色
func (l *RoleAddLogic) RoleAdd(in *sys.RoleAddReq) (*sys.RoleAddResp, error) {
	role := &sysmodel.SysRole{
		Name:     in.Name,
		Remark:   sql.NullString{String: in.Remark, Valid: true},
		CreateBy: in.CreateBy,
		DelFlag:  0,
		Status:   in.Status,
	}
	_, err := l.svcCtx.RoleModel.Insert(l.ctx, role)

	if err != nil {
		logc.Errorf(l.ctx, "新增角色信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}
	return &sys.RoleAddResp{}, nil
}
