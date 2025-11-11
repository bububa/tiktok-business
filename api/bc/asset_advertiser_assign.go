package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetAdvertiserAssign 在商务中心内将 TikTok 账号和广告账号绑定
// 您可以使用本接口在商务中心内将 TikTok 账号和广告账号绑定。
// 绑定完成后，对 TikTok 账号绑定的广告账号有访问权限的商务中心成员即可在这些广告账号中使用该 TikTok 账号进行广告投放。
func AssetAdvertiserAssign(ctx context.Context, clt *core.SDKClient, req *bc.AssetAdvertiserAssignRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/asset/advertiser/assign/", req, nil, accessToken)
}
