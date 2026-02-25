package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/comment"
)

// ImageUpload 上传评论图片
// 此接口可用于为新评论或回复现有评论上传图片，返回的 image_uri 可在调用 /business/comment/create/ 或 /business/comment/reply/create/ 接口时一并传入。
func ImageUpload(ctx context.Context, clt *core.SDKClient, req *comment.ImageUploadRequest, accessToken string) (*comment.ImageUploadResult, error) {
	var resp comment.ImageUploadResponse
	if err := clt.Upload(ctx, "v1.3/business/comment/image/upload/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
