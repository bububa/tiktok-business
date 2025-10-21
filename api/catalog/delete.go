package catalog

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog"
)

// Delete 删除商品库
func Delete(ctx context.Context, clt *core.SDKClient, req *catalog.DeleteRequest, accessToken string) (string, error) {
	var ret catalog.DeleteResponse
	if err := clt.Post(ctx, "v1.3/catalog/delete/", req, &ret, accessToken); err != nil {
		return "", err
	}
	return ret.Data.CatalogID, nil
}
