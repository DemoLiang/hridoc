// Code generated by goctl. DO NOT EDIT.

package umsmodel

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
	umsMemberReadHistoryFieldNames          = builder.RawFieldNames(&UmsMemberReadHistory{})
	umsMemberReadHistoryRows                = strings.Join(umsMemberReadHistoryFieldNames, ",")
	umsMemberReadHistoryRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberReadHistoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	umsMemberReadHistoryRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberReadHistoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	umsMemberReadHistoryModel interface {
		Insert(ctx context.Context, data *UmsMemberReadHistory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UmsMemberReadHistory, error)
		Update(ctx context.Context, data *UmsMemberReadHistory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsMemberReadHistoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberReadHistory struct {
		Id              int64          `db:"id"`                // 编号
		MemberId        int64          `db:"member_id"`         // 会员id
		MemberNickName  string         `db:"member_nick_name"`  // 会员姓名
		MemberIcon      string         `db:"member_icon"`       // 会员头像
		ProductId       int64          `db:"product_id"`        // 商品id
		ProductName     string         `db:"product_name"`      // 商品名称
		ProductPic      string         `db:"product_pic"`       // 商品图片
		ProductSubTitle sql.NullString `db:"product_sub_title"` // 商品标题
		ProductPrice    float64        `db:"product_price"`     // 商品价格
		CreateTime      time.Time      `db:"create_time"`       // 浏览时间
	}
)

func newUmsMemberReadHistoryModel(conn sqlx.SqlConn) *defaultUmsMemberReadHistoryModel {
	return &defaultUmsMemberReadHistoryModel{
		conn:  conn,
		table: "`ums_member_read_history`",
	}
}

func (m *defaultUmsMemberReadHistoryModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUmsMemberReadHistoryModel) FindOne(ctx context.Context, id int64) (*UmsMemberReadHistory, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberReadHistoryRows, m.table)
	var resp UmsMemberReadHistory
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

func (m *defaultUmsMemberReadHistoryModel) Insert(ctx context.Context, data *UmsMemberReadHistory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, umsMemberReadHistoryRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.MemberId, data.MemberNickName, data.MemberIcon, data.ProductId, data.ProductName, data.ProductPic, data.ProductSubTitle, data.ProductPrice)
	return ret, err
}

func (m *defaultUmsMemberReadHistoryModel) Update(ctx context.Context, data *UmsMemberReadHistory) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberReadHistoryRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.MemberId, data.MemberNickName, data.MemberIcon, data.ProductId, data.ProductName, data.ProductPic, data.ProductSubTitle, data.ProductPrice, data.Id)
	return err
}

func (m *defaultUmsMemberReadHistoryModel) tableName() string {
	return m.table
}
