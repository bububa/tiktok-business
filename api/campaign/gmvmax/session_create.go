package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// SessionCreate 创建最大投放量或创意作品加热时段
// 您可以使用本接口为商品 GMV Max 推广系列中的特定商品创建最大投放量时段，或为商品 GMV Max 推广系列中的特定视频创建创意作品加热时段。
// - GMV Max 的最大投放量模式通过追加预算，在选定时段内最大限度提升推广系列中所选商品的总收入。更多详情参考关于 GMV Max 的最大投放量优化。
// - 商品 GMV Max 推广系列中的创意作品加热功能允许商家手动为特定视频配置额外日预算，用于视频推广。
func SessionCreate(ctx context.Context, clt *core.SDKClient, req *gmvmax.SessionCreateRequest, accessToken string) (string, error) {
	var resp gmvmax.SessionCreateResponse
	if err := clt.Post(ctx, "v1.3/gmv_max/session/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.SessionID, nil
}
