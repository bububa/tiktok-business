package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// TargetingInfo 根据 ID 获取地域定向标签的信息
// 您可以使用本接口根据 ID 获取定向标签的信息。
// 您可获取三种定向标签的详情：邮政编码 ID（目前仅支持美国、加拿大、巴西、印度尼西亚、泰国、越南）、地域 ID 或互联网服务提供商 ID。
// 请查看下文的“不同的定向标签所需的参数设置”小节，了解获取不同定向标签时需要传入的参数的对比。
func TargetingInfo(ctx context.Context, clt *core.SDKClient, req *tool.TargetingInfoRequest, accessToken string) (*tool.TargetingInfoResult, error) {
	var ret tool.TargetingInfoResponse
	if err := clt.Get(ctx, "v1.3/tool/targeting/info/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
