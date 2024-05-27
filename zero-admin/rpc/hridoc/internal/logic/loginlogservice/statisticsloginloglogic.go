package loginlogservicelogic

import (
	"context"

	"zero-admin/rpc/hridoc/internal/svc"
	"zero-admin/rpc/proto/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

// StatisticsLoginLogLogic
/*
Author: LiuFeiHua
Date: 2024/01/15 11:55
*/
type StatisticsLoginLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStatisticsLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsLoginLogLogic {
	return &StatisticsLoginLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// StatisticsLoginLog 统计后台用户登录---(查询当天登录人数（根据IP）,统计当前周登录人数（根据IP）,统计当前月登录人数（根据IP）)
func (l *StatisticsLoginLogLogic) StatisticsLoginLog(in *sys.StatisticsLoginLogReq) (*sys.StatisticsLoginLogResp, error) {
	//查询当天登录人数（根据IP）
	dayLoginCount, _ := l.svcCtx.LoginLogModel.StatisticsLoginLog(l.ctx, 1)
	//统计当前周登录人数（根据IP）
	weekLoginCount, _ := l.svcCtx.LoginLogModel.StatisticsLoginLog(l.ctx, 2)
	//统计当前月登录人数（根据IP）
	monthLoginCount, _ := l.svcCtx.LoginLogModel.StatisticsLoginLog(l.ctx, 3)
	return &sys.StatisticsLoginLogResp{
		CurrentDayLoginCount:   dayLoginCount,
		CurrentWeekLoginCount:  weekLoginCount,
		CurrentMonthLoginCount: monthLoginCount,
	}, nil
}
