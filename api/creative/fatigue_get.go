package creative

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/creative"
)

// FatigueGet 获取创意疲劳检测结果
// 您可以使用本接口检测广告在过去的某个时间区间是否发生了创意疲劳。
// 若您希望通过Webhook订阅单个广告、广告组中所有广告或广告账户下所有广告的疲劳状态，请查看订阅Webhook-订阅广告疲劳状态小节。
func FatigueGet(ctx context.Context, clt *core.SDKClient, req *creative.FatigueGetRequest, accessToken string) (*creative.FatigueGetResult, error) {
	var resp creative.FatigueGetResponse
	if err := clt.Get(ctx, "v1.3/creative_fatigue/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
