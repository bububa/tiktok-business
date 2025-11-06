package business

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business"
)

// HashtagSuggestion 为 TikTok 账号获取推荐话题标签
// 您可以使用本接口传入关键词，从而为您的TikTok 账号视频获取推荐话题标签列表。
// 返回的话题标签（name）可在使用/business/video/publish/将视频发布到 TikTok 账号时，用在caption字段中作为视频文案中的一部分。
func HashtagSuggestion(ctx context.Context, clt *core.SDKClient, req *business.HashtagSuggestionRequest, accessToken string) ([]business.Hashtag, error) {
	var resp business.HashtagSuggestionResponse
	if err := clt.Get(ctx, "v1.3/business/hashtag/suggestion/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Suggestions, nil
}
