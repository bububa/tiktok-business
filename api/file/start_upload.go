package file

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file"
)

// StartUpload 请求分片上传
func StartUpload(ctx context.Context, clt *core.SDKClient, req *file.StartUploadRequest, accessToken string) (*file.StartUploadResult, error) {
	var ret file.StartUploadResponse
	if err := clt.Post(ctx, "v1.3/file/start/upload/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
