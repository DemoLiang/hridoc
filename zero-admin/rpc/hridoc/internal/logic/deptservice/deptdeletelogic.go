package deptservicelogic

import (
	"context"
	"zero-admin/rpc/proto/sys"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeptDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:00
*/
type DeptDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptDeleteLogic {
	return &DeptDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeptDelete 删除部门信息
func (l *DeptDeleteLogic) DeptDelete(in *sys.DeptDeleteReq) (*sys.DeptDeleteResp, error) {
	err := l.svcCtx.DeptModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sys.DeptDeleteResp{}, nil
}
