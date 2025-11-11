package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// PixelLinkGet 获取与pixel绑定的广告账户列表
func PixelLinkGet(ctx context.Context, clt *core.SDKClient, req *bc.PixelLinkGetRequest, accessToken string) (*bc.PixelLinkGetResult, error) {
	var resp bc.PixelLinkGetResponse
	if err := clt.Get(ctx, "v1.3/bc/pixel/link/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
