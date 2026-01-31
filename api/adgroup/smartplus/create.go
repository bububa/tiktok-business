package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup/smartplus"
)

// Create Create an Upgraded Smart+ Ad Group
func Create(ctx context.Context, clt *core.SDKClient, req *smartplus.CreateRequest, accessToken string) (*smartplus.Adgroup, error) {
	var ret smartplus.CreateResponse
	if err := clt.Post(ctx, "v1.3/smart_plus/adgroup/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
