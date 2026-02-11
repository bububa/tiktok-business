package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// SearchKeywordIdeaKeyword 发现新关键词
// 此接口可用于发现新关键词并探索关键词建议，触达对你的业务感兴趣的用户。
func SearchKeywordIdeaKeyword(ctx context.Context, clt *core.SDKClient, req *tool.SearchKeywordIdeaKeywordRequest, accessToken string) (*tool.SearchKeywordIdeaKeywordResult, error) {
	var ret tool.SearchKeywordIdeaKeywordResponse
	if err := clt.Get(ctx, "v1.3/tool/search_keyword/idea_keyword/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
