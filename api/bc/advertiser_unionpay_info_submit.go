package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AdvertiserUnionpayInfoSubmit 为营业执照提交银联验证
// 您可以使用本接口为营业执照提交银联验证。
// 提交银联验证后，需等待约 20 秒，随后使用/advertiser/info/查询广告账户的 status。若广告账户由于银联验证失败而拒审，您需先使用/advertiser/update/重新提交您的营业执照信息（license_no 和license_image_id），再使用/bc/advertiser/unionpay_info/submit/重新提交银联验证。
// 若想了解在商务中心中创建广告账户的完整流程，可查看创建广告账户
func AdvertiserUnionpayInfoSubmit(ctx context.Context, clt *core.SDKClient, req *bc.AdvertiserUnionpayInfoSubmitRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/advertiser/unionpay_info/submit/", req, nil, accessToken)
}
