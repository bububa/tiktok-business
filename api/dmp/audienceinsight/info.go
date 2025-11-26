package audienceinsight

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/audienceinsight"
)

// Info 获取潜在受众详情
func Info(ctx context.Context, clt *core.SDKClient, req *audienceinsight.InfoRequest, accessToken string) (*audienceinsight.Info, error) {
	var resp audienceinsight.InfoResponse
	if err := clt.Post(ctx, "v1.3/audience/insight/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
