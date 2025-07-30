package oauth

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/oauth"
)

// AdvertiserGet 获取广告账号列表
func AdvertiserGet(ctx context.Context, clt *core.SDKClient) ([]oauth.Advertiser, error) {
	req := oauth.AdvertiserGetRequest{
		AppID:  clt.AppID(),
		Secret: clt.Secret(),
	}
	var ret oauth.AdvertiserGetResponse
	if err := clt.Get(ctx, "v1.3/oauth2/advertiser/get/", &req, &ret, ""); err != nil {
		return nil, err
	}
	return ret.Data.List, nil
}
