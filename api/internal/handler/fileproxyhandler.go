// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/DemoLiang/hridoc/api/internal/logic"
	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func fileProxyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("url 参数不能为空"))
			return
		}

		l := logic.NewFileProxyLogic(r.Context(), svcCtx)
		reader, filename, contentType, err := l.ProxyFile(url)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		defer reader.Close()

		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", filename))
		_, _ = io.Copy(w, reader)
	}
}
