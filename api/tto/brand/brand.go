package brand

import (
	"context"
	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tto/brand"
)

func Create(ctx context.Context, clt *core.SDKClient, req *brand.CreateRequest, accessToken string) (*brand.Profile, error) {
	var resp brand.CreateResponse
	if err := clt.Post(ctx, "v1.3/tto/tcm/brand/profile/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func Get(ctx context.Context, clt *core.SDKClient, req *brand.GetRequest, accessToken string) (*brand.GetResult, error) {
	var resp brand.GetResponse
	if err := clt.Get(ctx, "v1.3/tto/tcm/brand/profile/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
