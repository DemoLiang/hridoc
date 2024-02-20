// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package flashpromotionproductrelationservice

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

	FlashPromotionProductRelationService interface {
		FlashPromotionProductRelationAdd(ctx context.Context, in *FlashPromotionProductRelationAddReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationAddResp, error)
		FlashPromotionProductRelationList(ctx context.Context, in *FlashPromotionProductRelationListReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationListResp, error)
		FlashPromotionProductRelationUpdate(ctx context.Context, in *FlashPromotionProductRelationUpdateReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationUpdateResp, error)
		FlashPromotionProductRelationDelete(ctx context.Context, in *FlashPromotionProductRelationDeleteReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationDeleteResp, error)
		// 根据商品id查询
		QueryFlashPromotionByProduct(ctx context.Context, in *QueryFlashPromotionByProductReq, opts ...grpc.CallOption) (*QueryFlashPromotionByProductResp, error)
	}

	defaultFlashPromotionProductRelationService struct {
		cli zrpc.Client
	}
)

func NewFlashPromotionProductRelationService(cli zrpc.Client) FlashPromotionProductRelationService {
	return &defaultFlashPromotionProductRelationService{
		cli: cli,
	}
}

func (m *defaultFlashPromotionProductRelationService) FlashPromotionProductRelationAdd(ctx context.Context, in *FlashPromotionProductRelationAddReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationAddResp, error) {
	client := smsclient.NewFlashPromotionProductRelationServiceClient(m.cli.Conn())
	return client.FlashPromotionProductRelationAdd(ctx, in, opts...)
}

func (m *defaultFlashPromotionProductRelationService) FlashPromotionProductRelationList(ctx context.Context, in *FlashPromotionProductRelationListReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationListResp, error) {
	client := smsclient.NewFlashPromotionProductRelationServiceClient(m.cli.Conn())
	return client.FlashPromotionProductRelationList(ctx, in, opts...)
}

func (m *defaultFlashPromotionProductRelationService) FlashPromotionProductRelationUpdate(ctx context.Context, in *FlashPromotionProductRelationUpdateReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationUpdateResp, error) {
	client := smsclient.NewFlashPromotionProductRelationServiceClient(m.cli.Conn())
	return client.FlashPromotionProductRelationUpdate(ctx, in, opts...)
}

func (m *defaultFlashPromotionProductRelationService) FlashPromotionProductRelationDelete(ctx context.Context, in *FlashPromotionProductRelationDeleteReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationDeleteResp, error) {
	client := smsclient.NewFlashPromotionProductRelationServiceClient(m.cli.Conn())
	return client.FlashPromotionProductRelationDelete(ctx, in, opts...)
}

// 根据商品id查询
func (m *defaultFlashPromotionProductRelationService) QueryFlashPromotionByProduct(ctx context.Context, in *QueryFlashPromotionByProductReq, opts ...grpc.CallOption) (*QueryFlashPromotionByProductResp, error) {
	client := smsclient.NewFlashPromotionProductRelationServiceClient(m.cli.Conn())
	return client.QueryFlashPromotionByProduct(ctx, in, opts...)
}
