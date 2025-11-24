package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// CustomAudienceList 获取受众列表
// 您可以使用本接口获取一个广告主账号下所有受众的信息。返回列表是分页的，最大分页大小为100。该接口返回的信息包括每个受众的ID、名称、类型、创建时间、过期时间等。该接口会返回属于您的受众和共享受众。您可以使用is_creator查看您是否为该受众的拥有者。
// 如果需要获取受众详细信息，例如当前状态以及历史操作记录详情，请使用/dmp/custom_audience/get/接口。
func CustomAudienceList(ctx context.Context, clt *core.SDKClient, req *dmp.CustomAudienceListRequest, accessToken string) (*dmp.CustomAudienceListResult, error) {
	var resp dmp.CustomAudienceListResponse
	if err := clt.Get(ctx, "v1.3/dmp/custom_audience/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
