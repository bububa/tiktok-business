package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// ActionCategory 获取行为分类
// 您可以使用本接口获取行为分类的枚举值。
// 通过行为分类，广告主可以根据用户最近的应用内行为向用户推送相关联的广告内容。用户的应用内行为可分为两类，视频相关行为（例如浏览、评论或分享），以及创作者相关行为（关注或浏览创作者主页）。相应地，开发者调用/adgroup/create/时能够使用的action_scene为CREATOR_RELATED或VIDEO_RELATED（只能选一个）。
func ActionCategory(ctx context.Context, clt *core.SDKClient, req *tool.ActionCategoryRequest, accessToken string) ([]tool.ActionCategory, error) {
	var ret tool.ActionCategoryResponse
	if err := clt.Get(ctx, "v1.3/tool/action_category/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.ActionCategories, nil
}
