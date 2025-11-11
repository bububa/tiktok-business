package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// SessionList 获取推广系列中的最大投放量或创意作品加热时段
// 您可以使用本接口获取商品 GMV Max 推广系列中生效的最大投放量或创意作品加热时段列表。
// 已经结束的时段不支持通过 /campaign/gmv_max/session/list/ 获取。若想获取此类时段，可使用 /campaign/gmv_max/session/get/。
func SessionList(ctx context.Context, clt *core.SDKClient, req *gmvmax.SessionListRequest, accessToken string) ([]gmvmax.Session, error) {
	var resp gmvmax.SessionListResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/session/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.SessionList, nil
}
