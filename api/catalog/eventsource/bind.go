package eventsource

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/eventsource"
)

// Bind 绑定商品库与事件源
// 您可以使用本接口将商务中心的商品库与事件源（应用事件源或网站Pixel事件源）进行绑定，以追踪相关数据。
func Bind(ctx context.Context, clt *core.SDKClient, req *eventsource.BindRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/eventsource/bind/", req, nil, accessToken)
}
