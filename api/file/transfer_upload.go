package file

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file"
)

// TransferUpload 请求分片上传
func TransferUpload(ctx context.Context, clt *core.SDKClient, req *file.TransferUploadRequest, accessToken string) (*file.TransferUploadResult, error) {
	var ret file.TransferUploadResponse
	if err := clt.Upload(ctx, "v1.3/file/transfer/upload", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
