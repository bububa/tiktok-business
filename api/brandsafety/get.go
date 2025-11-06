package brandsafety

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/brandsafety"
)

// Get 获取广告账户的品牌安全中心设置
// 您可以使用本接口检索广告账户的品牌安全中心设置。
// 品牌安全中心允许您定义默认的广告账户层级的第一方品牌安全设置，包括 TikTok 库存筛选和适用性控制措施，例如类别排除和行业敏感内容。
func Get(ctx context.Context, clt *core.SDKClient, req *brandsafety.GetRequest, accessToken string) (*brandsafety.BrandSafetySetting, error) {
	var ret brandsafety.GetResponse
	if err := clt.Get(ctx, "v1.3/tiktok_inventory_filters/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
