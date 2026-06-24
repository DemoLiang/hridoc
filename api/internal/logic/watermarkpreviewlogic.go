// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/pkg/watermark"
	"github.com/zeromicro/go-zero/core/logx"
)

type WatermarkPreviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWatermarkPreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WatermarkPreviewLogic {
	return &WatermarkPreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WatermarkPreviewLogic) WatermarkPreview(opts *watermark.Options) ([]byte, error) {
	return watermark.PreviewSample(opts)
}
