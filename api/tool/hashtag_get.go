package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// HashtagGet 通过 ID 获取定向话题标签
// 您可以使用本接口传入话题标签关键词ID，来获取相关的话题标签以及可用状态。
func HashtagGet(ctx context.Context, clt *core.SDKClient, req *tool.HashtagGetRequest, accessToken string) ([]tool.Keyword, error) {
	var ret tool.HashtagGetResponse
	if err := clt.Get(ctx, "v1.3/tool/hashtag/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.KeywordsStatus, nil
}
