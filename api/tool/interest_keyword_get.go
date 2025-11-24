package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// InterestKeywordGet 通过 ID 获取其他兴趣分类
func InterestKeywordGet(ctx context.Context, clt *core.SDKClient, req *tool.InterestKeywordGetRequest, accessToken string) ([]tool.Keyword, error) {
	var ret tool.InterestKeywordGetResponse
	if err := clt.Get(ctx, "v1.3/tool/interest_keyword/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.Keywords, nil
}
