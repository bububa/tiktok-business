package business

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business"
)

// PublishStatus 获取 TikTok 帖子的发布状态
func PublishStatus(ctx context.Context, clt *core.SDKClient, req *business.PublishStatusRequest, accessToken string) (*business.PublishStatus, error) {
	var resp business.PublishStatusResponse
	if err := clt.Get(ctx, "v1.3/business/publish/status/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
