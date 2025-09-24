package pixel

import (
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 更新Pixel API Request
type UpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// PixelID Pixel ID
	PixelID string `json:"pixel_id,omitempty"`
	// PixelName Pixel 名称。
	// 名称不能超过40字符（不区分语言），不能包含emoji。
	PixelName string `json:"pixel_name,omitempty"`
	// AdvancedMathingFields Pixel 的手动高级匹配设置。
	// 此设置仅适用于通过合作伙伴集成生成的 Pixel，不适用于手动配置的 Pixel。
	AdvancedMatchingFields *AdvancedMatchingFields `json:"advanced_matching_fields,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
