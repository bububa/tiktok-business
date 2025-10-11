package file

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file"
)

// FinishUpload 完成分片上传
func FinishUpload(ctx context.Context, clt *core.SDKClient, req *file.FinishUploadRequest, accessToken string) (*file.File, error) {
	var ret file.FinishUploadResponse
	if err := clt.Post(ctx, "v1.3/file/finish/upload/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
