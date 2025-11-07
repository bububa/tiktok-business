package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// ExclusiveAuthorizationCreate 授予广告账号 TikTok Shop 的专属授权
// 您可以使用本接口授予某个广告账号使用某个 TikTok Shop 创建 GMV Max 推广系列的专属授权。
// 对于每个 TikTok Shop，只能授权一个广告账户使用该 TikTok Shop 创建 GMV Max 推广系列，以优化收入和销售活动。
func ExclusiveAuthorizationCreate(ctx context.Context, clt *core.SDKClient, req *gmvmax.ExclusiveAuthorizationCreateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/gmv_max/exclusive_authorization/create/", req, nil, accessToken)
}
