package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetGroupGet 获取资产组信息
// 您可以使用本接口获取商务中心下一个资产组中的资产或成员信息。您必须是商务中心管理员用户
func AssetGroupGet(ctx context.Context, clt *core.SDKClient, req *bc.AssetGroupGetRequest, accessToken string) (*bc.AssetGroupGetResult, error) {
	var resp bc.AssetGroupGetResponse
	if err := clt.Get(ctx, "v1.3/bc/asset/group/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
