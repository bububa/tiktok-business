package catalog

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog"
)

// Get 获取商品库列表
func Get(ctx context.Context, clt *core.SDKClient, req *catalog.GetRequest, accessToken string) (*catalog.GetResult, error) {
	var ret catalog.GetResponse
	if err := clt.Get(ctx, "v1.3/catalog/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
