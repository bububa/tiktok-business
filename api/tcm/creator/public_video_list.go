package creator

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/creator"
)

// PublicVideoList 获取公开媒体洞察
// 您可以使用本接口获取 TTCM（TikTok Creator Marketplace）创作者视频的相关信息。
func PublicVideoList(ctx context.Context, clt *core.SDKClient, req *creator.PublicVideoListRequest, accessToken string) (*creator.PublicVideoListResult, error) {
	var resp creator.PublicVideoListResponse
	if err := clt.Get(ctx, "v1.3/tcm/creator/public/video/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
