package music

import (
	"context"
	"errors"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model/file/music"
)

// Upload 上传视频
func Upload(ctx context.Context, clt *core.SDKClient, req *music.UploadRequest, accessToken string) (*music.Music, error) {
	var ret music.UploadResponse
	if req.UploadType == enum.UPLOAD_BY_FILE {
		if req.MusicFile == nil {
			return nil, errors.New("missing music_file")
		}
		if err := clt.Upload(ctx, "v1.3/file/music/upload/", req, &ret, accessToken); err != nil {
			return nil, err
		}
	} else {
		postReq := music.PostUploadRequest(*req)
		if err := clt.Post(ctx, "v1.3/file/music/upload/", &postReq, &ret, accessToken); err != nil {
			return nil, err
		}
	}
	return ret.Data, nil
}
