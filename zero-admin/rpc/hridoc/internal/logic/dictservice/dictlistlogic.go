package dictservicelogic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/proto/sys"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DictListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:03
*/
type DictListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictListLogic {
	return &DictListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DictList 字典列表
func (l *DictListLogic) DictList(in *sys.DictListReq) (*sys.DictListResp, error) {
	all, err := l.svcCtx.DictModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.DictModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询字典列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
	var list []*sys.DictListData
	for _, dict := range *all {
		list = append(list, &sys.DictListData{
			Id:             dict.Id,
			Value:          dict.Value,
			Label:          dict.Label,
			Type:           dict.Type,
			Description:    dict.Description,
			Sort:           int64(dict.Sort),
			Remarks:        dict.Remarks.String,
			CreateBy:       dict.CreateBy,
			CreateTime:     dict.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   dict.UpdateBy.String,
			LastUpdateTime: dict.UpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        dict.DelFlag,
		})
	}

	logc.Infof(l.ctx, "查询字典列表信息,参数：%+v,响应：%+v", in, list)
	return &sys.DictListResp{
		Total: count,
		List:  list,
	}, nil

}
