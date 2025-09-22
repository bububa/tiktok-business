package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// DiagnosisGet 获取广告组诊断
func DiagnosisGet(ctx context.Context, clt *core.SDKClient, req *tool.DiagnosisGetRequest, accessToken string) ([]tool.DiagnosisGetResult, error) {
	var ret tool.DiagnosisGetResponse
	if err := clt.Get(ctx, "v1.3/tool/diagnosis/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.Results, nil
}
