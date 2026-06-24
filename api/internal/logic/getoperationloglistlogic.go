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

type GetOperationLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOperationLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOperationLogListLogic {
	return &GetOperationLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOperationLogListLogic) GetOperationLogList(req *types.OperationLogListReq) (resp *types.OperationLogListResp, err error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	var logs []types.OperationLogInfo
	var total int64

	where := "1=1"
	args := []any{}

	module := strings.TrimSpace(req.Module)
	if module != "" {
		where += " and `module` = ?"
		args = append(args, module)
	}
	action := strings.TrimSpace(req.Action)
	if action != "" {
		where += " and `action` = ?"
		args = append(args, action)
	}
	if req.OperatorId > 0 {
		where += " and `operator_id` = ?"
		args = append(args, req.OperatorId)
	}

	countQuery := fmt.Sprintf("select count(*) from `operation_log` where %s", where)
	err = l.svcCtx.DB.QueryRowCtx(l.ctx, &total, countQuery, args...)
	if err != nil {
		logx.Errorf("count operation log failed: %v", err)
		return &types.OperationLogListResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	if total > 0 {
		listQuery := fmt.Sprintf("select id, operator_id, operator_name, module, action, target, detail, ip, created_at from `operation_log` where %s order by id desc limit ? offset ?", where)
		queryArgs := append(args, pageSize, (page-1)*pageSize)
		var rows []struct {
			Id           int64          `db:"id"`
			OperatorId   int64          `db:"operator_id"`
			OperatorName string         `db:"operator_name"`
			Module       string         `db:"module"`
			Action       string         `db:"action"`
			Target       sql.NullString `db:"target"`
			Detail       sql.NullString `db:"detail"`
			Ip           sql.NullString `db:"ip"`
			CreatedAt    sql.NullTime   `db:"created_at"`
		}
		err := l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &rows, listQuery, queryArgs...)
		if err != nil {
			logx.Errorf("query operation log list failed: %v", err)
			return &types.OperationLogListResp{
				BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
			}, nil
		}

		for _, r := range rows {
			logs = append(logs, types.OperationLogInfo{
				Id:           r.Id,
				OperatorId:   r.OperatorId,
				OperatorName: r.OperatorName,
				Module:       r.Module,
				Action:       r.Action,
				Target:       nullString(r.Target),
				Detail:       nullString(r.Detail),
				Ip:           nullString(r.Ip),
				CreatedAt:    formatDateTime(r.CreatedAt),
			})
		}
	}

	return &types.OperationLogListResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.OperationLogListData{
			PageResp: types.PageResp{Total: total, Page: page},
			List:     logs,
		},
	}, nil
}

