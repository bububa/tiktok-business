package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// SegmentMapping 添加或删除受众细分群映射
// 您可以使用本接口添加或删除受众细分群的映射。支持的映射包括单一ID映射（每个用户或设备对应一个受众ID）和多ID映射（每个用户或设备对应多个受众ID）。
// 请查看流式API - 流式API相关提示，了解同一个企业号下有多个广告账号，和电子邮件地址标准化的相关提示。
func SegmentMapping(ctx context.Context, clt *core.SDKClient, req *dmp.SegmentMappingRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/segment/mapping/", req, nil, accessToken)
}
