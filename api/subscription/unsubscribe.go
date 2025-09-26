package subscription

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/subscription"
)

// Unsubscribe 取消订阅
func Unsubscribe(ctx context.Context, clt *core.SDKClient, subscriptionID string) (string, error) {
	req := subscription.UnsubscribeRequest{
		AppID:          clt.AppID(),
		Secret:         clt.Secret(),
		SubscriptionID: subscriptionID,
	}
	var ret subscription.SubscribeResponse
	if err := clt.Post(ctx, "v1.3/subscription/unsubscribe/", &req, &ret, ""); err != nil {
		return "", err
	}
	return ret.Data.SubscriptionID, nil
}
