package changelog

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/changelog"
)

// TaskDownload Get the downloaded file
func TaskDownload(ctx context.Context, clt *core.SDKClient, req *changelog.TaskDownloadRequest, accessToken string) (*changelog.TaskDownloadResult, error) {
	var resp changelog.TaskDownloadResponse
	if err := clt.Get(ctx, "v1.3/changelog/task/download/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
