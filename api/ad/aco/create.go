package aco

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad/aco"
)

// Create 创建智能创意广告 API Request
func Create(ctx context.Context, clt *core.SDKClient, req *aco.CreateRequest, accessToken string) (*aco.Ad, error) {
	var ret aco.CreateResponse
	if err := clt.Post(ctx, "v1.3/ad/aco/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
