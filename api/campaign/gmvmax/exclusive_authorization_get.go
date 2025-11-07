package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// ExclusiveAuthorizationGet 获取广告账号的 TikTok Shop 专属授权状态
// 您可以使用本接口检查某个广告账号针对某个 TikTok Shop 是否获取了创建 GMV Max 推广系列的专属授权。
func ExclusiveAuthorizationGet(ctx context.Context, clt *core.SDKClient, req *gmvmax.ExclusiveAuthorizationGetRequest, accessToken string) (*gmvmax.ExclusiveAuthorization, error) {
	var resp gmvmax.ExclusiveAuthorizationGetResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/exclusive_authorization/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
