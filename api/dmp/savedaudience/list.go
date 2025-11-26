package savedaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/savedaudience"
)

// List 获取已保存受众详情
// 您可以使用本接口获取与某个广告账户绑定的已保存受众详情。
func List(ctx context.Context, clt *core.SDKClient, req *savedaudience.ListRequest, accessToken string) ([]savedaudience.Audience, error) {
	var resp savedaudience.ListResponse
	if err := clt.Get(ctx, "v1.3/dmp/saved_audience/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.SavedAudiences, nil
}
