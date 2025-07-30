package report

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/report"
)

// TaskCheck 获取异步报表任务状态
func TaskCheck(ctx context.Context, clt *core.SDKClient, req *report.TaskCheckRequest, accessToken string) (*report.TaskCheckResult, error) {
	var ret report.TaskCheckResponse
	if err := clt.Get(ctx, "v1.3/report/task/check/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
