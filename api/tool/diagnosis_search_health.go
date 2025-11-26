package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// DiagnosisSearchHealth 获取搜索广告推广系列的健康状况诊断
// 您可以使用本接口获取搜索广告推广系列中广告组或广告的健康状况诊断信息，包括搜索量、关键词相关性以及出价和预算建议。
func DiagnosisSearchHealth(ctx context.Context, clt *core.SDKClient, req *tool.DiagnosisSearchHealthRequest, accessToken string) (*tool.DiagnosisSearchHealthResult, error) {
	var ret tool.DiagnosisSearchHealthResponse
	if err := clt.Get(ctx, "v1.3/tool/diagnosis/search/health/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
