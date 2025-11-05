package creator

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/creator"
)

// Get 通过访问令牌获取创作者
// 您可以使用本接口获取授权该 TikTok for Business 应用的创作者。
func Get(ctx context.Context, clt *core.SDKClient, accessToken string) ([]creator.Creator, error) {
	var resp creator.GetResponse
	if err := clt.Get(ctx, "v1.3/oauth2/creator/get/", nil, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
