package campaign

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign"
)

// Get 获取推广系列列表
func Get(ctx context.Context, clt *core.SDKClient, req *campaign.GetRequest, accessToken string) (*campaign.GetResult, error) {
	var ret campaign.GetResponse
	if err := clt.Get(ctx, "v1.3/campaign/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
