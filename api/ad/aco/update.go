package aco

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad/aco"
)

// Update 更新智能创意广告素材 API Request
func Update(ctx context.Context, clt *core.SDKClient, req *aco.UpdateRequest, accessToken string) (*aco.Ad, error) {
	var ret aco.UpdateResponse
	if err := clt.Post(ctx, "v1.3/ad/aco/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
