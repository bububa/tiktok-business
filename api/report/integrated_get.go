package report

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/report"
)

// IntegratedGet 请求同步报表
// 您可以使用本接口创建同步报表任务。
// 本接口目前最多能返回 20,000 条广告的报表数据。若您的广告数目超过 20,000 条，请使用campaign_ids/adgroup_ids/ad_ids作为筛选条件，分批获取全部广告的报表数据。若您使用campaign_ids/adgroup_ids/ad_ids作为筛选条件，单次最多传入 100 个 ID。
func IntegratedGet(ctx context.Context, clt *core.SDKClient, req *report.IntegratedGetRequest, accessToken string) (*report.IntegratedGetResult, error) {
	var ret report.IntegratedGetResponse
	if err := clt.Get(ctx, "v1.3/report/integrated/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
