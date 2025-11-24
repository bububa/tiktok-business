package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// HashtagRecommend 搜索定向话题标签
// 您可以使用本接口基于种子关键词搜索定向话题标签。
// 返回的话题标签 ID （keyword_id）可传入 /adgroup/create/ 中的 action_category_ids 字段，用于定向观看过带有该话题标签的 TikTok 视频的受众。您需同时将 action_scene 设置为 HASHTAG_RELATED。
func HashtagRecommend(ctx context.Context, clt *core.SDKClient, req *tool.HashtagRecommendRequest, accessToken string) ([]tool.RecommendedKeyword, error) {
	var ret tool.HashtagRecommendResponse
	if err := clt.Get(ctx, "v1.3/tool/hashtag/recommend/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.RecommendedKeywords, nil
}
