package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetGroupUpdate 更新资产组
// 您可以使用本接口更新一个资产组下的资产、成员或名称。您必须是商务中心管理员用户
func AssetGroupUpdate(ctx context.Context, clt *core.SDKClient, req *bc.AssetGroupUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/asset/group/update/", req, nil, accessToken)
}
