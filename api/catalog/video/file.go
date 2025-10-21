package video

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/video"
)

// File 通过文件 URL 上传商品库视频
// 您可以使用本接口以文件 URL 的方式上传商品库视频，商品库视频将通过异步任务与电商商品库绑定。您将会从接口返回中得到一个任务处理 ID（feed_log_id）。通过该ID您可以调用 /catalog/video/log/ 查看任务的处理状态和最终结果。
// 目前本接口只允许上传包含视频 URL 的 CSV 格式文件。 商品库视频上传的 CSV 模板为catalog_video_template。
func File(ctx context.Context, clt *core.SDKClient, req *video.FileRequest, accessToken string) (string, error) {
	var resp video.FileResponse
	if err := clt.Post(ctx, "v1.3/catalog/video/file/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.FeedLogID, nil
}
