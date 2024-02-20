package handler

import (
	"net/http"

	"zero-admin/api/internal/logic/member/statistics"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MemberStatisticsInfoUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateMemberStatisticsInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMemberStatisticsInfoUpdateLogic(r.Context(), ctx)
		resp, err := l.MemberStatisticsInfoUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
