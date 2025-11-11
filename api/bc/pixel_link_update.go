package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// PixelLinkUpdate 将pixel和广告账户绑定/解绑
func PixelLinkUpdate(ctx context.Context, clt *core.SDKClient, req *bc.PixelLinkUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/pixel/link/update/", req, nil, accessToken)
}
