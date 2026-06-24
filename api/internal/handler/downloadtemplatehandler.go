// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"fmt"
	"net/http"

	"github.com/DemoLiang/hridoc/api/internal/logic"
	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func downloadTemplateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDownloadTemplateLogic(r.Context(), svcCtx)
		resp, err := l.DownloadTemplate()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.Header().Set("X-Download-Url", "/api/excel/template/download")
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

func downloadTemplateFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDownloadTemplateLogic(r.Context(), svcCtx)
		data, err := l.DownloadTemplateFile()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=\"import_template.xlsx\"")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
		_, _ = w.Write(data)
	}
}
