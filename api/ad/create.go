package ad

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad"
)

// Create 创建广告 API Request
func Create(ctx context.Context, clt *core.SDKClient, req *ad.CreateRequest, accessToken string) (*ad.Ad, error) {
	var ret ad.CreateResponse
	if err := clt.Post(ctx, "v1.3/ad/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
