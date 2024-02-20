package umsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-admin/rpc/ums/umsclient"
)

var _ UmsMemberProductCollectionModel = (*customUmsMemberProductCollectionModel)(nil)

type (
	// UmsMemberProductCollectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberProductCollectionModel.
	UmsMemberProductCollectionModel interface {
		umsMemberProductCollectionModel
		Count(ctx context.Context, in *umsclient.MemberProductCollectionListReq) (int64, error)
		FindAll(ctx context.Context, in *umsclient.MemberProductCollectionListReq) (*[]UmsMemberProductCollection, error)
		DeleteByIdAndMemberId(ctx context.Context, id int64, MemberId int64) error
	}

	customUmsMemberProductCollectionModel struct {
		*defaultUmsMemberProductCollectionModel
	}
)

// NewUmsMemberProductCollectionModel returns a model for the database table.
func NewUmsMemberProductCollectionModel(conn sqlx.SqlConn) UmsMemberProductCollectionModel {
	return &customUmsMemberProductCollectionModel{
		defaultUmsMemberProductCollectionModel: newUmsMemberProductCollectionModel(conn),
	}
}

func (m *customUmsMemberProductCollectionModel) FindAll(ctx context.Context, in *umsclient.MemberProductCollectionListReq) (*[]UmsMemberProductCollection, error) {
	where := "1=1"

	if in.MemberId != 0 {
		where = where + fmt.Sprintf(" AND member_id = '%d'", in.MemberId)
	}

	if in.ProductId != 0 {
		where = where + fmt.Sprintf(" AND product_id = '%d'", in.ProductId)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", umsMemberProductCollectionRows, m.table, where)

	var resp []UmsMemberProductCollection
	err := m.conn.QueryRows(&resp, query, (in.Current-1)*in.PageSize, in.PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUmsMemberProductCollectionModel) Count(ctx context.Context, in *umsclient.MemberProductCollectionListReq) (int64, error) {
	where := "1=1"
	if in.MemberId != 0 {
		where = where + fmt.Sprintf(" AND member_id = '%d'", in.MemberId)
	}

	if in.ProductId != 0 {
		where = where + fmt.Sprintf(" AND product_id = '%d'", in.ProductId)
	}

	query := fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)

	var count int64
	err := m.conn.QueryRow(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customUmsMemberProductCollectionModel) DeleteByIdAndMemberId(ctx context.Context, id int64, MemberId int64) error {

	//清空收藏列表
	query := fmt.Sprintf("DELETE FROM %s WHERE member_id = ? and id = ?", m.table)

	_, err := m.conn.ExecCtx(ctx, query, MemberId, id)
	return err
}
