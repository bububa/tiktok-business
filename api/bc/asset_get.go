package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetGet 获取资产列表
// 您可以使用本接口获取商务中心下您或者某个用户拥有权限的资产信息。
func AssetGet(ctx context.Context, clt *core.SDKClient, req *bc.AssetGetRequest, accessToken string) (*bc.AssetGetResult, error) {
	var resp bc.AssetGetResponse
	if err := clt.Get(ctx, "v1.3/bc/asset/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
