package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/smartplus"
)

// Update Update an Upgraded Smart+ Campaign
// Use this endpoint to update an Upgraded Smart+ Campaign.
// This endpoint supports incremental updates.
func Update(ctx context.Context, clt *core.SDKClient, req *smartplus.UpdateRequest, accessToken string) (*smartplus.Campaign, error) {
	var ret smartplus.UpdateResponse
	if err := clt.Post(ctx, "v1.3/smart_plus/campaign/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
