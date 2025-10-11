package video

import (
	"context"
	"errors"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model/file/video"
)

// AdUpload 上传视频
func AdUpload(ctx context.Context, clt *core.SDKClient, req *video.AdUploadRequest, accessToken string) ([]video.Video, error) {
	var ret video.AdUploadResponse
	if req.UploadType == enum.UPLOAD_BY_FILE {
		if req.VideoFile == nil {
			return nil, errors.New("missing video_file")
		}
		if err := clt.Upload(ctx, "v1.3/file/video/ad/upload/", req, &ret, accessToken); err != nil {
			return nil, err
		}
	} else {
		postReq := video.AdPostUploadRequest(*req)
		if err := clt.Post(ctx, "v1.3/file/video/ad/upload/", &postReq, &ret, accessToken); err != nil {
			return nil, err
		}
	}
	return ret.Data, nil
}
