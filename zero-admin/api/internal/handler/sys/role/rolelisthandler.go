package role

import (
	"net/http"
	"zero-admin/api/internal/logic/sys/role"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RoleListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewRoleListLogic(r.Context(), ctx)
		resp, err := l.RoleList(req)
		if err != nil {
			l.Logger.Infof("-------------查询角色列表异常:%s", err.Error())
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
