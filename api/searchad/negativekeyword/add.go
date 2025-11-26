package negativekeyword

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/searchad/negativekeyword"
)

// Add 创建否定关键词
// 您可以使用本接口创建搜索广告的否定关键词。若用户的搜索词与否定关键词匹配，对应的推广系列或广告组中的广告不会显示。
// 您可以为广告组或推广系列设置否定关键词：
// 若想为广告组设置否定关键词，需确保该广告组已开启自动搜索广告位。
// 若想为推广系列设置否定关键词，需确保该推广系列下至少有一个已开启自动搜索广告位的广告组。
// 若想了解如何为广告组开启自动搜索广告位，请查看创建搜索广告。
func Add(ctx context.Context, clt *core.SDKClient, req *negativekeyword.AddRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/search_ad/negative_keyword/add/", req, nil, accessToken)
}
