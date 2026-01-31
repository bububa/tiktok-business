package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup/smartplus"
)

// Get Get Upgraded Smart+ Ad Groups
func Get(ctx context.Context, clt *core.SDKClient, req *smartplus.GetRequest, accessToken string) (*smartplus.GetResult, error) {
	var ret smartplus.GetResponse
	if err := clt.Get(ctx, "v1.3/smart_plus/adgroup/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
