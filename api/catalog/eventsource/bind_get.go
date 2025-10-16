package eventsource

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/eventsource"
)

// BindGet 获取事件源绑定信息
// 您可以使用本接口获取商务中心下商品库的事件源绑定信息。
func BindGet(ctx context.Context, clt *core.SDKClient, req *eventsource.BindGetRequest, accessToken string) (*eventsource.BindGetResult, error) {
	var ret eventsource.BindGetResponse
	if err := clt.Get(ctx, "v1.3/eventsource_bind/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
