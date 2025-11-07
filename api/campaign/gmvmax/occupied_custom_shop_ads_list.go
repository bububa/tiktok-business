package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// OccupiedCustomShopAdsList 检查认证身份和商品是否在购物广告中占用
// 您可以使用本接口检查某个认证身份或商品是否被处于启用状态的购物广告（视频购物广告，商品购物广告或直播购物广告）占用
func OccupiedCustomShopAdsList(ctx context.Context, clt *core.SDKClient, req *gmvmax.OccupiedCustomShopAdsListRequest, accessToken string) ([]ad.Ad, error) {
	var resp gmvmax.OccupiedCustomShopAdsListResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/occupied_custom_shop_ads/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.OccupiedShopAds, nil
}
