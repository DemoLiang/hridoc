// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type PreviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreviewLogic {
	return &PreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PreviewLogic) Preview(req *types.PreviewReq) (resp *types.PreviewResp, err error) {
	if len(req.UserIds) == 0 {
		return &types.PreviewResp{
			BaseResp: types.BaseResp{Code: 0, Message: "success"},
			Data:     types.PreviewData{Users: []types.PreviewUser{}},
		}, nil
	}
	if len(req.CategoryCodes) == 0 {
		return &types.PreviewResp{
			BaseResp: types.BaseResp{Code: 0, Message: "success"},
			Data:     types.PreviewData{Users: []types.PreviewUser{}},
		}, nil
	}

	userPlaceholders := make([]string, len(req.UserIds))
	userArgs := make([]any, len(req.UserIds))
	for i, id := range req.UserIds {
		userPlaceholders[i] = "?"
		userArgs[i] = id
	}

	catPlaceholders := make([]string, len(req.CategoryCodes))
	catArgs := make([]any, len(req.CategoryCodes))
	for i, code := range req.CategoryCodes {
		catPlaceholders[i] = "?"
		catArgs[i] = code
	}

	userQuery := fmt.Sprintf("select id, name from user where id in (%s) and status = 1", strings.Join(userPlaceholders, ","))
	var userRows []struct {
		Id   int64  `db:"id"`
		Name string `db:"name"`
	}
	if err = l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &userRows, userQuery, userArgs...); err != nil {
		logx.Errorf("query users failed: %v", err)
		return &types.PreviewResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	userMap := make(map[int64]string, len(userRows))
	for _, u := range userRows {
		userMap[u.Id] = u.Name
	}

	catQuery := fmt.Sprintf("select id, code, name from cert_category where code in (%s) and status = 1", strings.Join(catPlaceholders, ","))
	var catRows []struct {
		Id   int64  `db:"id"`
		Code string `db:"code"`
		Name string `db:"name"`
	}
	if err = l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &catRows, catQuery, catArgs...); err != nil {
		logx.Errorf("query categories failed: %v", err)
		return &types.PreviewResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	catIdMap := make(map[int64]struct{ Code, Name string }, len(catRows))
	catCodeMap := make(map[string]struct{ Id int64; Name string }, len(catRows))
	for _, c := range catRows {
		catIdMap[c.Id] = struct{ Code, Name string }{Code: c.Code, Name: c.Name}
		catCodeMap[c.Code] = struct{ Id int64; Name string }{Id: c.Id, Name: c.Name}
	}

	certPlaceholders := make([]string, len(req.UserIds))
	certArgs := make([]any, len(req.UserIds))
	for i, id := range req.UserIds {
		certPlaceholders[i] = "?"
		certArgs[i] = id
	}
	certQuery := fmt.Sprintf(`select c.id, c.user_id, c.category_id, c.name as cert_name
		from certificate c
		where c.user_id in (%s) and c.status = 1
		order by c.user_id, c.category_id, c.id desc`, strings.Join(certPlaceholders, ","))
	var certRows []struct {
		Id         int64  `db:"id"`
		UserId     int64  `db:"user_id"`
		CategoryId int64  `db:"category_id"`
		CertName   string `db:"cert_name"`
	}
	if err = l.svcCtx.DB.QueryRowsPartialCtx(l.ctx, &certRows, certQuery, certArgs...); err != nil {
		logx.Errorf("query certs failed: %v", err)
		return &types.PreviewResp{
			BaseResp: types.BaseResp{Code: errorx.ErrSystem, Message: "系统错误"},
		}, nil
	}

	certMap := make(map[int64]map[string]struct {
		CertId   int64
		CertName string
	})
	for _, r := range certRows {
		if _, ok := certMap[r.UserId]; !ok {
			certMap[r.UserId] = make(map[string]struct{ CertId int64; CertName string })
		}
		if catInfo, ok := catIdMap[r.CategoryId]; ok {
			if _, exists := certMap[r.UserId][catInfo.Code]; !exists {
				certMap[r.UserId][catInfo.Code] = struct {
					CertId   int64
					CertName string
				}{CertId: r.Id, CertName: r.CertName}
			}
		}
	}

	var users []types.PreviewUser
	for _, uid := range req.UserIds {
		name, ok := userMap[uid]
		if !ok {
			continue
		}
		pu := types.PreviewUser{
			UserId:   uid,
			UserName: name,
		}
		for _, code := range req.CategoryCodes {
			catInfo, ok := catCodeMap[code]
			if !ok {
				continue
			}
			pc := types.PreviewCategory{
				CategoryCode: code,
				CategoryName: catInfo.Name,
			}
			if certs, ok := certMap[uid]; ok {
				if c, ok := certs[code]; ok {
					pc.HasCert = true
					pc.CertId = c.CertId
					pc.CertName = c.CertName
				}
			}
			pu.Categories = append(pu.Categories, pc)
		}
		users = append(users, pu)
	}

	return &types.PreviewResp{
		BaseResp: types.BaseResp{Code: 0, Message: "success"},
		Data:     types.PreviewData{Users: users},
	}, nil
}
