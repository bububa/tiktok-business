package webhook

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/webhook"
)

// Update 创建或更新 TikTok 账号 Webhook 配置
// 您可以使用本接口创建或更新开发者应用层级与您的 TikTok 账号活动相关的 Webhook 事件配置。
// 若想了解如何创建这些 Webhook 事件配置，请查看通过 Webhook API 订阅 TikTok 账户的 Webhook 事件。
// 若不再想收到相关通知，可使用/business/webhook/delete/删除 Webhook 配置。
func Update(ctx context.Context, clt *core.SDKClient, req *webhook.UpdateRequest) (*webhook.Webhook, error) {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	var resp webhook.UpdateResponse
	if err := clt.Post(ctx, "v1.3/business/webhook/update/", req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
