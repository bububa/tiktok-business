package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AdvertiserUnionpayInfoCheck 查询营业执照的银联验证要求
// 您可以使用本接口查询您的营业执照是否需要进行银联验证。
// 对于在中国大陆注册且属于以下类型的广告账户，您需在将该营业执照传入/bc/advertiser/create/创建广告账户后，对该营业执照进行银联验证。
// 个体工商户
// 个人独资企业
// 有限责任公司（自然人独资）或有限责任公司（港澳台自然人独资）
// 有限责任公司（自然人投资或控股）
// 若您未对营业执照进行必要的银联验证，广告账户将在审核流程中被拒审。
// 若想了解在商务中心中创建广告账户的完整流程，可查看创建广告账户
func AdvertiserUnionpayInfoCheck(ctx context.Context, clt *core.SDKClient, req *bc.AdvertiserUnionpayInfoCheckRequest, accessToken string) (bool, error) {
	var resp bc.AdvertiserUnionpayInfoCheckResponse
	if err := clt.Get(ctx, "v1.3/bc/advertiser/unionpay_info/check/", req, &resp, accessToken); err != nil {
		return false, err
	}
	return resp.Data.UnionpayVerificationRequired, nil
}
