// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryDetailLogic {
	return &GetCategoryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryDetailLogic) GetCategoryDetail(req *types.UserDetailReq) (resp *types.CategoryListResp, err error) {
	cat, err := l.svcCtx.CertCategoryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.CategoryListResp{
				BaseResp: types.BaseResp{Code: errorx.ErrCategoryNotFound, Message: "证件类型不存在"},
			}, nil
		}
		logx.Errorf("get category detail failed: %v", err)
		return &types.CategoryListResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	return &types.CategoryListResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.CategoryListData{
			PageResp: types.PageResp{Total: 1, Page: 1},
			List: []types.CategoryInfo{{
				Id:          cat.Id,
				Name:        cat.Name,
				Code:        cat.Code,
				Description: nullString(cat.Description),
				Status:      cat.Status,
			}},
		},
	}, nil
}
