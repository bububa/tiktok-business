package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// SearchKeywordRecommend 获取推荐搜索关键词
// 您可以使用本接口获取搜索广告推广系列可用的推荐搜索关键词列表，以及关键词对应的预测月搜索量。
func SearchKeywordRecommend(ctx context.Context, clt *core.SDKClient, req *tool.SearchKeywordRecommendRequest, accessToken string) (*tool.SearchKeywordRecommendResult, error) {
	var ret tool.SearchKeywordRecommendResponse
	if err := clt.Get(ctx, "v1.3/tool/search_keyword/recommend/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
