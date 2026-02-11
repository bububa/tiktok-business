package campaign

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign"
)

// CopyTaskCheck 获取推广系列异步复制任务的结果
// 您可以使用本接口获取推广系列异步复制任务的结果。
func CopyTaskCheck(ctx context.Context, clt *core.SDKClient, req *campaign.CopyTaskCheckRequest, accessToken string) (*campaign.CopyTaskCheckResult, error) {
	var ret campaign.CopyTaskCheckResponse
	if err := clt.Get(ctx, "v1.3/campaign/copy/task/check/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
