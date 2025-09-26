package subscription

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/subscription"
)

// Subscribe 新建订阅
func Subscribe(ctx context.Context, clt *core.SDKClient, req *subscription.SubscribeRequest) (string, error) {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	var ret subscription.SubscribeResponse
	if err := clt.Post(ctx, "v1.3/subscription/subscribe/", req, &ret, ""); err != nil {
		return "", err
	}
	return ret.Data.SubscriptionID, nil
}
