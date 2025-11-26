package customaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/customaudience"
)

// Get 获取受众详情
// 您可以使用本接口获取指定受众群体的详细信息，包括受众群体当前状态以及历史操作记录详情。该接口将返回属于您的受众和与您分享的受众。您可以使用is_creator查看您是否为该受众的拥有者
func Get(ctx context.Context, clt *core.SDKClient, req *customaudience.GetRequest, accessToken string) ([]customaudience.AudienceWithHistory, error) {
	var resp customaudience.GetResponse
	if err := clt.Get(ctx, "v1.3/dmp/custom_audience/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
