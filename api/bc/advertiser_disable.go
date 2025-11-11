package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AdvertiserDisable 永久停用广告账户
// 您可以使用本接口永久停用商务中心中的广告账户。您必须是商务中心管理员用户。
// 推荐您不再需要使用某一广告账户时再进行永久停用账户的操作，因为一旦永久停用，广告账户将会被关闭，且广告账户将无法恢复到正常状态（即“审核通过”状态）。
func AdvertiserDisable(ctx context.Context, clt *core.SDKClient, req *bc.AdvertiserDisableRequest, accessToken string) (*bc.AdvertiserDisableResult, error) {
	var resp bc.AdvertiserDisableResponse
	if err := clt.Post(ctx, "v1.3/bc/advertiser/disable/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
