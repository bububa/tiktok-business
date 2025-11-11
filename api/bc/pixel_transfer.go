package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// PixelTransfer 将pixel转移至商务中心
// 您可以使用本接口将pixel转移到商务中心。
// 为了操作成功，必须满足以下两点要求：
// 如果转移到代理商商务中心，操作必须由代理商的管理员用户完成。同时，pixel原属于的广告账户也属于此代理商BC。
// 如果转移到直客商务中心，此BC必须有这个pixel原属于的广告账户的管理员权限。
func PixelTransfer(ctx context.Context, clt *core.SDKClient, req *bc.PixelTransferRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/pixel/transfer/", req, nil, accessToken)
}
