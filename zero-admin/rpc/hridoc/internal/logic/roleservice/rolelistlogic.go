package roleservicelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/hridoc/internal/svc"
	"zero-admin/rpc/proto/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

// RoleListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 16:06
*/
type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleList 角色列表
func (l *RoleListLogic) RoleList(in *sys.RoleListReq) (*sys.RoleListResp, error) {
	all, err := l.svcCtx.RoleModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.RoleModel.Count(l.ctx, in)

	if err != nil {
		logc.Errorf(l.ctx, "查询角色列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*sys.RoleListData
	for _, role := range *all {
		list = append(list, &sys.RoleListData{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark.String,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   role.UpdateBy.String,
			LastUpdateTime: role.UpdateTime.Time.Format("2006-01-02 15:04:05"),
			DelFlag:        role.DelFlag,
			Status:         role.Status,
		})
	}

	logc.Infof(l.ctx, "查询角色列表信息,参数：%+v,响应：%+v", in, list)
	return &sys.RoleListResp{
		Total: count,
		List:  list,
	}, nil

}
