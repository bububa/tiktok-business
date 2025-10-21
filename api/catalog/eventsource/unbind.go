package eventsource

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/eventsource"
)

// Unbind 解绑商品库与事件源
// 您可以使用本接口将商务中心的商品库与事件源解绑。
func Unbind(ctx context.Context, clt *core.SDKClient, req *eventsource.BindRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/eventsource/unbind/", req, nil, accessToken)
}
