package pixel

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/pixel"
)

// EventStats 获取Pixel事件数据
func EventStats(ctx context.Context, clt *core.SDKClient, req *pixel.EventStatsRequest, accessToken string) ([]pixel.PixelEventStats, error) {
	var ret pixel.EventStatsResponse
	if err := clt.Get(ctx, "v1.3/pixel/event/stats/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.List, nil
}
