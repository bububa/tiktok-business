package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetBindQuota 获取资产绑定信息
// 您可以使用本接口查看资产已经绑定了多少广告以及剩下可以绑定的广告数。每个资产仅可用于创建一定数量的广告
func AssetBindQuota(ctx context.Context, clt *core.SDKClient, req *bc.AssetBindQuotaRequest, accessToken string) (*bc.AssetBindQuota, error) {
	var resp bc.AssetBindQuotaResponse
	if err := clt.Get(ctx, "v1.3/bc/asset/bind/quota/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
