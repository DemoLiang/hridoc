// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskStatusLogic {
	return &GetTaskStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskStatusLogic) GetTaskStatus(req *types.TaskStatusReq) (resp *types.TaskStatusResp, err error) {
	task, err := l.svcCtx.ExportTaskModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.TaskStatusResp{
				BaseResp: types.BaseResp{Code: errorx.ErrTaskNotFound, Message: "任务不存在"},
			}, nil
		}
		logx.Errorf("find export task failed: %v", err)
		return &types.TaskStatusResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	data := types.TaskStatusData{
		Id:       task.Id,
		TaskName: task.TaskName.String,
		Status:   task.Status,
		FailReason: func() string {
			if task.FailReason.Valid {
				return task.FailReason.String
			}
			return ""
		}(),
		UserCount: task.UserCount,
		CertCount: task.CertCount,
		MissCount: task.MissCount,
		FileUrl: func() string {
			if task.FileUrl.Valid {
				return task.FileUrl.String
			}
			return ""
		}(),
		CreatedAt:   task.CreatedAt.Format("2006-01-02 15:04:05"),
		CompletedAt: func() string {
			if task.CompletedAt.Valid {
				return task.CompletedAt.Time.Format("2006-01-02 15:04:05")
			}
			return ""
		}(),
	}

	return &types.TaskStatusResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data:     data,
	}, nil
}
