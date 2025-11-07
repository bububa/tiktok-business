package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// Create 创建 GMV Max 推广系列
func Create(ctx context.Context, clt *core.SDKClient, req *gmvmax.CreateRequest, accessToken string) (*gmvmax.Campaign, error) {
	var resp gmvmax.CreateResponse
	if err := clt.Post(ctx, "v1.3/campaign/gmv_max/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
