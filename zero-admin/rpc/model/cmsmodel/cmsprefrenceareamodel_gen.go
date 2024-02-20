// Code generated by goctl. DO NOT EDIT.

package cmsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	cmsPrefrenceAreaFieldNames          = builder.RawFieldNames(&CmsPrefrenceArea{})
	cmsPrefrenceAreaRows                = strings.Join(cmsPrefrenceAreaFieldNames, ",")
	cmsPrefrenceAreaRowsExpectAutoSet   = strings.Join(stringx.Remove(cmsPrefrenceAreaFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	cmsPrefrenceAreaRowsWithPlaceHolder = strings.Join(stringx.Remove(cmsPrefrenceAreaFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	cmsPrefrenceAreaModel interface {
		Insert(ctx context.Context, data *CmsPrefrenceArea) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*CmsPrefrenceArea, error)
		Update(ctx context.Context, data *CmsPrefrenceArea) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCmsPrefrenceAreaModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CmsPrefrenceArea struct {
		Id         int64  `db:"id"`
		Name       string `db:"name"`
		SubTitle   string `db:"sub_title"`
		Pic        string `db:"pic"` // 展示图片
		Sort       int64  `db:"sort"`
		ShowStatus int64  `db:"show_status"`
	}
)

func newCmsPrefrenceAreaModel(conn sqlx.SqlConn) *defaultCmsPrefrenceAreaModel {
	return &defaultCmsPrefrenceAreaModel{
		conn:  conn,
		table: "`cms_prefrence_area`",
	}
}

func (m *defaultCmsPrefrenceAreaModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCmsPrefrenceAreaModel) FindOne(ctx context.Context, id int64) (*CmsPrefrenceArea, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cmsPrefrenceAreaRows, m.table)
	var resp CmsPrefrenceArea
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCmsPrefrenceAreaModel) Insert(ctx context.Context, data *CmsPrefrenceArea) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, cmsPrefrenceAreaRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.SubTitle, data.Pic, data.Sort, data.ShowStatus)
	return ret, err
}

func (m *defaultCmsPrefrenceAreaModel) Update(ctx context.Context, data *CmsPrefrenceArea) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, cmsPrefrenceAreaRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.SubTitle, data.Pic, data.Sort, data.ShowStatus, data.Id)
	return err
}

func (m *defaultCmsPrefrenceAreaModel) tableName() string {
	return m.table
}
