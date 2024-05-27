package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/proto/sys"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ConfigDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:11
*/
type ConfigDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigDeleteLogic {
	return ConfigDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ConfigDelete 删除配置
func (l *ConfigDeleteLogic) ConfigDelete(req types.DeleteConfigReq) (*types.DeleteConfigResp, error) {
	_, err := l.svcCtx.ConfigService.ConfigDelete(l.ctx, &sys.ConfigDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("删除参数配置失败")
	}

	return &types.DeleteConfigResp{}, nil
}
