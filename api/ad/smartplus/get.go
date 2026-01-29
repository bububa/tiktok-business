package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad/smartplus"
)

// Get Get Upgraded Smart+ Ads
func Get(ctx context.Context, clt *core.SDKClient, req *smartplus.GetRequest, accessToken string) (*smartplus.GetResult, error) {
	var ret smartplus.GetResponse
	if err := clt.Get(ctx, "v1.3/smart_plus/ad/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
