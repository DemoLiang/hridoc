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

type GetCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryListLogic) GetCategoryList(req *types.CategoryListReq) (resp *types.CategoryListResp, err error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	var list []types.CategoryInfo
	var total int64

	keyword := strings.TrimSpace(req.Keyword)
	where := "1=1"
	args := []any{}
	if keyword != "" {
		where += " and (`name` like ? or `code` like ?)"
		args = append(args, "%"+keyword+"%", "%"+keyword+"%")
	}

	countQuery := fmt.Sprintf("select count(*) from `cert_category` where %s", where)
	err = l.svcCtx.DB.QueryRowCtx(l.ctx, &total, countQuery, args...)
	if err != nil {
		logx.Errorf("count category failed: %v", err)
		return &types.CategoryListResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	if total > 0 {
		listQuery := fmt.Sprintf("select id, name, code, description, status from `cert_category` where %s order by id desc limit ? offset ?", where)
		queryArgs := append(args, pageSize, (page-1)*pageSize)
		var rows []struct {
			Id          int64          `db:"id"`
			Name        string         `db:"name"`
			Code        string         `db:"code"`
			Description sql.NullString `db:"description"`
			Status      int64          `db:"status"`
		}
		err := l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &rows, listQuery, queryArgs...)
		if err != nil {
			logx.Errorf("query category list failed: %v", err)
			return &types.CategoryListResp{
				BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
			}, nil
		}

		for _, c := range rows {
			list = append(list, types.CategoryInfo{
				Id:          c.Id,
				Name:        c.Name,
				Code:        c.Code,
				Description: nullString(c.Description),
				Status:      c.Status,
			})
		}
	}

	return &types.CategoryListResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.CategoryListData{
			PageResp: types.PageResp{Total: total, Page: page},
			List:     list,
		},
	}, nil
}
