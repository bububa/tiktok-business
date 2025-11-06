package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/video"
)

// Settings 获取 TikTok 账号帖子隐私设置
func Settings(ctx context.Context, clt *core.SDKClient, req *video.SettingsRequest, accessToken string) (*video.Settings, error) {
	var resp video.SettingsResponse
	if err := clt.Get(ctx, "v1.3/business/video/settings/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
