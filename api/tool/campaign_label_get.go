package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// CampaignLabelGet 获取广告账号的推广系列标签
func CampaignLabelGet(ctx context.Context, clt *core.SDKClient, req *tool.CampaignLabelGetRequest, accessToken string) (*tool.CampaignLabelGetResult, error) {
	var ret tool.CampaignLabelGetResponse
	if err := clt.Get(ctx, "v1.3/campaign_label/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
