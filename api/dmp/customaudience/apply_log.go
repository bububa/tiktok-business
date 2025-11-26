package customaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/customaudience"
)

// ApplyLog 获取受众应用记录
// 您可以使用本接口获取自定义受众的最新应用记录。请注意，本接口仅返回应用操作（action_mode = Apply）的记录。
func ApplyLog(ctx context.Context, clt *core.SDKClient, req *customaudience.ApplyLogRequest, accessToken string) (*customaudience.ApplyLogResult, error) {
	var resp customaudience.ApplyLogResponse
	if err := clt.Get(ctx, "v1.3/dmp/custom_audience/apply/log/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
