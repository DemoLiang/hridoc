package dictservicelogic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/proto/sys"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DictAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:02
*/
type DictAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictAddLogic {
	return &DictAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DictAdd 添加字典信息
func (l *DictAddLogic) DictAdd(in *sys.DictAddReq) (*sys.DictAddResp, error) {
	dict := &sysmodel.SysDict{
		Value:       in.Value,
		Label:       in.Label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        float64(in.Sort),
		CreateBy:    in.CreateBy,
		Remarks:     sql.NullString{String: in.Remarks, Valid: true},
		DelFlag:     in.DelFlag,
	}
	if _, err := l.svcCtx.DictModel.Insert(l.ctx, dict); err != nil {
		return nil, err
	}

	return &sys.DictAddResp{}, nil
}
