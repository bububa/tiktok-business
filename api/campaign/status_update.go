package campaign

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign"
)

// StatusUpdate 更新推广系列操作状态 API Request
func StatusUpdate(ctx context.Context, clt *core.SDKClient, req *campaign.StatusUpdateRequest, accessToken string) (*campaign.StatusUpdateResult, error) {
	var ret campaign.StatusUpdateResponse
	if err := clt.Post(ctx, "v1.3/campaign/status/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
