package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// ContextualTagGet 获取可用内容定向标签
// 您可以使用本接口获取可用于定向的内容相关定向标签。您需要在请求中传入广告主ID，推广目标和地区代码。
func ContextualTagGet(ctx context.Context, clt *core.SDKClient, req *tool.ContextualTagGetRequest, accessToken string) ([]tool.ContextualTag, error) {
	var ret tool.ContextualTagGetResponse
	if err := clt.Get(ctx, "v1.3/tool/contextual_tag/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.ContextualTagList, nil
}
