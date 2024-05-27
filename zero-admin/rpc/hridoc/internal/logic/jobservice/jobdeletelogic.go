package jobservicelogic

import (
	"context"
	"zero-admin/rpc/proto/sys"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// JobDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:04
*/
type JobDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobDeleteLogic {
	return &JobDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// JobDelete 删除岗位
func (l *JobDeleteLogic) JobDelete(in *sys.JobDeleteReq) (*sys.JobDeleteResp, error) {
	err := l.svcCtx.JobModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sys.JobDeleteResp{}, nil
}
