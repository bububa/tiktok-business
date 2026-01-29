package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup/smartplus"
)

// Update Update an Upgraded Smart+ Ad Group
func Update(ctx context.Context, clt *core.SDKClient, req *smartplus.UpdateRequest, accessToken string) (*smartplus.Adgroup, error) {
	var ret smartplus.UpdateResponse
	if err := clt.Post(ctx, "v1.3/smartplus/adgroup/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
