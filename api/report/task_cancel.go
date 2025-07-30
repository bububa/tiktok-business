package report

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model/report"
)

// TaskCancel 取消异步报表任务
func TaskCancel(ctx context.Context, clt *core.SDKClient, req *report.TaskCancelRequest, accessToken string) (enum.ReportTaskStatus, error) {
	var ret report.TaskCancelResponse
	if err := clt.Post(ctx, "v1.3/report/task/cancel/", req, &ret, accessToken); err != nil {
		return "", err
	}
	return ret.Data.Status, nil
}
