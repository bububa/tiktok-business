package spc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/spc"
)

// QuotaGet 获取投放中 Smart+ 推广系列的动态配额
func QuotaGet(ctx context.Context, clt *core.SDKClient, req *spc.QuotaGetRequest, accessToken string) (*spc.Quota, error) {
	var ret spc.QuotaGetResponse
	if err := clt.Get(ctx, "v1.3/campaign/spc/quota/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
