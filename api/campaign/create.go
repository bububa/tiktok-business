package campaign

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign"
)

// Create 创建推广系列 API Request
func Create(ctx context.Context, clt *core.SDKClient, req *campaign.CreateRequest, accessToken string) (*campaign.Campaign, error) {
	var ret campaign.CreateResponse
	if err := clt.Post(ctx, "v1.3/campaign/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
