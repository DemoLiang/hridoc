package loginlogservicelogic

import (
	"context"
	"time"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/proto/sys"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:07
*/
type LoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogAddLogic {
	return &LoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginLogAdd 添加登录日志
func (l *LoginLogAddLogic) LoginLogAdd(in *sys.LoginLogAddReq) (*sys.LoginLogAddResp, error) {
	loginLog := &sysmodel.SysLoginLog{
		UserName:   in.UserName,
		Status:     in.Status,
		Ip:         in.Ip,
		CreateTime: time.Now(),
		CreateBy:   in.CreateBy,
	}
	if _, err := l.svcCtx.LoginLogModel.Insert(l.ctx, loginLog); err != nil {
		return nil, err
	}

	return &sys.LoginLogAddResp{}, nil
}
