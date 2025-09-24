package pixel

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/pixel"
)

// Create 创建pixel
func Create(ctx context.Context, clt *core.SDKClient, req *pixel.CreateRequest, accessToken string) (*pixel.Pixel, error) {
	var ret pixel.CreateResponse
	if err := clt.Post(ctx, "v1.3/pixel/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
