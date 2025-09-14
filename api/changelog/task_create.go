package changelog

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/changelog"
)

// TaskCreate 创建更改日志下载任务
func TaskCreate(ctx context.Context, clt *core.SDKClient, req *changelog.TaskCreateRequest, accessToken string) (string, error) {
	var resp changelog.TaskCreateResponse
	if err := clt.Post(ctx, "v1.3/changelog/task/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.TaskID, nil
}
