package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/video"
)

// Publish 将公开视频帖子发布到自有账号
// 您可以使用本接口将公开视频帖子发布到 TikTok 账号。
// 您需指定需要发布的视频，并可同时设置视频标题，以及是否允许不同类型的用户互动形式（评论/画面拼贴/合拍）。
// 若想了解可订阅的帖子发布相关的 Webhook 事件，可查看帖子发布事件。若想为您的开发者应用配置回调 URL，可使用 Webhook API 接口 /business/webhook/update/。
func Publish(ctx context.Context, clt *core.SDKClient, req *video.PublishRequest, accessToken string) (string, error) {
	var resp video.PublishResponse
	if err := clt.Post(ctx, "v1.3/business/video/publish/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.ShareID, nil
}
