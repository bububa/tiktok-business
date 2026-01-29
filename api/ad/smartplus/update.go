package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad/smartplus"
)

// Update Update an Upgraded Smart+ Ad
func Update(ctx context.Context, clt *core.SDKClient, req *smartplus.UpdateRequest, accessToken string) (*smartplus.Ad, error) {
	var ret smartplus.UpdateResponse
	if err := clt.Post(ctx, "v1.3/smart_plus/ad/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
