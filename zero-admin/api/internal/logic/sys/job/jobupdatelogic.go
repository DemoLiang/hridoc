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

// JobUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:19
*/
type JobUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) JobUpdateLogic {
	return JobUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// JobUpdate 更新岗位信息
func (l *JobUpdateLogic) JobUpdate(req types.UpdateJobReq) (*types.UpdateJobResp, error) {
	_, err := l.svcCtx.JobService.JobUpdate(l.ctx, &sys.JobUpdateReq{
		Id:           req.Id,
		JobName:      req.JobName,
		OrderNum:     req.OrderNum,
		Remarks:      req.Remarks,
		LastUpdateBy: l.ctx.Value("userName").(string),
		DelFlag:      req.DelFlag,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新岗位信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("删除岗位失败")
	}

	return &types.UpdateJobResp{
		Code:    "000000",
		Message: "删除岗位信息成功",
	}, nil
}
