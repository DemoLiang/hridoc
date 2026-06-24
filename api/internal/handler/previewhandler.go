// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/DemoLiang/hridoc/api/internal/logic"
	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func previewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		defer file.Close()

		excelBytes, err := io.ReadAll(file)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		var categoryCodes []string
		if codes := r.FormValue("categoryCodes"); codes != "" {
			categoryCodes = strings.Split(codes, ",")
		}

		req := types.PreviewReq{
			CategoryCodes: categoryCodes,
		}

		l := logic.NewPreviewLogic(r.Context(), svcCtx)
		resp, err := l.Preview(&req, excelBytes)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
