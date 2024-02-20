// Code generated by goctl. DO NOT EDIT.

package smsmodel

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
	smsCouponProductCategoryRelationFieldNames          = builder.RawFieldNames(&SmsCouponProductCategoryRelation{})
	smsCouponProductCategoryRelationRows                = strings.Join(smsCouponProductCategoryRelationFieldNames, ",")
	smsCouponProductCategoryRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(smsCouponProductCategoryRelationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	smsCouponProductCategoryRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(smsCouponProductCategoryRelationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	smsCouponProductCategoryRelationModel interface {
		Insert(ctx context.Context, data *SmsCouponProductCategoryRelation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SmsCouponProductCategoryRelation, error)
		Update(ctx context.Context, data *SmsCouponProductCategoryRelation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSmsCouponProductCategoryRelationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsCouponProductCategoryRelation struct {
		Id                  int64  `db:"id"`
		CouponId            int64  `db:"coupon_id"`
		ProductCategoryId   int64  `db:"product_category_id"`
		ProductCategoryName string `db:"product_category_name"` // 产品分类名称
		ParentCategoryName  string `db:"parent_category_name"`  // 父分类名称
	}
)

func newSmsCouponProductCategoryRelationModel(conn sqlx.SqlConn) *defaultSmsCouponProductCategoryRelationModel {
	return &defaultSmsCouponProductCategoryRelationModel{
		conn:  conn,
		table: "`sms_coupon_product_category_relation`",
	}
}

func (m *defaultSmsCouponProductCategoryRelationModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSmsCouponProductCategoryRelationModel) FindOne(ctx context.Context, id int64) (*SmsCouponProductCategoryRelation, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", smsCouponProductCategoryRelationRows, m.table)
	var resp SmsCouponProductCategoryRelation
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

func (m *defaultSmsCouponProductCategoryRelationModel) Insert(ctx context.Context, data *SmsCouponProductCategoryRelation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, smsCouponProductCategoryRelationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CouponId, data.ProductCategoryId, data.ProductCategoryName, data.ParentCategoryName)
	return ret, err
}

func (m *defaultSmsCouponProductCategoryRelationModel) Update(ctx context.Context, data *SmsCouponProductCategoryRelation) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, smsCouponProductCategoryRelationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.CouponId, data.ProductCategoryId, data.ProductCategoryName, data.ParentCategoryName, data.Id)
	return err
}

func (m *defaultSmsCouponProductCategoryRelationModel) tableName() string {
	return m.table
}
