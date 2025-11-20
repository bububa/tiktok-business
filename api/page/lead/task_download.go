package lead

import (
	"context"
	"io"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/page/lead"
)

// TaskDownload 下载线索
// 您可以使用本接口下载线索。您需要先通过/page/lead/task/ 接口创建下载任务并确保任务已完成（status = SUCCEED） 。
// 若接口调用成功，数据为UTC+0时间数据，且为 CSV (原始 CSV 文件 <= 10MB) 或 zip (原始 CSV 文件 > 10MB) 格式，而非 JSON 格式。
// 若接口产生错误，根据错误类型可能返回 JSON 格式或在返回的 CSV 文件中输出错误信息。
func TaskDownload(ctx context.Context, clt *core.SDKClient, req *lead.TaskDownloadRequest, accessToken string, w io.Writer) error {
	resp := lead.NewTaskDownloadResponse(w)
	if err := clt.Get(ctx, "v1.3/page/lead/task/download/", req, resp, accessToken); err != nil {
		return err
	}
	return nil
}
