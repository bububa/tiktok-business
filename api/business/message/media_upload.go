package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// MediaUpload Upload an image
func MediaUpload(ctx context.Context, clt *core.SDKClient, req *message.MediaUploadRequest, accessToken string) (string, error) {
	var resp message.MediaUploadResponse
	if err := clt.Upload(ctx, "v1.3/business/message/media/upload/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.MediaID, nil
}
