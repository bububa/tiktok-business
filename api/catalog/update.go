package catalog

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog"
)

// Update 修改商品库名称
func Update(ctx context.Context, clt *core.SDKClient, req *catalog.UpdateRequest, accessToken string) (string, error) {
	var ret catalog.UpdateResponse
	if err := clt.Post(ctx, "v1.3/catalog/update/", req, &ret, accessToken); err != nil {
		return "", err
	}
	return ret.Data.CatalogID, nil
}
