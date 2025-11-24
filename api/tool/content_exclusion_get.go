package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// ContentExclusionGet 获取可用的内容排除类别
// 您可以使用本接口获取类别排除或行业敏感内容控制类型，避免其出现在您的广告附近。
func ContentExclusionGet(ctx context.Context, clt *core.SDKClient, req *tool.ContentExclusionGetRequest, accessToken string) (*tool.ContentExclusionGetResult, error) {
	var ret tool.ContentExclusionGetResponse
	if err := clt.Get(ctx, "v1.3/tool/content_exclusion/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
