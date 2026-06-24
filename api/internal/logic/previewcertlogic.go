// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"strings"
	"time"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type PreviewCertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPreviewCertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreviewCertLogic {
	return &PreviewCertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PreviewCertLogic) extractObjectName(fileUrl string) string {
	parts := strings.Split(fileUrl, "/")
	if len(parts) < 2 {
		return ""
	}
	return parts[len(parts)-2] + "/" + parts[len(parts)-1]
}

func (l *PreviewCertLogic) PreviewCert(req *types.CertPreviewReq) (resp *types.PreviewUrlResp, err error) {
	if req.FileUrl == "" {
		return &types.PreviewUrlResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidParam, Message: "缺少文件地址"},
		}, nil
	}
	objectName := l.extractObjectName(req.FileUrl)
	if objectName == "" {
		return &types.PreviewUrlResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidParam, Message: "无效的文件地址"},
		}, nil
	}
	if l.svcCtx.MinIO == nil {
		return &types.PreviewUrlResp{
			BaseResp: types.BaseResp{Code: errorx.ErrMinIOFailed, Message: "MinIO 未就绪"},
		}, nil
	}
	url, er := l.svcCtx.MinIO.PresignedGetURL(l.ctx, objectName, time.Duration(1)*time.Hour)
	if er != nil {
		logx.Errorf("generate presigned url failed: %v", er)
		return &types.PreviewUrlResp{
			BaseResp: types.BaseResp{Code: errorx.ErrMinIOFailed, Message: "生成预览链接失败"},
		}, nil
	}
	return &types.PreviewUrlResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data:     types.PreviewUrlData{Url: url},
	}, nil
}
