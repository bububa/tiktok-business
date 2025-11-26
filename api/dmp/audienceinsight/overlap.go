package audienceinsight

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/audienceinsight"
)

// Overlap 获取受众重叠详情
// 使用此接口获取基准受众与最多4个其他Tiktok自定义受众细分之间的重叠。
func Overlap(ctx context.Context, clt *core.SDKClient, req *audienceinsight.OverlapRequest, accessToken string) (*audienceinsight.Overlap, error) {
	var resp audienceinsight.OverlapResponse
	if err := clt.Get(ctx, "v1.3/audience/insight/overlap/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
