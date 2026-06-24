// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/DemoLiang/hridoc/api/internal/logic"
	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createExportTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		var categoryCodes []string
		if codes := r.FormValue("categoryCodes"); codes != "" {
			categoryCodes = strings.Split(codes, ",")
		}

		var userIds []int64
		if idsStr := r.FormValue("userIds"); idsStr != "" {
			for _, s := range strings.Split(idsStr, ",") {
				if id, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64); err == nil {
					userIds = append(userIds, id)
				}
			}
		}

		var excelBytes []byte
		file, _, err := r.FormFile("file")
		if err == nil {
			defer file.Close()
			excelBytes, err = io.ReadAll(file)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
		}

		req := types.ExportTaskReq{
			CategoryCodes: categoryCodes,
			WatermarkText: r.FormValue("watermarkText"),
			WatermarkMode: r.FormValue("watermarkMode"),
			UserIds:       userIds,
		}
		if opacity, err := strconv.ParseFloat(r.FormValue("watermarkOpacity"), 64); err == nil {
			req.WatermarkOpacity = opacity
		}
		if fontSize, err := strconv.Atoi(r.FormValue("watermarkFontSize")); err == nil {
			req.WatermarkFontSize = fontSize
		}
		req.WatermarkColor = r.FormValue("watermarkColor")

		l := logic.NewCreateExportTaskLogic(r.Context(), svcCtx)
		resp, err := l.CreateExportTask(&req, excelBytes)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
