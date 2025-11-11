package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AdvertiserCreate 创建广告账户
// 您可以使用本接口为代理商或直客商务中心创建竞价广告账户。您必须是商务中心管理员用户。
// 若想创建竞价广告账户，请在advertiser_info下的type字段指定AUCTION。
// 若广告账户的注册地在中国大陆、中国香港地区、法国、巴西或墨西哥，那么在创建广告账户之前，您需要使用/bc/image/upload/上传营业执照或其他认证资质的图片。
// 若广告账户的注册地在中国大陆，则您可能需要为营业执照提交银联验证。若想了解银联验证和在商务中心创建广告账户的详情，可查看创建广告账户。
func AdvertiserCreate(ctx context.Context, clt *core.SDKClient, req *bc.AdvertiserCreateRequest, accessToken string) (string, error) {
	var resp bc.AdvertiserCreateResponse
	if err := clt.Post(ctx, "v1.3/bc/advertiser/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.AdvertiserID, nil
}
