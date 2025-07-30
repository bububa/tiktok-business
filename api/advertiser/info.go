package advertiser

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/advertiser"
)

// Info 获取广告账号信息
func Info(ctx context.Context, clt *core.SDKClient, req *advertiser.InfoRequest, accessToken string) ([]advertiser.Advertiser, error) {
	var ret advertiser.InfoResponse
	if err := clt.Get(ctx, "v1.3/advertiser/info/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.List, nil
}
