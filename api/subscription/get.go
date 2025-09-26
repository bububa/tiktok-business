package subscription

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/subscription"
)

// Get 获取开发者应用的订阅信息
func Get(ctx context.Context, clt *core.SDKClient, req *subscription.GetRequest) (*subscription.GetResult, error) {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	var ret subscription.GetResponse
	if err := clt.Get(ctx, "v1.3/subscription/get/", req, &ret, ""); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
