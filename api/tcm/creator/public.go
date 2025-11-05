package creator

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/creator"
)

// Public 获取公开账户洞察
// 您可以使用本接口获取 TTCM（TikTok Creator Marketplace）创作者的主页基本信息。
func Public(ctx context.Context, clt *core.SDKClient, req *creator.PublicRequest, accessToken string) (*creator.Creator, error) {
	var resp creator.PublicResponse
	if err := clt.Get(ctx, "v1.3/tcm/creator/public/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
