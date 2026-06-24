// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"io"
	"path"
	"strings"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type FileProxyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileProxyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileProxyLogic {
	return &FileProxyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileProxyLogic) ProxyFile(fileUrl string) (io.ReadCloser, string, string, error) {
	objectName := l.extractObjectName(fileUrl)
	if objectName == "" {
		return nil, "", "", fmt.Errorf("无效的文件链接")
	}

	reader, err := l.svcCtx.MinIO.GetObject(l.ctx, objectName)
	if err != nil {
		logx.Errorf("get object from minio failed: %v", err)
		return nil, "", "", fmt.Errorf("获取文件失败")
	}

	filename := path.Base(objectName)
	contentType := l.detectContentType(filename)
	return reader, filename, contentType, nil
}

func (l *FileProxyLogic) extractObjectName(fileUrl string) string {
	parsed := strings.Split(fileUrl, "?")[0]
	parts := strings.Split(parsed, "/")
	if len(parts) < 2 {
		return ""
	}
	return parts[len(parts)-2] + "/" + parts[len(parts)-1]
}

func (l *FileProxyLogic) detectContentType(filename string) string {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".pdf":
		return "application/pdf"
	case ".zip":
		return "application/zip"
	case ".xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	default:
		return "application/octet-stream"
	}
}
