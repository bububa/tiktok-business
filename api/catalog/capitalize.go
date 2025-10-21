package catalog

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog"
)

// Capitalize 将商品库迁移至商务中心
// 您可以使用本接口将处于广告账户下的商品库迁移至商务中心，将其资产化为商务中心下的资产。
func Capitalize(ctx context.Context, clt *core.SDKClient, req *catalog.CapitalizeRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/catalog/capitalize/", req, nil, accessToken)
}
