package photo

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/photo"
)

// Publish 将照片帖子发布到自有账号
// 您可以使用本接口将照片帖子发布到自有 TikTok 账号。
// 若想了解可订阅的帖子发布相关的 Webhook 事件，可查看帖子发布事件。若想为您的开发者应用配置回调 URL，可使用 Webhook API 接口 /business/webhook/update/。
func Publish(ctx context.Context, clt *core.SDKClient, req *photo.PublishRequest, accessToken string) (string, error) {
	var resp photo.PublishResponse
	if err := clt.Post(ctx, "v1.3/business/photo/publish/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.ShareID, nil
}
