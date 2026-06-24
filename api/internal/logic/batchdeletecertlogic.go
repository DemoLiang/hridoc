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

type BatchDeleteCertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchDeleteCertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteCertLogic {
	return &BatchDeleteCertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteCertLogic) BatchDeleteCert(req *types.BatchDeleteReq) (resp *types.BaseResp, err error) {
	if len(req.Ids) == 0 {
		return &types.BaseResp{Code: errorx.ErrInvalidParam, Message: "缺少删除 ID"}, nil
	}

	for _, id := range req.Ids {
		cert, err := l.svcCtx.CertificateModel.FindOne(l.ctx, id)
		if err != nil {
			if errorx.IsNotFound(err) {
				continue
			}
			logx.Errorf("find certificate %d failed: %v", id, err)
			continue
		}

		if l.svcCtx.MinIO != nil {
			if cert.FileUrl != "" {
				objectName := extractObjectName(cert.FileUrl)
				if objectName != "" {
					_ = l.svcCtx.MinIO.RemoveObject(l.ctx, objectName)
				}
			}
			if cert.ThumbUrl.Valid && cert.ThumbUrl.String != "" {
				thumbName := extractObjectName(cert.ThumbUrl.String)
				if thumbName != "" {
					_ = l.svcCtx.MinIO.RemoveObject(l.ctx, thumbName)
				}
			}
		}

		if err = l.svcCtx.CertificateModel.Delete(l.ctx, id); err != nil {
			logx.Errorf("delete certificate %d failed: %v", id, err)
		}
	}

	return &types.BaseResp{Code: 0, Message: "success"}, nil
}

