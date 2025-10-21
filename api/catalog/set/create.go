package set

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/set"
)

// Create 创建商品系列
// 您可以使用本接口在商务中心下的商品库创建一个商品系列。
func Create(ctx context.Context, clt *core.SDKClient, req *set.CreateRequest, accessToken string) (*set.CreateResult, error) {
	var resp set.CreateResponse
	if err := clt.Post(ctx, "v1.3/catalog/set/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
