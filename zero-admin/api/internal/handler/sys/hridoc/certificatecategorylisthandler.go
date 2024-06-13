package hridoc

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin/api/internal/logic/sys/hridoc"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
)

func CertificateCategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListCertificateCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := hridoc.NewCertificateCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.CertificateCategoryList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
