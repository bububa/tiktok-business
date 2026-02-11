package campaign

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign"
)

// CopyTaskCreate 创建推广系列异步复制任务
// 您可以使用本接口创建推广系列异步复制任务。
// 每次仅支持复制一个推广系列。若想了解详细步骤，请查看复制推广系列。创建复制任务后，需使用 /campaign/copy/task/check/ 查询任务的结果。
func CopyTaskCreate(ctx context.Context, clt *core.SDKClient, req *campaign.CopyTaskCreateRequest, accessToken string) (*campaign.CopyTaskCreateResult, error) {
	var ret campaign.CopyTaskCreateResponse
	if err := clt.Post(ctx, "v1.3/campaign/copy/task/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
