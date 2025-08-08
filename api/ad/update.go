package ad

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad"
)

// Update 更新广告 API Request
func Update(ctx context.Context, clt *core.SDKClient, req *ad.UpdateRequest, accessToken string) (*ad.UpdateResult, error) {
	var ret ad.UpdateResponse
	if err := clt.Post(ctx, "v1.3/ad/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
