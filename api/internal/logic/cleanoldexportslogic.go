// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CleanOldExportsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCleanOldExportsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanOldExportsLogic {
	return &CleanOldExportsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CleanOldExportsLogic) CleanOldExports() (resp *types.CleanResp, err error) {
	cleanupDays := l.svcCtx.Config.Export.CleanupDays
	if cleanupDays <= 0 {
		cleanupDays = 7
	}
	cutoff := time.Now().AddDate(0, 0, -cleanupDays).Format("2006-01-02 15:04:05")

	query := fmt.Sprintf("select id, file_url from export_task where (status = 2 or status = 3) and created_at < '%s'", cutoff)
	var rows []struct {
		Id     int64  `db:"id"`
		FileUrl string `db:"file_url"`
	}
	if err = l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &rows, query); err != nil {
		logx.Errorf("query old exports failed: %v", err)
		return &types.CleanResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "查询旧导出任务失败"},
		}, nil
	}

	var deleted int64
	for _, r := range rows {
		if r.FileUrl != "" {
			objectName := l.extractObjectName(r.FileUrl)
			if objectName != "" {
				if err := l.svcCtx.MinIO.RemoveObject(l.ctx, objectName); err != nil {
					logx.Errorf("remove minio object failed: %v", err)
				}
			}
		}
		if err := l.svcCtx.ExportTaskModel.Delete(l.ctx, r.Id); err != nil {
			logx.Errorf("delete export task %d failed: %v", r.Id, err)
			continue
		}
		deleted++
	}

	return &types.CleanResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data:     types.CleanData{DeletedFiles: deleted},
	}, nil
}

func (l *CleanOldExportsLogic) extractObjectName(fileUrl string) string {
	parts := strings.Split(fileUrl, "/")
	if len(parts) < 2 {
		return ""
	}
	return strings.Join(parts[len(parts)-2:], "/")
}
