package campaign

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign"
)

// Update 更新推广系列 API Request
func Update(ctx context.Context, clt *core.SDKClient, req *campaign.UpdateRequest, accessToken string) (*campaign.Campaign, error) {
	var ret campaign.UpdateResponse
	if err := clt.Post(ctx, "v1.3/campaign/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
