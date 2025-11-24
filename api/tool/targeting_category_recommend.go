package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// TargetingCategoryRecommend 获取推荐兴趣及行为分类
// 您可以使用本接口获取系统推荐的兴趣和行为分类。推荐基于同行业的历史表现数据。
func TargetingCategoryRecommend(ctx context.Context, clt *core.SDKClient, req *tool.TargetingCategoryRecommendRequest, accessToken string) (*tool.TargetingCategoryRecommendResult, error) {
	var ret tool.TargetingCategoryRecommendResponse
	if err := clt.Get(ctx, "v1.3/tool/targeting_category/recommend/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
