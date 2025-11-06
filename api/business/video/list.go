package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/video"
)

// List 获取 TikTok 账号视频数据
// 您可以使用本接口获取 TikTok 账号公开的视频帖子的覆盖人数和互动数据。返回的视频帖子列表包括展示在账户主页的原生视频，以及"只作为广告展示"的帖子（即专门用作公开可访问的广告、但不展示在账户主页的视频帖子）。
// 结果会分页显示。您可以使用 max_count 请求参数，指定返回中每页返回的视频最大数量。
func List(ctx context.Context, clt *core.SDKClient, req *video.ListRequest, accessToken string) (*video.ListResult, error) {
	var resp video.ListResponse
	if err := clt.Get(ctx, "v1.3/business/videos/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
