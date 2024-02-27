// Code generated by goctl. DO NOT EDIT.

package hridoc

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	certificateFieldNames          = builder.RawFieldNames(&Certificate{})
	certificateRows                = strings.Join(certificateFieldNames, ",")
	certificateRowsExpectAutoSet   = strings.Join(stringx.Remove(certificateFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	certificateRowsWithPlaceHolder = strings.Join(stringx.Remove(certificateFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheCertificateIdPrefix = "cache:certificate:id:"
)

type (
	certificateModel interface {
		Insert(ctx context.Context, data *Certificate) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Certificate, error)
		Update(ctx context.Context, data *Certificate) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCertificateModel struct {
		sqlc.CachedConn
		table string
	}

	Certificate struct {
		CreatedAt sql.NullTime   `db:"created_at"`
		UpdatedAt sql.NullTime   `db:"updated_at"`
		DeletedAt sql.NullTime   `db:"deleted_at"`
		Id        int64          `db:"id"`      // 证书ID
		Name      string         `db:"name"`    // 证书名称
		Path      sql.NullString `db:"path"`    // 存储位置
		UserId    sql.NullInt64  `db:"user_id"` // 证书持有人
	}
)

func newCertificateModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultCertificateModel {
	return &defaultCertificateModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`certificate`",
	}
}

func (m *defaultCertificateModel) withSession(session sqlx.Session) *defaultCertificateModel {
	return &defaultCertificateModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`certificate`",
	}
}

func (m *defaultCertificateModel) Delete(ctx context.Context, id int64) error {
	certificateIdKey := fmt.Sprintf("%s%v", cacheCertificateIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, certificateIdKey)
	return err
}

func (m *defaultCertificateModel) FindOne(ctx context.Context, id int64) (*Certificate, error) {
	certificateIdKey := fmt.Sprintf("%s%v", cacheCertificateIdPrefix, id)
	var resp Certificate
	err := m.QueryRowCtx(ctx, &resp, certificateIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", certificateRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCertificateModel) Insert(ctx context.Context, data *Certificate) (sql.Result, error) {
	certificateIdKey := fmt.Sprintf("%s%v", cacheCertificateIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, certificateRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.Path, data.UserId)
	}, certificateIdKey)
	return ret, err
}

func (m *defaultCertificateModel) Update(ctx context.Context, data *Certificate) error {
	certificateIdKey := fmt.Sprintf("%s%v", cacheCertificateIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, certificateRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.Path, data.UserId, data.Id)
	}, certificateIdKey)
	return err
}

func (m *defaultCertificateModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheCertificateIdPrefix, primary)
}

func (m *defaultCertificateModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", certificateRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCertificateModel) tableName() string {
	return m.table
}
