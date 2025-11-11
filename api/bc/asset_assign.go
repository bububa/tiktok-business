package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetAssign 将资产分配给用户
// 商务中心的管理员用户可以使用本接口将资产分配给一个用户。该用户必须已在当前商务中心内
func AssetAssign(ctx context.Context, clt *core.SDKClient, req *bc.AssetAssignRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/asset/assign/", req, nil, accessToken)
}
