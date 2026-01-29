package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/smartplus"
)

// StatusUpdate Update the operation statuses of Upgraded Smart+ Campaigns
func StatusUpdate(ctx context.Context, clt *core.SDKClient, req *smartplus.StatusUpdateRequest, accessToken string) (*smartplus.StatusUpdateResult, error) {
	var ret smartplus.StatusUpdateResponse
	if err := clt.Post(ctx, "v1.3/smartplus/campaign/status/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
