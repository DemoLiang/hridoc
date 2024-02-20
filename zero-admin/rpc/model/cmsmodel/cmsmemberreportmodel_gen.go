// Code generated by goctl. DO NOT EDIT.

package cmsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	cmsMemberReportFieldNames          = builder.RawFieldNames(&CmsMemberReport{})
	cmsMemberReportRows                = strings.Join(cmsMemberReportFieldNames, ",")
	cmsMemberReportRowsExpectAutoSet   = strings.Join(stringx.Remove(cmsMemberReportFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	cmsMemberReportRowsWithPlaceHolder = strings.Join(stringx.Remove(cmsMemberReportFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	cmsMemberReportModel interface {
		Insert(ctx context.Context, data *CmsMemberReport) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*CmsMemberReport, error)
		Update(ctx context.Context, data *CmsMemberReport) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCmsMemberReportModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CmsMemberReport struct {
		Id               int64          `db:"id"`
		ReportType       int64          `db:"report_type"`        // 举报类型：0->商品评价；1->话题内容；2->用户评论
		ReportMemberName string         `db:"report_member_name"` // 举报人
		CreateTime       time.Time      `db:"create_time"`
		ReportObject     string         `db:"report_object"`
		ReportStatus     int64          `db:"report_status"` // 举报状态：0->未处理；1->已处理
		HandleStatus     int64          `db:"handle_status"` // 处理结果：0->无效；1->有效；2->恶意
		Note             sql.NullString `db:"note"`
	}
)

func newCmsMemberReportModel(conn sqlx.SqlConn) *defaultCmsMemberReportModel {
	return &defaultCmsMemberReportModel{
		conn:  conn,
		table: "`cms_member_report`",
	}
}

func (m *defaultCmsMemberReportModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCmsMemberReportModel) FindOne(ctx context.Context, id int64) (*CmsMemberReport, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cmsMemberReportRows, m.table)
	var resp CmsMemberReport
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

func (m *defaultCmsMemberReportModel) Insert(ctx context.Context, data *CmsMemberReport) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, cmsMemberReportRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.ReportType, data.ReportMemberName, data.ReportObject, data.ReportStatus, data.HandleStatus, data.Note)
	return ret, err
}

func (m *defaultCmsMemberReportModel) Update(ctx context.Context, data *CmsMemberReport) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, cmsMemberReportRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ReportType, data.ReportMemberName, data.ReportObject, data.ReportStatus, data.HandleStatus, data.Note, data.Id)
	return err
}

func (m *defaultCmsMemberReportModel) tableName() string {
	return m.table
}
