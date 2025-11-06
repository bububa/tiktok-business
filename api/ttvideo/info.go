package ttvideo

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ttvideo"
)

// Info 获取 Spark Ads 帖子信息
// 您可以使用本接口通过从帖子所有者处获得的授权码，查询有权发布为广告的 Spark Ads 帖子（视频）信息。
func Info(ctx context.Context, clt *core.SDKClient, req *ttvideo.InfoRequest, accessToken string) (*ttvideo.Item, error) {
	var resp ttvideo.InfoResponse
	if err := clt.Get(ctx, "v1.3/tt_video/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
