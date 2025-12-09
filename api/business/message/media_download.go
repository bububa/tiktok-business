package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// MediaDownload Download an image or a video from a message
func MediaDownload(ctx context.Context, clt *core.SDKClient, req *message.MediaDownloadRequest, accessToken string) (string, error) {
	var resp message.MediaDownloadResponse
	if err := clt.Post(ctx, "v1.3/business/message/media/download/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.DownloadURL, nil
}
