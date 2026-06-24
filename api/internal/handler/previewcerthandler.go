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

func previewCertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CertPreviewReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// Presigned URL 有效期内支持重复访问，直接复用它亦可
		// 后端帮助浏览器绕过同源/CORS 也可以新增一个 proxy 接口，这里先沿用 presigned 方案
		l := logic.NewPreviewCertLogic(r.Context(), svcCtx)
		resp, err := l.PreviewCert(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
