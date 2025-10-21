package product

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/product"
)

// Upload 通过 JSON 结构上传商品
func Upload(ctx context.Context, clt *core.SDKClient, req *product.UploadRequest, accessToken string) (string, error) {
	var resp product.UploadResponse
	if err := clt.Post(ctx, "v1.3/catalog/product/upload/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.FeedLogID, nil
}
