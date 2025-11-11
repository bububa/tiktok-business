package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AssetBindQuotaRequest 获取资产绑定信息 API Request
type AssetBindQuotaRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// AssetType 资产类型。
	// 枚举值：
	// ADVERTISER：广告账户（广告主账号）。
	// CATALOG：商品库。
	// TIKTOK_SHOP：TikTok Shop（商店）。
	// PIXEL：Pixel。
	// LEAD：线索。
	// TT_ACCOUNT：TikTok 账号。此资产类型对应 BC_AUTH_TT身份，详情见认证身份
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// AssetID 资产 ID。
	// 若 asset_type 为 ADVERTISER，本字段代表广告账户 ID。
	// 若 asset_type 为 CATALOG，本字段代表商品库 ID。
	// 若 asset_type 为 TIKTOK_SHOP，本字段代表 TikTok Shop 的 ID。
	// 若 asset_type 为 PIXEL，本字段代表 Pixel ID。
	// 若 asset_type 为 LEAD，本字段代表表单库 ID。若想获取您有权访问的表单库的详情，可使用/page/library/get/。
	// 若 asset_type 为 TT_ACCOUNT，本字段代表商务中心中的 TikTok 账号的 ID。
	AssetID string `json:"asset_id,omitempty"`
}

// Encode implements GetRequest
func (r *AssetBindQuotaRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("asset_type", string(r.AssetType))
	values.Set("asset_id", r.AssetID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AssetBindQuotaResponse 获取资产绑定信息 API Response
type AssetBindQuotaResponse struct {
	model.BaseResponse
	Data *AssetBindQuota `json:"data,omitempty"`
}

// AssetBindQuota 资产绑定信息
type AssetBindQuota struct {
	// TotalQuota 可绑定的总数
	TotalQuota int `json:"total_quota,omitempty"`
	// UsedQuota 已绑定数
	UsedQuota int `json:"used_quota,omitempty"`
	// AvailableQuota 剩余可用数
	AvailableQuota int `json:"available_quota,omitempty"`
}
