// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskDownloadLogic {
	return &GetTaskDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskDownloadLogic) GetTaskDownload(req *types.TaskStatusReq) (resp *types.DownloadResp, err error) {
	task, err := l.svcCtx.ExportTaskModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if errorx.IsNotFound(err) {
			return &types.DownloadResp{BaseResp: types.BaseResp{Code: errorx.ErrTaskNotFound, Message: "任务不存在"}}, nil
		}
		logx.Errorf("find export task failed: %v", err)
		return &types.DownloadResp{BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"}}, nil
	}

	if task.Status != 2 {
		return &types.DownloadResp{BaseResp: types.BaseResp{Code: errorx.ErrTaskFailed, Message: "任务尚未完成"}}, nil
	}

	if !task.FileUrl.Valid || task.FileUrl.String == "" {
		return &types.DownloadResp{BaseResp: types.BaseResp{Code: errorx.ErrTaskFailed, Message: "下载链接不存在"}}, nil
	}

	objectName := l.extractObjectName(task.FileUrl.String)
	if objectName == "" {
		return &types.DownloadResp{BaseResp: types.BaseResp{Code: errorx.ErrTaskFailed, Message: "无效的文件链接"}}, nil
	}

	url, err := l.svcCtx.MinIO.PresignedGetURL(l.ctx, objectName, time.Duration(1)*time.Hour)
	if err != nil {
		logx.Errorf("generate presigned url failed: %v", err)
		return &types.DownloadResp{BaseResp: types.BaseResp{Code: errorx.ErrMinIOFailed, Message: "生成下载链接失败"}}, nil
	}

	return &types.DownloadResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data:     types.DownloadData{Url: url},
	}, nil
}

func (l *GetTaskDownloadLogic) GetTaskDownloadFile(req *types.TaskStatusReq) (io.ReadCloser, string, string, error) {
	task, err := l.svcCtx.ExportTaskModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if errorx.IsNotFound(err) {
			return nil, "", "", fmt.Errorf("任务不存在")
		}
		logx.Errorf("find export task failed: %v", err)
		return nil, "", "", fmt.Errorf("系统错误")
	}

	if task.Status != 2 {
		return nil, "", "", fmt.Errorf("任务尚未完成")
	}

	if !task.FileUrl.Valid || task.FileUrl.String == "" {
		return nil, "", "", fmt.Errorf("下载链接不存在")
	}

	objectName := l.extractObjectName(task.FileUrl.String)
	if objectName == "" {
		return nil, "", "", fmt.Errorf("无效的文件链接")
	}

	reader, err := l.svcCtx.MinIO.GetObject(l.ctx, objectName)
	if err != nil {
		logx.Errorf("get object from minio failed: %v", err)
		return nil, "", "", fmt.Errorf("获取文件失败")
	}

	filename := fmt.Sprintf("task_%d.zip", req.Id)
	return reader, filename, "application/zip", nil
}

func (l *GetTaskDownloadLogic) extractObjectName(fileUrl string) string {
	parts := strings.Split(fileUrl, "/")
	if len(parts) < 2 {
		return ""
	}
	return parts[len(parts)-2] + "/" + parts[len(parts)-1]
}
