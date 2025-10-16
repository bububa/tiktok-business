package catalog

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog"
)

// Overview 获取商品库概览
// 您可以使用本接口获取商品库中处于不同状态（通过审核、审核被拒、审核处理中）的商品的数量。
func Overview(ctx context.Context, clt *core.SDKClient, req *catalog.OverviewRequest, accessToken string) (*catalog.Overview, error) {
	var ret catalog.OverviewResponse
	if err := clt.Get(ctx, "v1.3/catalog/overview/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
