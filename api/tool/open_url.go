package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// OpenURL 获取 TikTok 应用内链接
// 您可以使用本接口获取 TikTok 公开网址的对应应用内链接。获取的地址可以用于创建覆盖和频次广告中，需要将地址传入广告的deeplink字段。
func OpenURL(ctx context.Context, clt *core.SDKClient, req *tool.OpenURLRequest, accessToken string) (*tool.OpenURLResult, error) {
	var ret tool.OpenURLResponse
	if err := clt.Get(ctx, "v1.3/tool/open_url/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
