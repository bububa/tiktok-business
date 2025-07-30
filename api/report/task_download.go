package report

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/report"
)

// TaskDownload 下载异步报表任务结果
func TaskDownload(ctx context.Context, clt *core.SDKClient, req *report.TaskDownloadRequest, accessToken string) (*report.TaskDownloadResult, error) {
	var ret report.TaskDownloadResponse
	if err := clt.Get(ctx, "v1.3/report/task/download/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
