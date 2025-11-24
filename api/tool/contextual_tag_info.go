package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// ContextualTagInfo 获取内容定向标签详情
// 您可以使用本接口获取内容相关定向标签的详细信息。您需在请求中传入广告主ID和/tool/contextual_tag/get/返回的内容相关定向标签ID。
func ContextualTagInfo(ctx context.Context, clt *core.SDKClient, req *tool.ContextualTagInfoRequest, accessToken string) ([]tool.ContextualTag, error) {
	var ret tool.ContextualTagInfoResponse
	if err := clt.Get(ctx, "v1.3/tool/contextual_tag/info/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.ContextualTagList, nil
}
