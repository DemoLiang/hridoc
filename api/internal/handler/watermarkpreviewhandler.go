// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/DemoLiang/hridoc/api/internal/logic"
	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/pkg/watermark"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func watermarkPreviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		text := r.FormValue("watermarkText")
		if text == "" {
			httpx.Error(w, errors.New("水印文字不能为空"))
			return
		}

		opts := &watermark.Options{
			Text:  text,
			Mode:  r.FormValue("watermarkMode"),
			Color: r.FormValue("watermarkColor"),
		}
		if opacity, err := strconv.ParseFloat(r.FormValue("watermarkOpacity"), 64); err == nil {
			opts.Opacity = opacity
		}
		if fontSize, err := strconv.Atoi(r.FormValue("watermarkFontSize")); err == nil {
			opts.FontSize = fontSize
		}

		l := logic.NewWatermarkPreviewLogic(r.Context(), svcCtx)
		data, err := l.WatermarkPreview(opts)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Disposition", "inline; filename=\"watermark_preview.jpg\"")
		_, _ = w.Write(data)
	}
}
