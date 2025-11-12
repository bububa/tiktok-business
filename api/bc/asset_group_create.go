package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetGroupCreate 创建资产组
// 您可以使用本接口在商务中心下创建一个资产组。您必须是商务中心管理员用户
func AssetGroupCreate(ctx context.Context, clt *core.SDKClient, req *bc.AssetGroupCreateRequest, accessToken string) (string, error) {
	var resp bc.AssetGroupCreateResponse
	if err := clt.Post(ctx, "v1.3/bc/asset/group/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.AssetGroupID, nil
}
