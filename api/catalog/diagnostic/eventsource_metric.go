package diagnostic

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/diagnostic"
)

// EventSourceMetric 获取商品库事件趋势及匹配率
// 您可以使用本接口获取接收到的应用事件或 Pixel 事件的趋势，可应用于再营销的事件数量，以及商品库匹配率。
func EventSourceMetric(ctx context.Context, clt *core.SDKClient, req *diagnostic.EventSourceMetricRequest, accessToken string) ([]diagnostic.EventSourceMetric, error) {
	var resp diagnostic.EventSourceMetricResponse
	if err := clt.Get(ctx, "v1.3/diagnostic/catalog/eventsource/metric", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
