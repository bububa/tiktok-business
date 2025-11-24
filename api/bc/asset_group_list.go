package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetGroupList 获取资产组列表
// 您可以使用本接口获取商务中心下的资产组列表。您必须是商务中心管理员用户
func AssetGroupList(ctx context.Context, clt *core.SDKClient, req *bc.AssetGroupListRequest, accessToken string) (*bc.AssetGroupListResult, error) {
	var resp bc.AssetGroupListResponse
	if err := clt.Get(ctx, "v1.3/bc/asset/group/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
