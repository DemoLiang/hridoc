// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package couponhistoryservice

import (
	"context"

	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CouponAddReq                                  = smsclient.CouponAddReq
	CouponAddResp                                 = smsclient.CouponAddResp
	CouponCountReq                                = smsclient.CouponCountReq
	CouponCountResp                               = smsclient.CouponCountResp
	CouponDeleteReq                               = smsclient.CouponDeleteReq
	CouponDeleteResp                              = smsclient.CouponDeleteResp
	CouponFindByIdReq                             = smsclient.CouponFindByIdReq
	CouponFindByIdResp                            = smsclient.CouponFindByIdResp
	CouponFindByIdsReq                            = smsclient.CouponFindByIdsReq
	CouponFindByIdsResp                           = smsclient.CouponFindByIdsResp
	CouponFindByProductIdAndProductCategoryIdReq  = smsclient.CouponFindByProductIdAndProductCategoryIdReq
	CouponFindByProductIdAndProductCategoryIdResp = smsclient.CouponFindByProductIdAndProductCategoryIdResp
	CouponHistoryAddReq                           = smsclient.CouponHistoryAddReq
	CouponHistoryAddResp                          = smsclient.CouponHistoryAddResp
	CouponHistoryDeleteReq                        = smsclient.CouponHistoryDeleteReq
	CouponHistoryDeleteResp                       = smsclient.CouponHistoryDeleteResp
	CouponHistoryDetailListData                   = smsclient.CouponHistoryDetailListData
	CouponHistoryDetailListReq                    = smsclient.CouponHistoryDetailListReq
	CouponHistoryDetailListResp                   = smsclient.CouponHistoryDetailListResp
	CouponHistoryListData                         = smsclient.CouponHistoryListData
	CouponHistoryListReq                          = smsclient.CouponHistoryListReq
	CouponHistoryListResp                         = smsclient.CouponHistoryListResp
	CouponHistoryUpdateReq                        = smsclient.CouponHistoryUpdateReq
	CouponHistoryUpdateResp                       = smsclient.CouponHistoryUpdateResp
	CouponListData                                = smsclient.CouponListData
	CouponListReq                                 = smsclient.CouponListReq
	CouponListResp                                = smsclient.CouponListResp
	CouponProductCategoryRelationAddReq           = smsclient.CouponProductCategoryRelationAddReq
	CouponProductCategoryRelationAddResp          = smsclient.CouponProductCategoryRelationAddResp
	CouponProductCategoryRelationDeleteReq        = smsclient.CouponProductCategoryRelationDeleteReq
	CouponProductCategoryRelationDeleteResp       = smsclient.CouponProductCategoryRelationDeleteResp
	CouponProductCategoryRelationListData         = smsclient.CouponProductCategoryRelationListData
	CouponProductCategoryRelationListReq          = smsclient.CouponProductCategoryRelationListReq
	CouponProductCategoryRelationListResp         = smsclient.CouponProductCategoryRelationListResp
	CouponProductCategoryRelationUpdateReq        = smsclient.CouponProductCategoryRelationUpdateReq
	CouponProductCategoryRelationUpdateResp       = smsclient.CouponProductCategoryRelationUpdateResp
	CouponProductRelationAddReq                   = smsclient.CouponProductRelationAddReq
	CouponProductRelationAddResp                  = smsclient.CouponProductRelationAddResp
	CouponProductRelationDeleteReq                = smsclient.CouponProductRelationDeleteReq
	CouponProductRelationDeleteResp               = smsclient.CouponProductRelationDeleteResp
	CouponProductRelationListData                 = smsclient.CouponProductRelationListData
	CouponProductRelationListReq                  = smsclient.CouponProductRelationListReq
	CouponProductRelationListResp                 = smsclient.CouponProductRelationListResp
	CouponProductRelationUpdateReq                = smsclient.CouponProductRelationUpdateReq
	CouponProductRelationUpdateResp               = smsclient.CouponProductRelationUpdateResp
	CouponUpdateReq                               = smsclient.CouponUpdateReq
	CouponUpdateResp                              = smsclient.CouponUpdateResp
	FlashPromotionAddReq                          = smsclient.FlashPromotionAddReq
	FlashPromotionAddResp                         = smsclient.FlashPromotionAddResp
	FlashPromotionDeleteReq                       = smsclient.FlashPromotionDeleteReq
	FlashPromotionDeleteResp                      = smsclient.FlashPromotionDeleteResp
	FlashPromotionListByDateReq                   = smsclient.FlashPromotionListByDateReq
	FlashPromotionListByDateResp                  = smsclient.FlashPromotionListByDateResp
	FlashPromotionListData                        = smsclient.FlashPromotionListData
	FlashPromotionListReq                         = smsclient.FlashPromotionListReq
	FlashPromotionListResp                        = smsclient.FlashPromotionListResp
	FlashPromotionLogAddReq                       = smsclient.FlashPromotionLogAddReq
	FlashPromotionLogAddResp                      = smsclient.FlashPromotionLogAddResp
	FlashPromotionLogDeleteReq                    = smsclient.FlashPromotionLogDeleteReq
	FlashPromotionLogDeleteResp                   = smsclient.FlashPromotionLogDeleteResp
	FlashPromotionLogListData                     = smsclient.FlashPromotionLogListData
	FlashPromotionLogListReq                      = smsclient.FlashPromotionLogListReq
	FlashPromotionLogListResp                     = smsclient.FlashPromotionLogListResp
	FlashPromotionLogUpdateReq                    = smsclient.FlashPromotionLogUpdateReq
	FlashPromotionLogUpdateResp                   = smsclient.FlashPromotionLogUpdateResp
	FlashPromotionProductRelationAddReq           = smsclient.FlashPromotionProductRelationAddReq
	FlashPromotionProductRelationAddResp          = smsclient.FlashPromotionProductRelationAddResp
	FlashPromotionProductRelationDeleteReq        = smsclient.FlashPromotionProductRelationDeleteReq
	FlashPromotionProductRelationDeleteResp       = smsclient.FlashPromotionProductRelationDeleteResp
	FlashPromotionProductRelationListData         = smsclient.FlashPromotionProductRelationListData
	FlashPromotionProductRelationListReq          = smsclient.FlashPromotionProductRelationListReq
	FlashPromotionProductRelationListResp         = smsclient.FlashPromotionProductRelationListResp
	FlashPromotionProductRelationUpdateReq        = smsclient.FlashPromotionProductRelationUpdateReq
	FlashPromotionProductRelationUpdateResp       = smsclient.FlashPromotionProductRelationUpdateResp
	FlashPromotionSessionAddReq                   = smsclient.FlashPromotionSessionAddReq
	FlashPromotionSessionAddResp                  = smsclient.FlashPromotionSessionAddResp
	FlashPromotionSessionByTimeReq                = smsclient.FlashPromotionSessionByTimeReq
	FlashPromotionSessionByTimeResp               = smsclient.FlashPromotionSessionByTimeResp
	FlashPromotionSessionDeleteReq                = smsclient.FlashPromotionSessionDeleteReq
	FlashPromotionSessionDeleteResp               = smsclient.FlashPromotionSessionDeleteResp
	FlashPromotionSessionListData                 = smsclient.FlashPromotionSessionListData
	FlashPromotionSessionListReq                  = smsclient.FlashPromotionSessionListReq
	FlashPromotionSessionListResp                 = smsclient.FlashPromotionSessionListResp
	FlashPromotionSessionUpdateReq                = smsclient.FlashPromotionSessionUpdateReq
	FlashPromotionSessionUpdateResp               = smsclient.FlashPromotionSessionUpdateResp
	FlashPromotionUpdateReq                       = smsclient.FlashPromotionUpdateReq
	FlashPromotionUpdateResp                      = smsclient.FlashPromotionUpdateResp
	HomeAdvertiseAddReq                           = smsclient.HomeAdvertiseAddReq
	HomeAdvertiseAddResp                          = smsclient.HomeAdvertiseAddResp
	HomeAdvertiseDeleteReq                        = smsclient.HomeAdvertiseDeleteReq
	HomeAdvertiseDeleteResp                       = smsclient.HomeAdvertiseDeleteResp
	HomeAdvertiseListData                         = smsclient.HomeAdvertiseListData
	HomeAdvertiseListReq                          = smsclient.HomeAdvertiseListReq
	HomeAdvertiseListResp                         = smsclient.HomeAdvertiseListResp
	HomeAdvertiseUpdateReq                        = smsclient.HomeAdvertiseUpdateReq
	HomeAdvertiseUpdateResp                       = smsclient.HomeAdvertiseUpdateResp
	HomeBrandAddData                              = smsclient.HomeBrandAddData
	HomeBrandAddReq                               = smsclient.HomeBrandAddReq
	HomeBrandAddResp                              = smsclient.HomeBrandAddResp
	HomeBrandDeleteReq                            = smsclient.HomeBrandDeleteReq
	HomeBrandDeleteResp                           = smsclient.HomeBrandDeleteResp
	HomeBrandListData                             = smsclient.HomeBrandListData
	HomeBrandListReq                              = smsclient.HomeBrandListReq
	HomeBrandListResp                             = smsclient.HomeBrandListResp
	HomeBrandUpdateReq                            = smsclient.HomeBrandUpdateReq
	HomeBrandUpdateResp                           = smsclient.HomeBrandUpdateResp
	HomeNewProductAddData                         = smsclient.HomeNewProductAddData
	HomeNewProductAddReq                          = smsclient.HomeNewProductAddReq
	HomeNewProductAddResp                         = smsclient.HomeNewProductAddResp
	HomeNewProductDeleteReq                       = smsclient.HomeNewProductDeleteReq
	HomeNewProductDeleteResp                      = smsclient.HomeNewProductDeleteResp
	HomeNewProductListData                        = smsclient.HomeNewProductListData
	HomeNewProductListReq                         = smsclient.HomeNewProductListReq
	HomeNewProductListResp                        = smsclient.HomeNewProductListResp
	HomeNewProductUpdateReq                       = smsclient.HomeNewProductUpdateReq
	HomeNewProductUpdateResp                      = smsclient.HomeNewProductUpdateResp
	HomeRecommendProductAddData                   = smsclient.HomeRecommendProductAddData
	HomeRecommendProductAddReq                    = smsclient.HomeRecommendProductAddReq
	HomeRecommendProductAddResp                   = smsclient.HomeRecommendProductAddResp
	HomeRecommendProductDeleteReq                 = smsclient.HomeRecommendProductDeleteReq
	HomeRecommendProductDeleteResp                = smsclient.HomeRecommendProductDeleteResp
	HomeRecommendProductListData                  = smsclient.HomeRecommendProductListData
	HomeRecommendProductListReq                   = smsclient.HomeRecommendProductListReq
	HomeRecommendProductListResp                  = smsclient.HomeRecommendProductListResp
	HomeRecommendProductUpdateReq                 = smsclient.HomeRecommendProductUpdateReq
	HomeRecommendProductUpdateResp                = smsclient.HomeRecommendProductUpdateResp
	HomeRecommendSubjectAddData                   = smsclient.HomeRecommendSubjectAddData
	HomeRecommendSubjectAddReq                    = smsclient.HomeRecommendSubjectAddReq
	HomeRecommendSubjectAddResp                   = smsclient.HomeRecommendSubjectAddResp
	HomeRecommendSubjectDeleteReq                 = smsclient.HomeRecommendSubjectDeleteReq
	HomeRecommendSubjectDeleteResp                = smsclient.HomeRecommendSubjectDeleteResp
	HomeRecommendSubjectListData                  = smsclient.HomeRecommendSubjectListData
	HomeRecommendSubjectListReq                   = smsclient.HomeRecommendSubjectListReq
	HomeRecommendSubjectListResp                  = smsclient.HomeRecommendSubjectListResp
	HomeRecommendSubjectUpdateReq                 = smsclient.HomeRecommendSubjectUpdateReq
	HomeRecommendSubjectUpdateResp                = smsclient.HomeRecommendSubjectUpdateResp
	QueryFlashPromotionByProductReq               = smsclient.QueryFlashPromotionByProductReq
	QueryFlashPromotionByProductResp              = smsclient.QueryFlashPromotionByProductResp
	QueryMemberCouponListReq                      = smsclient.QueryMemberCouponListReq
	QueryMemberCouponListResp                     = smsclient.QueryMemberCouponListResp
	UpdateCouponStatusReq                         = smsclient.UpdateCouponStatusReq
	UpdateCouponStatusResp                        = smsclient.UpdateCouponStatusResp

	CouponHistoryService interface {
		CouponHistoryAdd(ctx context.Context, in *CouponHistoryAddReq, opts ...grpc.CallOption) (*CouponHistoryAddResp, error)
		CouponHistoryList(ctx context.Context, in *CouponHistoryListReq, opts ...grpc.CallOption) (*CouponHistoryListResp, error)
		CouponHistoryUpdate(ctx context.Context, in *CouponHistoryUpdateReq, opts ...grpc.CallOption) (*CouponHistoryUpdateResp, error)
		CouponHistoryDelete(ctx context.Context, in *CouponHistoryDeleteReq, opts ...grpc.CallOption) (*CouponHistoryDeleteResp, error)
		// 登录时获取用户还没有使用的获取优惠券数量
		CouponCount(ctx context.Context, in *CouponCountReq, opts ...grpc.CallOption) (*CouponCountResp, error)
		// 获取会员优惠券
		QueryMemberCouponList(ctx context.Context, in *QueryMemberCouponListReq, opts ...grpc.CallOption) (*QueryMemberCouponListResp, error)
		// 更新优惠券状态
		UpdateCouponStatus(ctx context.Context, in *UpdateCouponStatusReq, opts ...grpc.CallOption) (*UpdateCouponStatusResp, error)
		// 获取该用户所有优惠券(包括商品和优惠券,商品分类和优惠券的关联关糸)
		QueryCouponHistoryDetailList(ctx context.Context, in *CouponHistoryDetailListReq, opts ...grpc.CallOption) (*CouponHistoryDetailListResp, error)
	}

	defaultCouponHistoryService struct {
		cli zrpc.Client
	}
)

func NewCouponHistoryService(cli zrpc.Client) CouponHistoryService {
	return &defaultCouponHistoryService{
		cli: cli,
	}
}

func (m *defaultCouponHistoryService) CouponHistoryAdd(ctx context.Context, in *CouponHistoryAddReq, opts ...grpc.CallOption) (*CouponHistoryAddResp, error) {
	client := smsclient.NewCouponHistoryServiceClient(m.cli.Conn())
	return client.CouponHistoryAdd(ctx, in, opts...)
}

func (m *defaultCouponHistoryService) CouponHistoryList(ctx context.Context, in *CouponHistoryListReq, opts ...grpc.CallOption) (*CouponHistoryListResp, error) {
	client := smsclient.NewCouponHistoryServiceClient(m.cli.Conn())
	return client.CouponHistoryList(ctx, in, opts...)
}

func (m *defaultCouponHistoryService) CouponHistoryUpdate(ctx context.Context, in *CouponHistoryUpdateReq, opts ...grpc.CallOption) (*CouponHistoryUpdateResp, error) {
	client := smsclient.NewCouponHistoryServiceClient(m.cli.Conn())
	return client.CouponHistoryUpdate(ctx, in, opts...)
}

func (m *defaultCouponHistoryService) CouponHistoryDelete(ctx context.Context, in *CouponHistoryDeleteReq, opts ...grpc.CallOption) (*CouponHistoryDeleteResp, error) {
	client := smsclient.NewCouponHistoryServiceClient(m.cli.Conn())
	return client.CouponHistoryDelete(ctx, in, opts...)
}

// 登录时获取用户还没有使用的获取优惠券数量
func (m *defaultCouponHistoryService) CouponCount(ctx context.Context, in *CouponCountReq, opts ...grpc.CallOption) (*CouponCountResp, error) {
	client := smsclient.NewCouponHistoryServiceClient(m.cli.Conn())
	return client.CouponCount(ctx, in, opts...)
}

// 获取会员优惠券
func (m *defaultCouponHistoryService) QueryMemberCouponList(ctx context.Context, in *QueryMemberCouponListReq, opts ...grpc.CallOption) (*QueryMemberCouponListResp, error) {
	client := smsclient.NewCouponHistoryServiceClient(m.cli.Conn())
	return client.QueryMemberCouponList(ctx, in, opts...)
}

// 更新优惠券状态
func (m *defaultCouponHistoryService) UpdateCouponStatus(ctx context.Context, in *UpdateCouponStatusReq, opts ...grpc.CallOption) (*UpdateCouponStatusResp, error) {
	client := smsclient.NewCouponHistoryServiceClient(m.cli.Conn())
	return client.UpdateCouponStatus(ctx, in, opts...)
}

// 获取该用户所有优惠券(包括商品和优惠券,商品分类和优惠券的关联关糸)
func (m *defaultCouponHistoryService) QueryCouponHistoryDetailList(ctx context.Context, in *CouponHistoryDetailListReq, opts ...grpc.CallOption) (*CouponHistoryDetailListResp, error) {
	client := smsclient.NewCouponHistoryServiceClient(m.cli.Conn())
	return client.QueryCouponHistoryDetailList(ctx, in, opts...)
}
