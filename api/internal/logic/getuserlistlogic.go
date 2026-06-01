// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	var users []types.UserInfo
	var total int64

	keyword := strings.TrimSpace(req.Keyword)
	where := "1=1"
	args := []any{}
	if keyword != "" {
		where += " and (`name` like ? or `phone` like ? or `id_card` like ?)"
		args = append(args, "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	// count
	countQuery := fmt.Sprintf("select count(*) from `user` where %s", where)
	err = l.svcCtx.DB.QueryRowCtx(l.ctx, &total, countQuery, args...)
	if err != nil {
		logx.Errorf("count user failed: %v", err)
		return &types.UserListResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	if total > 0 {
		listQuery := fmt.Sprintf("select id, name, phone, email, id_card, education, role, status from `user` where %s order by id desc limit ? offset ?", where)
		queryArgs := append(args, pageSize, (page-1)*pageSize)
		var rows []struct {
			Id        int64          `db:"id"`
			Name      string         `db:"name"`
			Phone     sql.NullString `db:"phone"`
			Email     sql.NullString `db:"email"`
			IdCard    string         `db:"id_card"`
			Education sql.NullString `db:"education"`
			Role      int64          `db:"role"`
			Status    int64          `db:"status"`
		}
		err := l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &rows, listQuery, queryArgs...)
		if err != nil {
			logx.Errorf("query user list failed: %v", err)
			return &types.UserListResp{
				BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
			}, nil
		}

		for _, u := range rows {
			users = append(users, types.UserInfo{
				Id:        u.Id,
				Name:      u.Name,
				Phone:     nullString(u.Phone),
				Email:     nullString(u.Email),
				IdCard:    u.IdCard,
				Education: nullString(u.Education),
				Role:      u.Role,
				Status:    u.Status,
			})
		}
	}

	return &types.UserListResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.UserListData{
			PageResp: types.PageResp{Total: total, Page: page},
			List:     users,
		},
	}, nil
}
