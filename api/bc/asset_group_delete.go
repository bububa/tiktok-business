package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetGroupDelete 删除资产组
// 您可以使用本接口删除商务中心下的资产组。您必须是商务中心管理员用户。
// 请注意，该接口仅删除您分配给成员的资产组权限，而不会删除您的成员和资产组本身
func AssetGroupDelete(ctx context.Context, clt *core.SDKClient, req *bc.AssetGroupDeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/asset/group/delete/", req, nil, accessToken)
}
