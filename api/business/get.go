package business

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business"
)

// Get 获取 TikTok 账号主页数据
// 您可以使用本接口获取有关 TikTok 账号粉丝群体和主页互动的详细分析和洞察。
func Get(ctx context.Context, clt *core.SDKClient, req *business.GetRequest, accessToken string) (*business.Insight, error) {
	var resp business.GetResponse
	if err := clt.Get(ctx, "v1.3/business/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
