package creator

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/creator"
)

// Discover 发现创作者
// 您可以使用本接口根据粉丝数、平均视频播放量、受众年龄及其他筛选条件搜索创作者。您还可以使用粉丝数、平均视频播放量和互动率对返回的结果进行排序
func Discover(ctx context.Context, clt *core.SDKClient, req *creator.DiscoverRequest, accessToken string) (*creator.DiscoverResult, error) {
	var resp creator.DiscoverResponse
	if err := clt.Post(ctx, "v1.3/tcm/creator/discover/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
