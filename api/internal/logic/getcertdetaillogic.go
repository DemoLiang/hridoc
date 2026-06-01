// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCertDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCertDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCertDetailLogic {
	return &GetCertDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCertDetailLogic) GetCertDetail(req *types.CertDetailReq) (resp *types.CertListResp, err error) {
	query := `select c.id, c.user_id, u.name as user_name, c.category_id, cc.name as category_name,
		c.name, c.cert_no, c.issuer, c.issue_date, c.expire_date, c.level, c.file_url, c.file_type, c.thumb_url, c.status
		from certificate c
		left join user u on c.user_id = u.id
		left join cert_category cc on c.category_id = cc.id
		where c.id = ? limit 1`

	var row struct {
		Id           int64          `db:"id"`
		UserId       int64          `db:"user_id"`
		UserName     string         `db:"user_name"`
		CategoryId   int64          `db:"category_id"`
		CategoryName string         `db:"category_name"`
		Name         string         `db:"name"`
		CertNo       sql.NullString `db:"cert_no"`
		Issuer       sql.NullString `db:"issuer"`
		IssueDate    sql.NullTime   `db:"issue_date"`
		ExpireDate   sql.NullTime   `db:"expire_date"`
		Level        sql.NullString `db:"level"`
		FileUrl      string         `db:"file_url"`
		FileType     string         `db:"file_type"`
		ThumbUrl     sql.NullString `db:"thumb_url"`
		Status       int64          `db:"status"`
	}

	err = l.svcCtx.DB.QueryRowCtx(l.ctx, &row, query, req.Id)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.CertListResp{
				BaseResp: types.BaseResp{Code: errorx.ErrCertNotFound, Message: "证件不存在"},
			}, nil
		}
		logx.Errorf("get cert detail failed: %v", err)
		return &types.CertListResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	return &types.CertListResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.CertListData{
			PageResp: types.PageResp{Total: 1, Page: 1},
			List: []types.CertificateInfo{{
				Id:           row.Id,
				UserId:       row.UserId,
				UserName:     row.UserName,
				CategoryId:   row.CategoryId,
				CategoryName: row.CategoryName,
				Name:         row.Name,
				CertNo:       nullString(row.CertNo),
				Issuer:       nullString(row.Issuer),
				IssueDate:    formatDate(row.IssueDate),
				ExpireDate:   formatDate(row.ExpireDate),
				Level:        nullString(row.Level),
				FileUrl:      row.FileUrl,
				FileType:     row.FileType,
				ThumbUrl:     nullString(row.ThumbUrl),
				Status:       row.Status,
			}},
		},
	}, nil
}
