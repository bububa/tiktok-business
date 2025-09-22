package adspreview

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/creative/adspreview"
)

// Create 广告预览
func Create(ctx context.Context, clt *core.SDKClient, req adspreview.ICreateRequest, accessToken string) (*adspreview.CreateResult, error) {
	var ret adspreview.CreateResponse
	if err := clt.Post(ctx, "v1.3/creative/ads_preview/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
