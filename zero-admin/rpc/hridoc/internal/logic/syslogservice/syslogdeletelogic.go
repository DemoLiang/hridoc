package syslogservicelogic

import (
	"context"
	"zero-admin/rpc/proto/sys"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SysLogDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:08
*/
type SysLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogDeleteLogic {
	return &SysLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SysLogDelete 删除操作日志
func (l *SysLogDeleteLogic) SysLogDelete(in *sys.SysLogDeleteReq) (*sys.SysLogDeleteResp, error) {
	err := l.svcCtx.SysLogModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sys.SysLogDeleteResp{}, nil
}
