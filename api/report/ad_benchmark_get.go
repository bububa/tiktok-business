package report

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/report"
)

// AdBenchmarkGet 创意基础报表
func AdBenchmarkGet(ctx context.Context, clt *core.SDKClient, req *report.AdBenchmarkGetRequest, accessToken string) (*report.AdBenchmarkGetResult, error) {
	var ret report.AdBenchmarkGetResponse
	if err := clt.Get(ctx, "v1.3/report/ad_benchmark/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
