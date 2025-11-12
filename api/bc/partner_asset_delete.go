package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// PartnerAssetDelete 取消资产分享
// 商务中心的管理员用户可以使用本接口取消与某个合作伙伴分享资产
func PartnerAssetDelete(ctx context.Context, clt *core.SDKClient, req *bc.PartnerAssetDeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/partner/asset/delete/", req, nil, accessToken)
}
