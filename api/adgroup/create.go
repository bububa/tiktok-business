package adgroup

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup"
)

// Create 创建广告组 API Request
func Create(ctx context.Context, clt *core.SDKClient, req *adgroup.CreateRequest, accessToken string) (*adgroup.Adgroup, error) {
	var ret adgroup.CreateResponse
	if err := clt.Post(ctx, "v1.3/adgroup/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
