package set

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/set"
)

// Delete 删除商品系列
// 您可以使用本接口删除商务中心下商品库的一个已存在的商品系列
func Delete(ctx context.Context, clt *core.SDKClient, req *set.DeleteRequest, accessToken string) ([]string, error) {
	var resp set.DeleteResponse
	if err := clt.Post(ctx, "v1.3/catalog/set/delete/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ProductSetIDs, nil
}
