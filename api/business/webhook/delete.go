package webhook

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/webhook"
)

// Delete 删除 TikTok 账号 Webhook 配置
// 您可使用本接口删除开发者应用层级已有的与您的 TikTok 账号活动相关的 Webhook 事件配置。
func Delete(ctx context.Context, clt *core.SDKClient, req *webhook.DeleteRequest) (*webhook.Webhook, error) {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	var resp webhook.DeleteResponse
	if err := clt.Post(ctx, "v1.3/business/webhook/delete/", req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
