package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// StoreList 获取 GMV Max 推广系列可用的 TikTok Shop
// 您可以使用本接口获取某个广告账号有权限访问的 TikTok Shop 列表，以及这些 TikTok Shop 是否可用于创建GMV Max 推广系列。
func StoreList(ctx context.Context, clt *core.SDKClient, req *gmvmax.StoreListRequest, accessToken string) ([]gmvmax.Store, error) {
	var resp gmvmax.StoreListResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/store/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.StoreList, nil
}
