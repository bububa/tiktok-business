package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AdvertiserQualificationGet 获取商务中心中已有的资质信息
// 您可以使用本接口获取商务中心中已有资质的信息。
// 从返回中获取 qualification_id 后，您可将此字段传入 /bc/advertiser/create/ 的 qualification_id 复用该资质创建广告账户
func AdvertiserQualificationGet(ctx context.Context, clt *core.SDKClient, req *bc.AdvertiserQualificationGetRequest, accessToken string) (*bc.AdvertiserQualificationGetResult, error) {
	var resp bc.AdvertiserQualificationGetResponse
	if err := clt.Get(ctx, "v1.3/bc/advertiser/qualification/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
