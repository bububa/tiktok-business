package adgroup

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup"
)

// Update 更新广告组 API Request
func Update(ctx context.Context, clt *core.SDKClient, req *adgroup.UpdateRequest, accessToken string) (*adgroup.Adgroup, error) {
	var ret adgroup.UpdateResponse
	if err := clt.Post(ctx, "v1.3/adgroup/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
