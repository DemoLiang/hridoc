// Code generated by goctl. DO NOT EDIT.

package umsmodel

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
	umsMemberTaskFieldNames          = builder.RawFieldNames(&UmsMemberTask{})
	umsMemberTaskRows                = strings.Join(umsMemberTaskFieldNames, ",")
	umsMemberTaskRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	umsMemberTaskRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	umsMemberTaskModel interface {
		Insert(ctx context.Context, data *UmsMemberTask) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UmsMemberTask, error)
		Update(ctx context.Context, data *UmsMemberTask) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsMemberTaskModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberTask struct {
		Id           int64  `db:"id"`
		Name         string `db:"name"`
		Growth       int64  `db:"growth"`       // 赠送成长值
		Intergration int64  `db:"intergration"` // 赠送积分
		Type         int64  `db:"type"`         // 任务类型：0->新手任务；1->日常任务
	}
)

func newUmsMemberTaskModel(conn sqlx.SqlConn) *defaultUmsMemberTaskModel {
	return &defaultUmsMemberTaskModel{
		conn:  conn,
		table: "`ums_member_task`",
	}
}

func (m *defaultUmsMemberTaskModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUmsMemberTaskModel) FindOne(ctx context.Context, id int64) (*UmsMemberTask, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberTaskRows, m.table)
	var resp UmsMemberTask
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

func (m *defaultUmsMemberTaskModel) Insert(ctx context.Context, data *UmsMemberTask) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, umsMemberTaskRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Growth, data.Intergration, data.Type)
	return ret, err
}

func (m *defaultUmsMemberTaskModel) Update(ctx context.Context, data *UmsMemberTask) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberTaskRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Growth, data.Intergration, data.Type, data.Id)
	return err
}

func (m *defaultUmsMemberTaskModel) tableName() string {
	return m.table
}
