package file

import (
	"context"
	"errors"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model/file"
)

// TemporarilyUpload 上传单个文件
func TemporarilyUpload(ctx context.Context, clt *core.SDKClient, req *file.TemporarilyUploadRequest, accessToken string) (*file.File, error) {
	var ret file.TemporarilyUploadResponse
	if req.UploadType == enum.UploadType_FILE {
		if req.File == nil {
			return nil, errors.New("missing file")
		}
		if err := clt.Upload(ctx, "v1.3/file/temporarily/upload/", req, &ret, accessToken); err != nil {
			return nil, err
		}
	} else {
		postReq := file.TemporarilyPostUploadRequest(*req)
		if err := clt.Post(ctx, "v1.3/file/temporarily/upload/", &postReq, &ret, accessToken); err != nil {
			return nil, err
		}
	}
	return ret.Data, nil
}
