package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetAdminGet 以管理员身份获取资产列表
// 管理员用户可以使用本接口获取商务中心下的所有资产的信息。
// 资产信息包括广告账户状态，商务中心与资产的关系类型，商务中心与资产的关系状态。管理员用户还可以使用这些条件对结果进行筛选。
func AssetAdminGet(ctx context.Context, clt *core.SDKClient, req *bc.AssetAdminGetRequest, accessToken string) (*bc.AssetGetResult, error) {
	var resp bc.AssetGetResponse
	if err := clt.Get(ctx, "v1.3/bc/asset/admin/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
