package syslogservicelogic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/proto/sys"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SysLogAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:08
*/
type SysLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogAddLogic {
	return &SysLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SysLogAdd 添加操作日志
func (l *SysLogAddLogic) SysLogAdd(in *sys.SysLogAddReq) (*sys.SysLogAddResp, error) {
	sysLog := &sysmodel.SysLog{
		UserName:       in.UserName,
		Operation:      in.Operation,
		Method:         in.Method,
		RequestParams:  in.RequestParams,
		ResponseParams: sql.NullString{String: in.ResponseParams, Valid: true},
		Time:           in.Time,
		Ip:             sql.NullString{String: in.Ip, Valid: true},
		OperationTime:  time.Now(),
	}
	if _, err := l.svcCtx.SysLogModel.Insert(l.ctx, sysLog); err != nil {
		return nil, err
	}

	return &sys.SysLogAddResp{}, nil
}
