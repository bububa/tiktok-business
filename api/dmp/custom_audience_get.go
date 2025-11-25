package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// CustomAudienceGet 获取受众详情
// 您可以使用本接口获取指定受众群体的详细信息，包括受众群体当前状态以及历史操作记录详情。该接口将返回属于您的受众和与您分享的受众。您可以使用is_creator查看您是否为该受众的拥有者
func CustomAudienceGet(ctx context.Context, clt *core.SDKClient, req *dmp.CustomAudienceGetRequest, accessToken string) ([]dmp.CustomAudienceWithHistory, error) {
	var resp dmp.CustomAudienceGetResponse
	if err := clt.Get(ctx, "v1.3/dmp/custom_audience/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
