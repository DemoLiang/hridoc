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

type AddCertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCertLogic {
	return &AddCertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCertLogic) AddCert(req *types.AddCertReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.BaseResp{Code: errorx.ErrUserNotFound, Message: "用户不存在"}, nil
		}
		logx.Errorf("find user failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	_, err = l.svcCtx.CertCategoryModel.FindOne(l.ctx, req.CategoryId)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.BaseResp{Code: errorx.ErrCategoryNotFound, Message: "证件类型不存在"}, nil
		}
		logx.Errorf("find category failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	cert := &model.Certificate{
		UserId:     req.UserId,
		CategoryId: req.CategoryId,
		Name:       req.Name,
		CertNo:     sql.NullString{String: req.CertNo, Valid: req.CertNo != ""},
		Issuer:     sql.NullString{String: req.Issuer, Valid: req.Issuer != ""},
		IssueDate:  parseDate(req.IssueDate),
		ExpireDate: parseDate(req.ExpireDate),
		Level:      sql.NullString{String: req.Level, Valid: req.Level != ""},
		FileUrl:    req.FileUrl,
		FileType:   req.FileType,
		Status:     1,
	}

	_, err = l.svcCtx.CertificateModel.Insert(l.ctx, cert)
	if err != nil {
		logx.Errorf("insert certificate failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	return &types.BaseResp{Code: 0, Message: "success"}, nil
}
