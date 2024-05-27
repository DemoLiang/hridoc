package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/proto/sys"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeptUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
type DeptUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptUpdateLogic {
	return DeptUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeptUpdate 更新部门信息
func (l *DeptUpdateLogic) DeptUpdate(req types.UpdateDeptReq) (*types.UpdateDeptResp, error) {
	_, err := l.svcCtx.DeptService.DeptUpdate(l.ctx, &sys.DeptUpdateReq{
		Id:           req.Id,
		Name:         req.Name,
		ParentId:     req.ParentId,
		OrderNum:     req.OrderNum,
		LastUpdateBy: l.ctx.Value("userName").(string),
		ParentIds:    req.ParentIds,
		DelFlag:      req.DelFlag,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新机构信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新机构失败")
	}

	return &types.UpdateDeptResp{
		Code:    "000000",
		Message: "更新机构成功",
	}, nil
}
