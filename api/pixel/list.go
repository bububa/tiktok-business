package pixel

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/pixel"
)

// List 获取Pixel列表
func List(ctx context.Context, clt *core.SDKClient, req *pixel.ListRequest, accessToken string) (*pixel.ListResult, error) {
	var ret pixel.ListResponse
	if err := clt.Get(ctx, "v1.3/pixel/list/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
