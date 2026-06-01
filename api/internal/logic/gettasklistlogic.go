// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskListLogic {
	return &GetTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskListLogic) GetTaskList(req *types.TaskListReq) (resp *types.TaskListResp, err error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	offset := (page - 1) * pageSize

	countQuery := "select count(*) from export_task"
	var total int64
	if err = l.svcCtx.DB.QueryRowCtx(l.ctx, &total, countQuery); err != nil {
		logx.Errorf("count export task failed: %v", err)
		return &types.TaskListResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	query := fmt.Sprintf("select id, task_name, user_count, cert_count, miss_count, status, fail_reason, file_url, created_at, completed_at from export_task order by id desc limit %d offset %d", pageSize, offset)
	var rows []struct {
		Id          int64  `db:"id"`
		TaskName    string `db:"task_name"`
		UserCount   int64  `db:"user_count"`
		CertCount   int64  `db:"cert_count"`
		MissCount   int64  `db:"miss_count"`
		Status      int64  `db:"status"`
		FailReason  string `db:"fail_reason"`
		FileUrl     string `db:"file_url"`
		CreatedAt   string `db:"created_at"`
		CompletedAt string `db:"completed_at"`
	}
	if err = l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &rows, query); err != nil {
		logx.Errorf("query export task list failed: %v", err)
		return &types.TaskListResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	var list []types.TaskStatusData
	for _, r := range rows {
		list = append(list, types.TaskStatusData{
			Id:          r.Id,
			TaskName:    r.TaskName,
			UserCount:   r.UserCount,
			CertCount:   r.CertCount,
			MissCount:   r.MissCount,
			Status:      r.Status,
			FailReason:  r.FailReason,
			FileUrl:     r.FileUrl,
			CreatedAt:   r.CreatedAt,
			CompletedAt: r.CompletedAt,
		})
	}

	return &types.TaskListResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.TaskListData{
			PageResp: types.PageResp{Total: total, Page: page},
			List:     list,
		},
	}, nil
}
