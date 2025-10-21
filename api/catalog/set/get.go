package set

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/set"
)

// Get 获取商品系列列表
// 您可以使用本接口获取商务中心下商品库的某一商品系列或商品系列列表。
func Get(ctx context.Context, clt *core.SDKClient, req *set.GetRequest, accessToken string) ([]set.ProductSet, error) {
	var resp set.GetResponse
	if err := clt.Get(ctx, "v1.3/catalog/set/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
