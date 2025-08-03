package spc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/spc"
)

// Update 更新Smart+推广系列 API Request
func Update(ctx context.Context, clt *core.SDKClient, req *spc.UpdateRequest, accessToken string) (*spc.Campaign, error) {
	var ret spc.UpdateResponse
	if err := clt.Post(ctx, "v1.3/campaign/spc/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
