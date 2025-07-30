package report

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/report"
)

// TaskCreate 创建异步报表任务
func TaskCreate(ctx context.Context, clt *core.SDKClient, req *report.TaskCreateRequest, accessToken string) (string, error) {
	var ret report.TaskCreateResponse
	if err := clt.Post(ctx, "v1.3/report/task/create/", req, &ret, accessToken); err != nil {
		return "", err
	}
	return ret.Data.TaskID, nil
}
