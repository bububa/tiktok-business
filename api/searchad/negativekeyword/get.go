package negativekeyword

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/searchad/negativekeyword"
)

// Get 获取否定关键词
// 您可以使用本接口获取搜索广告的否定关键词列表。
func Get(ctx context.Context, clt *core.SDKClient, req *negativekeyword.GetRequest, accessToken string) (*negativekeyword.GetResult, error) {
	var resp negativekeyword.GetResponse
	if err := clt.Get(ctx, "v1.3/search_ad/negative_keyword/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
