package product

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/product"
)

// Delete 删除商品
func Delete(ctx context.Context, clt *core.SDKClient, req *product.DeleteRequest, accessToken string) (string, error) {
	var resp product.DeleteResponse
	if err := clt.Post(ctx, "v1.3/catalog/product/delete/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.FeedLogID, nil
}
