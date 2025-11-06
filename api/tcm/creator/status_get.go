package creator

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/creator"
)

// StatusGet 查询创作者 TTCM 状态
// 您可以使用本接口按用户名获取创作者 的TTCM（TikTok Creator Marketplace）入驻状态列表。
func StatusGet(ctx context.Context, clt *core.SDKClient, req *creator.StatusGetRequest, accessToken string) ([]creator.OnboardingStatus, error) {
	var resp creator.StatusGetResponse
	if err := clt.Get(ctx, "v1.3/tcm/creator/status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.OnboardingStatus, nil
}
