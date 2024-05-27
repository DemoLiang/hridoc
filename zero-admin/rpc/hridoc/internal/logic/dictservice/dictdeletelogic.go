package dictservicelogic

import (
	"context"
	"zero-admin/rpc/proto/sys"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DictDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:02
*/
type DictDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDeleteLogic {
	return &DictDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DictDelete 删除字典信息
func (l *DictDeleteLogic) DictDelete(in *sys.DictDeleteReq) (*sys.DictDeleteResp, error) {
	err := l.svcCtx.DictModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sys.DictDeleteResp{}, nil
}
