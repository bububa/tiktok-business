package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetAccountAuthorization 获取 TikTok 账号广告投放授权链接
// 您可以使用本接口生成授权链接，从而请求 TikTok 账号的广告投放权限。
// 获取授权链接后，您可以使用二维码生成器将其转换为二维码，并分享给 TikTok 账号所有者。TikTok 账号所有者扫描二维码后，即可将 Spark Ads 拉取的“现有帖子”权限和 Spark Ads 推送的“TikTok 帖子”权限授予给商务中心。
func AssetAccountAuthorization(ctx context.Context, clt *core.SDKClient, req *bc.AssetAccountAuthorizationRequest, accessToken string) (string, error) {
	var resp bc.AssetAccountAuthorizationResponse
	if err := clt.Get(ctx, "v1.3/bc/asset/account/authorization/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.BcAuthQRCode, nil
}
