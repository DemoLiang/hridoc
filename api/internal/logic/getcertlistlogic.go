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

type GetCertListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCertListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCertListLogic {
	return &GetCertListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCertListLogic) GetCertList(req *types.CertListReq) (resp *types.CertListResp, err error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	var list []types.CertificateInfo
	var total int64

	where := "1=1"
	args := []any{}
	if req.UserId > 0 {
		where += " and c.user_id = ?"
		args = append(args, req.UserId)
	}
	if req.CategoryId > 0 {
		where += " and c.category_id = ?"
		args = append(args, req.CategoryId)
	}
	keyword := strings.TrimSpace(req.Keyword)
	if keyword != "" {
		where += " and (c.name like ? or c.cert_no like ?)"
		args = append(args, "%"+keyword+"%", "%"+keyword+"%")
	}

	countQuery := fmt.Sprintf("select count(*) from certificate c where %s", where)
	err = l.svcCtx.DB.QueryRowCtx(l.ctx, &total, countQuery, args...)
	if err != nil {
		logx.Errorf("count certificate failed: %v", err)
		return &types.CertListResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	if total > 0 {
		listQuery := fmt.Sprintf(`select c.id, c.user_id, u.name as user_name, c.category_id, cc.name as category_name,
			c.name, c.cert_no, c.issuer, c.issue_date, c.expire_date, c.level, c.file_url, c.file_type, c.thumb_url, c.status
			from certificate c
			left join user u on c.user_id = u.id
			left join cert_category cc on c.category_id = cc.id
			where %s order by c.id desc limit ? offset ?`, where)
		queryArgs := append(args, pageSize, (page-1)*pageSize)

		var rows []struct {
			Id           int64          `db:"id"`
			UserId       int64          `db:"user_id"`
			UserName     sql.NullString `db:"user_name"`
			CategoryId   int64          `db:"category_id"`
			CategoryName sql.NullString `db:"category_name"`
			Name         string         `db:"name"`
			CertNo       sql.NullString `db:"cert_no"`
			Issuer       sql.NullString `db:"issuer"`
			IssueDate    sql.NullTime   `db:"issue_date"`
			ExpireDate   sql.NullTime   `db:"expire_date"`
			Level        sql.NullString `db:"level"`
			FileUrl      sql.NullString `db:"file_url"`
			FileType     sql.NullString `db:"file_type"`
			ThumbUrl     sql.NullString `db:"thumb_url"`
			Status       int64          `db:"status"`
		}

		err := l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &rows, listQuery, queryArgs...)
		if err != nil {
			logx.Errorf("query certificate list failed: %v", err)
			return &types.CertListResp{
				BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
			}, nil
		}

		for _, r := range rows {
			list = append(list, types.CertificateInfo{
				Id:           r.Id,
				UserId:       r.UserId,
				UserName:     nullString(r.UserName),
				CategoryId:   r.CategoryId,
				CategoryName: nullString(r.CategoryName),
				Name:         r.Name,
				CertNo:       nullString(r.CertNo),
				Issuer:       nullString(r.Issuer),
				IssueDate:    formatDate(r.IssueDate),
				ExpireDate:   formatDate(r.ExpireDate),
				Level:        nullString(r.Level),
				FileUrl:      nullString(r.FileUrl),
				FileType:     nullString(r.FileType),
				ThumbUrl:     nullString(r.ThumbUrl),
				Status:       r.Status,
			})
		}
	}

	return &types.CertListResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.CertListData{
			PageResp: types.PageResp{Total: total, Page: page},
			List:     list,
		},
	}, nil
}
