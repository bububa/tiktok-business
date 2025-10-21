package product

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/product"
)

// File 通过文件 URL 上传商品
func File(ctx context.Context, clt *core.SDKClient, req *product.FileRequest, accessToken string) (string, error) {
	var ret product.FileResponse
	if err := clt.Post(ctx, "v1.3/catalog/product/file/", req, &ret, accessToken); err != nil {
		return "", err
	}
	return ret.Data.FeedLogID, nil
}
