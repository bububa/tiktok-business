package spc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/spc"
)

// Get 获取 Smart+ 推广系列
func Get(ctx context.Context, clt *core.SDKClient, req *spc.GetRequest, accessToken string) ([]spc.Campaign, error) {
	var ret spc.GetResponse
	if err := clt.Get(ctx, "v1.3/campaign/spc/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.List, nil
}
