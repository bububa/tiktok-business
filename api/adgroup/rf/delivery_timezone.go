package rf

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup/rf"
)

// DeliveryTimezone 获取覆盖和频次广告投放地区的时区信息
// 您可以使用本接口根据地点（地区代码）获取覆盖和频次广告的投放地区的时区信息（仅适用于覆盖和频次广告）。
func DeliveryTimezone(ctx context.Context, clt *core.SDKClient, req *rf.DeliveryTimezoneRequest, accessToken string) ([]rf.TimezoneInfo, error) {
	var resp rf.DeliveryTimezoneResponse
	if err := clt.Get(ctx, "v1.3/rf/delivery/timezone/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.TimezoneInfo, nil
}
