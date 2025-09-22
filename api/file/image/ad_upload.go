package image

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model/file/image"
)

// AdUpload 上传图片
func AdUpload(ctx context.Context, clt *core.SDKClient, req *image.AdUploadRequest, accessToken string) (*image.Image, error) {
	var ret image.AdUploadResponse
	if req.UploadType == enum.UPLOAD_BY_FILE {
		if err := clt.Upload(ctx, "v1.3/file/image/ad/upload", req, &ret, accessToken); err != nil {
			return nil, err
		}
	} else {
		postReq := image.AdPostUploadRequest(*req)
		if err := clt.Post(ctx, "v1.3/file/image/ad/upload", &postReq, &ret, accessToken); err != nil {
			return nil, err
		}
	}
	return ret.Data, nil
}
