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

type UpdateCertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCertLogic {
	return &UpdateCertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCertLogic) UpdateCert(req *types.UpdateCertReq) (resp *types.BaseResp, err error) {
	cert, err := l.svcCtx.CertificateModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.BaseResp{Code: errorx.ErrCertNotFound, Message: "证件不存在"}, nil
		}
		logx.Errorf("find certificate failed: %v", err)
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

	cert.CategoryId = req.CategoryId
	cert.Name = req.Name
	cert.CertNo = sql.NullString{String: req.CertNo, Valid: req.CertNo != ""}
	cert.Issuer = sql.NullString{String: req.Issuer, Valid: req.Issuer != ""}
	cert.IssueDate = parseDate(req.IssueDate)
	cert.ExpireDate = parseDate(req.ExpireDate)
	cert.Level = sql.NullString{String: req.Level, Valid: req.Level != ""}
	if req.FileUrl != "" {
		cert.FileUrl = req.FileUrl
	}
	if req.FileType != "" {
		cert.FileType = req.FileType
	}
	cert.Status = req.Status

	err = l.svcCtx.CertificateModel.Update(l.ctx, cert)
	if err != nil {
		logx.Errorf("update certificate failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	return &types.BaseResp{Code: 0, Message: "success"}, nil
}
