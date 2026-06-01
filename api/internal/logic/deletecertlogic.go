// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"strings"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCertLogic {
	return &DeleteCertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func extractObjectName(urlStr string) string {
	if idx := strings.Index(urlStr, "?"); idx != -1 {
		urlStr = urlStr[:idx]
	}
	bucket := "hridoc"
	idx := strings.LastIndex(urlStr, "/"+bucket+"/")
	if idx == -1 {
		parts := strings.Split(urlStr, "/")
		if len(parts) > 0 {
			return parts[len(parts)-1]
		}
		return ""
	}
	return urlStr[idx+len(bucket)+2:]
}

func (l *DeleteCertLogic) DeleteCert(req *types.DeleteReq) (resp *types.BaseResp, err error) {
	cert, err := l.svcCtx.CertificateModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.BaseResp{Code: errorx.ErrCertNotFound, Message: "证件不存在"}, nil
		}
		logx.Errorf("find certificate failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	if l.svcCtx.MinIO != nil {
		if cert.FileUrl != "" {
			objectName := extractObjectName(cert.FileUrl)
			if objectName != "" {
				_ = l.svcCtx.MinIO.RemoveObject(l.ctx, objectName)
			}
		}
		if cert.ThumbUrl.Valid && cert.ThumbUrl.String != "" {
			thumbName := extractObjectName(cert.ThumbUrl.String)
			if thumbName != "" {
				_ = l.svcCtx.MinIO.RemoveObject(l.ctx, thumbName)
			}
		}
	}

	err = l.svcCtx.CertificateModel.Delete(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("delete certificate failed: %v", err)
		return &types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}, nil
	}

	return &types.BaseResp{Code: 0, Message: "success"}, nil
}
