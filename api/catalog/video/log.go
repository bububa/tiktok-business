package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/video"
)

// Log 获取商品库视频的操作结果
// 您可以使用本接口了解商品库视频上传是否成功，以及上传失败的后续操作建议。
func Log(ctx context.Context, clt *core.SDKClient, req *video.LogRequest, accessToken string) (*video.LogResult, error) {
	var resp video.LogResponse
	if err := clt.Get(ctx, "v1.3/catalog/video/log/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
