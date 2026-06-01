// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/model"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCategoryLogic) AddCategory(req *types.AddCategoryReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.CertCategoryModel.FindOneByCode(l.ctx, req.Code)
	if err == nil {
		return &types.BaseResp{Code: errorx.ErrIdCardExists, Message: "证件类型编码已存在"}, nil
	}
	if !errorx.IsNotFound(err) {
		logx.Errorf("query category failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	_, err = l.svcCtx.CertCategoryModel.Insert(l.ctx, &model.CertCategory{
		Name:        req.Name,
		Code:        req.Code,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Status:      1,
	})
	if err != nil {
		logx.Errorf("insert category failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	return &types.BaseResp{Code: 0, Message: "success"}, nil
}
