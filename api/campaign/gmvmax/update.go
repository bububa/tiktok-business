package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// Update 更新 GMV Max 推广系列
// 您可以使用本接口更新 GMV Max 推广系列。
// 本接口支持增量更新
func Update(ctx context.Context, clt *core.SDKClient, req *gmvmax.UpdateRequest, accessToken string) (*gmvmax.Campaign, error) {
	var resp gmvmax.UpdateResponse
	if err := clt.Post(ctx, "v1.3/campaign/gmv_max/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
