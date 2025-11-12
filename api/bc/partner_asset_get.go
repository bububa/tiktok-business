package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// PartnerAssetGet 获取合作伙伴资产
// 商务中心的管理员用户可以使用本接口获取分享给合作伙伴的资产以及由合作伙伴分享的资产
func PartnerAssetGet(ctx context.Context, clt *core.SDKClient, req *bc.PartnerAssetGetRequest, accessToken string) (*bc.AssetGetResult, error) {
	var resp bc.AssetGetResponse
	if err := clt.Get(ctx, "v1.3/bc/partner/asset/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
