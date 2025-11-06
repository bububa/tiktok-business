package ttvideo

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ttvideo"
)

// List 获取 Spark Ads 帖子列表
// 您可以使用本接口获取已授权给某一广告账户的 Spark Ads 帖子列表。
func List(ctx context.Context, clt *core.SDKClient, req *ttvideo.ListRequest, accessToken string) (*ttvideo.ListResult, error) {
	var resp ttvideo.ListResponse
	if err := clt.Get(ctx, "v1.3/tt_video/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
