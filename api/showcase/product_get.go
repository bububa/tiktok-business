package showcase

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/showcase"
)

// ProductGet 获取橱窗中的可用商品
func ProductGet(ctx context.Context, clt *core.SDKClient, req *showcase.ProductGetRequest, accessToken string) (*showcase.ProductGetResult, error) {
	var ret showcase.ProductGetResponse
	if err := clt.Get(ctx, "v1.3/showcase/product/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
