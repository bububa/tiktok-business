package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AdvertiserAttribute 获取广告账户的币种和注册地区
// 您可以使用本接口获取商务中心内广告账户的币种和注册地区。
// 您可将返回的币种和注册地区传给商务中心报表中的筛选字段registered_area和currency_of_account
func AdvertiserAttribute(ctx context.Context, clt *core.SDKClient, req *bc.AdvertiserAttributeRequest, accessToken string) (*bc.AdvertiserAttribute, error) {
	var resp bc.AdvertiserAttributeResponse
	if err := clt.Get(ctx, "v1.3/bc/advertiser/attribute/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
