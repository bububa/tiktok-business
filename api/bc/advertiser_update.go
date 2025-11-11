package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AdvertiserUpdate 更新广告账户
// 您可以使用本接口更新商务中心下的广告账户。您还可以使用本接口来更新自动分配模式下商务中心的广告账户月度预算。支持增量更新。您必须是商务中心管理员用户
func AdvertiserUpdate(ctx context.Context, clt *core.SDKClient, req *bc.AdvertiserUpdateRequest, accessToken string) ([]bc.AdvertiserUpdateResult, error) {
	var resp bc.AdvertiserUpdateResponse
	if err := clt.Post(ctx, "v1.3/bc/advertiser/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
