package creator

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/creator"
)

// AuthorizedVideoList 授权后获取媒体洞察
// 您可以使用本接口获取 TikTok 创作者视频的报告洞察数据。
func AuthorizedVideoList(ctx context.Context, clt *core.SDKClient, req *creator.AuthorizedVideoListRequest, accessToken string) (*creator.AuthorizedVideoListResult, error) {
	var resp creator.AuthorizedVideoListResponse
	if err := clt.Get(ctx, "v1.3/tcm/creator/authorized/video/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
