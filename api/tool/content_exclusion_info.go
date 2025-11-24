package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// ContentExclusionInfo 获取内容排除类别信息
// 您可以使用本接口获取类别排除或行业敏感内容控制类型的详细信息，避免其出现在您的广告周围。
func ContentExclusionInfo(ctx context.Context, clt *core.SDKClient, req *tool.ContentExclusionInfoRequest, accessToken string) ([]tool.Category, error) {
	var ret tool.ContentExclusionInfoResponse
	if err := clt.Get(ctx, "v1.3/tool/content_exclusion/info/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.ContentExclusionList, nil
}
