package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// InterestCategory 获取一般兴趣分类列表
func InterestCategory(ctx context.Context, clt *core.SDKClient, req *tool.InterestCategoryRequest, accessToken string) ([]tool.InterestCategory, error) {
	var ret tool.InterestCategoryResponse
	if err := clt.Get(ctx, "v1.3/tool/interest_category/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.InterestCategories, nil
}
