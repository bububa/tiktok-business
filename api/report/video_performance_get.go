package report

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/report"
)

// VideoPerformanceGet 获取秒级表现数据
// 您可以使用本接口获取广告的秒级表现数据，或视频素材的视频洞察数据。
func VideoPerformanceGet(ctx context.Context, clt *core.SDKClient, req *report.VideoPerformanceGetRequest, accessToken string) (*report.VideoPerformanceGetResult, error) {
	var ret report.VideoPerformanceGetResponse
	if err := clt.Get(ctx, "v1.3/report/video_performance/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
