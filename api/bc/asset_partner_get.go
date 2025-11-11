package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetPartnerGet 根据资产获取合作伙伴
// 商务中心的管理员用户可以使用本接口获取某个资产曾经被分享给的合作伙伴
func AssetPartnerGet(ctx context.Context, clt *core.SDKClient, req *bc.AssetPartnerGetRequest, accessToken string) (*bc.AssetPartnerGetResult, error) {
	var resp bc.AssetPartnerGetResponse
	if err := clt.Get(ctx, "v1.3/bc/asset/partner/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
