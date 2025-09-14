package changelog

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model/changelog"
)

// TaskCheck Check the status of a download task
func TaskCheck(ctx context.Context, clt *core.SDKClient, req *changelog.TaskCheckRequest, accessToken string) (enum.ChangelogTaskStatus, error) {
	var resp changelog.TaskCheckResponse
	if err := clt.Get(ctx, "v1.3/changelog/task/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.Status, nil
}
