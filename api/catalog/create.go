package catalog

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog"
)

// Create 创建商品库
func Create(ctx context.Context, clt *core.SDKClient, req *catalog.CreateRequest, accessToken string) (string, error) {
	var ret catalog.CreateResponse
	if err := clt.Post(ctx, "v1.3/catalog/create/", req, &ret, accessToken); err != nil {
		return "", err
	}
	return ret.Data.CatalogID, nil
}
