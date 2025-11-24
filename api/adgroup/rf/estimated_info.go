package rf

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup/rf"
)

// EstimatedInfo 获取覆盖和频次广告组预估信息
// 您可以使用本接口查询下单时的预估每日成本和频次分布。若广告组没有关联有效的预估信息，则不返回对应记录。
func EstimatedInfo(ctx context.Context, clt *core.SDKClient, req *rf.EstimatedInfoRequest, accessToken string) ([]rf.EstimatedInfo, error) {
	var resp rf.EstimatedInfoResponse
	if err := clt.Get(ctx, "v1.3/adgroup/rf/estimated/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.EstimatedInfo, nil
}
