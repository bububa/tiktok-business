package aco

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad/aco"
)

// StatusUpdate 更新广告操作状态 API Request
func StatusUpdate(ctx context.Context, clt *core.SDKClient, req *aco.StatusUpdateRequest, accessToken string) (*aco.StatusUpdateResult, error) {
	var ret aco.StatusUpdateResponse
	if err := clt.Post(ctx, "v1.3/ad/aco/material_status/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
