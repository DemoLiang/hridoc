package configservicelogic

import (
	"context"
	"zero-admin/rpc/proto/sys"

	"zero-admin/rpc/hridoc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateConfigRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigRoleLogic {
	return &UpdateConfigRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateConfigRoleLogic) UpdateConfigRole(in *sys.UpdateConfigRoleReq) (*sys.UpdateConfigRoleResp, error) {
	// todo: add your logic here and delete this line

	return &sys.UpdateConfigRoleResp{}, nil
}
