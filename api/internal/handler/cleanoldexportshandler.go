// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"github.com/DemoLiang/hridoc/api/internal/logic"
	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func cleanOldExportsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCleanOldExportsLogic(r.Context(), svcCtx)
		resp, err := l.CleanOldExports()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
