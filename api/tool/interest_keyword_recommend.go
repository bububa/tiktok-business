package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// InterestKeywordRecommend 搜索其他兴趣分类
// 您可以使用本接口搜索其他兴趣分类。你需要传入一个或多个种子关键词，系统会对该种子关键词进行匹配，然后获取一组相关的其他兴趣分类。
// 返回中 status 为 EFFECTIVE 的其他兴趣分类（keyword_id）可在调用 /adgroup/create/ 创建广告组时传入 interest_keyword_ids 字段。指定的其他兴趣分类将在广告组中用于定向受众。
func InterestKeywordRecommend(ctx context.Context, clt *core.SDKClient, req *tool.InterestKeywordRecommendRequest, accessToken string) ([]tool.RecommendedKeyword, error) {
	var ret tool.InterestKeywordRecommendResponse
	if err := clt.Get(ctx, "v1.3/tool/interest_keyword/recommend/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.RecommendedKeywords, nil
}
