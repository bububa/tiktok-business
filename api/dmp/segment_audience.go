package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// SegmentAudience 创建/删除受众细分群
// 您可以使用本接口创建/删除受众细分群。
// 请查看流式API - 流式API相关提示，了解同一个企业号下有多个广告账号，在受众细分中混用IDFA/AAID加密类型，和电子邮件地址标准化的相关提示
func SegmentAudience(ctx context.Context, clt *core.SDKClient, req *dmp.SegmentAudienceRequest, accessToken string) (string, error) {
	var resp dmp.SegmentAudienceResponse
	if err := clt.Post(ctx, "v1.3/segment/audience/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.AudienceID, nil
}
