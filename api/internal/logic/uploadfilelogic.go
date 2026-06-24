// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"strings"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func detectFileType(data []byte) string {
	if len(data) < 4 {
		return "unknown"
	}
	switch {
	case data[0] == 0xFF && data[1] == 0xD8:
		return "image"
	case data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47:
		return "image"
	case data[0] == 0x25 && data[1] == 0x50 && data[2] == 0x44 && data[3] == 0x46:
		return "pdf"
	case data[0] == 0x47 && data[1] == 0x49 && data[2] == 0x46:
		return "image"
	case data[0] == 0x52 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x46:
		return "image"
	default:
		return "unknown"
	}
}

func contentType(fileType string) string {
	switch fileType {
	case "pdf":
		return "application/pdf"
	case "image":
		return "image/jpeg"
	default:
		return "application/octet-stream"
	}
}

func fileExt(fileType string) string {
	switch fileType {
	case "pdf":
		return ".pdf"
	case "image":
		return ".jpg"
	default:
		return ".bin"
	}
}

func isAllowedType(fileType string, allowed []string) bool {
	if len(allowed) == 0 {
		return true
	}
	for _, t := range allowed {
		if strings.EqualFold(t, fileType) {
			return true
		}
		// generic "image" matches any image/* in allowed list
		if fileType == "image" && strings.HasPrefix(t, "image/") {
			return true
		}
	}
	return false
}

func generateThumbnail(data []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	maxDim := 200
	if bounds.Dx() <= maxDim && bounds.Dy() <= maxDim {
		return nil, nil
	}

	var newW, newH int
	if bounds.Dx() > bounds.Dy() {
		newW = maxDim
		newH = bounds.Dy() * maxDim / bounds.Dx()
	} else {
		newH = maxDim
		newW = bounds.Dx() * maxDim / bounds.Dy()
	}
	if newW < 1 {
		newW = 1
	}
	if newH < 1 {
		newH = 1
	}

	thumb := image.NewRGBA(image.Rect(0, 0, newW, newH))
	sx := float64(bounds.Dx()) / float64(newW)
	sy := float64(bounds.Dy()) / float64(newH)
	for y := 0; y < newH; y++ {
		for x := 0; x < newW; x++ {
			srcX := int(float64(x) * sx)
			srcY := int(float64(y) * sy)
			thumb.Set(x, y, img.At(srcX, srcY))
		}
	}

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, thumb, &jpeg.Options{Quality: 80}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (l *UploadFileLogic) UploadFile(req *types.UploadReq) (resp *types.UploadResp, err error) {
	if len(req.File) == 0 {
		return &types.UploadResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidParam, Message: "文件不能为空"},
		}, nil
	}

	maxSize := l.svcCtx.Config.Upload.MaxFileSize
	if maxSize > 0 && int64(len(req.File)) > maxSize {
		return &types.UploadResp{
			BaseResp: types.BaseResp{Code: errorx.ErrFileTooLarge, Message: "文件过大"},
		}, nil
	}

	fileType := detectFileType(req.File)
	if fileType == "unknown" {
		return &types.UploadResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidFileType, Message: "无法识别的文件类型"},
		}, nil
	}
	if !isAllowedType(fileType, l.svcCtx.Config.Upload.AllowedTypes) {
		return &types.UploadResp{
			BaseResp: types.BaseResp{Code: errorx.ErrInvalidFileType, Message: "不支持的文件类型"},
		}, nil
	}

	if l.svcCtx.MinIO == nil {
		return &types.UploadResp{
			BaseResp: types.BaseResp{Code: errorx.ErrMinIOFailed, Message: "MinIO 未就绪"},
		}, nil
	}

	objectName := fmt.Sprintf("certs/%s%s", uuid.New().String(), fileExt(fileType))
	ct := contentType(fileType)
	err = l.svcCtx.MinIO.PutObject(l.ctx, objectName, bytes.NewReader(req.File), int64(len(req.File)), ct)
	if err != nil {
		logx.Errorf("upload file to minio failed: %v", err)
		return &types.UploadResp{
			BaseResp: types.BaseResp{Code: errorx.ErrUploadFailed, Message: "文件上传失败"},
		}, nil
	}

	var thumbUrl string
	if fileType == "image" {
		thumbData, err := generateThumbnail(req.File)
		if err == nil && len(thumbData) > 0 {
			thumbName := fmt.Sprintf("thumb/%s", objectName)
			if err := l.svcCtx.MinIO.PutObject(l.ctx, thumbName, bytes.NewReader(thumbData), int64(len(thumbData)), "image/jpeg"); err == nil {
				thumbUrl = l.svcCtx.MinIO.ObjectURL(thumbName)
			}
		}
	}

	url := l.svcCtx.MinIO.ObjectURL(objectName)
	return &types.UploadResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data: types.UploadData{
			Url:      url,
			ThumbUrl: thumbUrl,
			FileType: fileType,
		},
	}, nil
}
