package aco

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad/aco"
)

// Get 获取智能创意广告素材
func Get(ctx context.Context, clt *core.SDKClient, req *aco.GetRequest, accessToken string) ([]aco.Ad, error) {
	var ret aco.GetResponse
	if err := clt.Get(ctx, "v1.3/ad/aco/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.List, nil
}
