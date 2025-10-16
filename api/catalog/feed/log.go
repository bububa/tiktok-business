package feed

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/feed"
)

// Log 获取更新源日志
// 您可以使用本接口获取最近10次更新源的操作信息。
func Log(ctx context.Context, clt *core.SDKClient, req *feed.LogRequest, accessToken string) ([]feed.Log, error) {
	var ret feed.LogResponse
	if err := clt.Get(ctx, "v1.3/catalog/feed/log/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.FeedLogs, nil
}
