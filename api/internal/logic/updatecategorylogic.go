// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdateCategoryReq) (resp *types.BaseResp, err error) {
	cat, err := l.svcCtx.CertCategoryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.BaseResp{Code: errorx.ErrCategoryNotFound, Message: "证件类型不存在"}, nil
		}
		logx.Errorf("find category failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	if cat.Code != req.Code {
		_, err = l.svcCtx.CertCategoryModel.FindOneByCode(l.ctx, req.Code)
		if err == nil {
			return &types.BaseResp{Code: errorx.ErrIdCardExists, Message: "证件类型编码已存在"}, nil
		}
		if !errorx.IsNotFound(err) {
			logx.Errorf("query category code failed: %v", err)
			return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
		}
	}

	cat.Name = req.Name
	cat.Code = req.Code
	cat.Description = sql.NullString{String: req.Description, Valid: req.Description != ""}
	cat.Status = req.Status

	err = l.svcCtx.CertCategoryModel.Update(l.ctx, cat)
	if err != nil {
		logx.Errorf("update category failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	return &types.BaseResp{Code: 0, Message: "success"}, nil
}
