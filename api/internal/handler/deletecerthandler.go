// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"github.com/DemoLiang/hridoc/api/internal/logic"
	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func deleteCertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteCertLogic(r.Context(), svcCtx)
		resp, err := l.DeleteCert(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
