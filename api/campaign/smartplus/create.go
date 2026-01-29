package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/smartplus"
)

// Create Create an Upgraded Smart+ Campaign
func Create(ctx context.Context, clt *core.SDKClient, req *smartplus.CreateRequest, accessToken string) (*smartplus.Campaign, error) {
	var ret smartplus.CreateResponse
	if err := clt.Post(ctx, "v1.3/smart_plus/campaign/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
