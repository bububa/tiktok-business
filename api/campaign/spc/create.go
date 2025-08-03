package spc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/spc"
)

// Create 创建Smart+推广系列 API Request
func Create(ctx context.Context, clt *core.SDKClient, req *spc.CreateRequest, accessToken string) (*spc.Campaign, error) {
	var ret spc.CreateResponse
	if err := clt.Post(ctx, "v1.3/campaign/spc/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
