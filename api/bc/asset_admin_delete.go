package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetAdminDelete 删除资产
// 商务中心管理员用户可以使用本接口删除商务中心下的资产。
// 目前支持删除的资产类型包括线索和TikTok账号。
func AssetAdminDelete(ctx context.Context, clt *core.SDKClient, req *bc.AssetAdminDeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/asset/admin/delete/", req, nil, accessToken)
}
