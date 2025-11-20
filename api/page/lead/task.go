package lead

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/page/lead"
)

// Task 创建线索下载任务
// 您可以使用本接口为您在请求参数中传入的page_id或ad_id产生的所有线索创建下载任务。
// 异步下载任务创建后，再次调用/page/lead/task/检查任务是否完成（status = SUCCEED）。若任务已完成，调用/page/lead/task/download/下载线索。
func Task(ctx context.Context, clt *core.SDKClient, req *lead.TaskRequest, accessToken string) (*lead.Task, error) {
	var resp lead.TaskResponse
	if err := clt.Post(ctx, "v1.3/page/lead/task/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
