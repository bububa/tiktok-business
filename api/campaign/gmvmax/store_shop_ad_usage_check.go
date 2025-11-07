package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// StoreShopAdUsageCheck 检查 TikTok Shop 是否可用于商品 GMV Max 推广系列
// 您可以使用本接口检查某个 TikTok Shop 是否已在处于启用状态的视频购物广告或商品购物广告中占用，以及是否支持创建推广此 TikTok Shop 中的所有商品的商品 GMV Max 推广系列
func StoreShopAdUsageCheck(ctx context.Context, clt *core.SDKClient, req *gmvmax.StoreShopAdUsageCheckRequest, accessToken string) (*gmvmax.StoreShopAdUsageCheckResult, error) {
	var resp gmvmax.StoreShopAdUsageCheckResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/store/shop_ad_usage_check/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
