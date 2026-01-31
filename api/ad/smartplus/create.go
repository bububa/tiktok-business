package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad/smartplus"
)

// Create Create an Upgraded Smart+ Ad
func Create(ctx context.Context, clt *core.SDKClient, req *smartplus.CreateRequest, accessToken string) (*smartplus.Ad, error) {
	var ret smartplus.CreateResponse
	if err := clt.Post(ctx, "v1.3/smart_plus/ad/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
