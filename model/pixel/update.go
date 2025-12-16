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
	// AutomaticAdvancedMatchingFields 该 Pixel 的自动高级匹配设置。
	// 自动高级匹配功能自动使用通过 TikTok Pixel 收集的客户信息，如电子邮箱、姓名和电话号码，从而改善匹配率、转化监测和再营销。
	// Note:为帮助保护客户隐私，敏感信息将先经过散列处理，再与 TikTok 共享。如果你是在受到较多监管的行业或敏感行业（例如金融服务或医疗保健行业）开展业务，请考虑使用手动高级匹配 (MAM)，而不是自动高级匹配 (AAM)。
	AutomaticAdvancedMatchingFields *AutomaticAdvancedMatchingFields `json:"automatic_advanced_matching_fields,omitempty"`
	// EnableFirstPartyCookies 是否允许第一方 Cookie。
	// 允许第一方 Cookie 将与 TikTok 共享来自你网站的第一方 Cookie 数据，可改善衡量效果并帮助触达更多用户。了解关于通过 TikTok Pixel 使用 Cookie。
	// 默认值： true。
	EnableFirstPartyCookies *bool `json:"enable_first_party_cookies,omitempty"`
	// EnableExpandedDataSharing 是否启用扩展数据共享。
	// 启用扩展数据共享将允许此 Pixel 收集来自你的网站的更多信息，包括页面内容（如元标记和商品详情）、访客操作（如点击、按钮互动和使用时长）和网站性能（如页面加载速度和响应能力）。此数据可用于提升 TikTok 的广告投放成效和推广系列成效。你可以通过 TikTok Pixel 助手查看收集的所有数据。了解通过 TikTok Pixel 实现扩展型数据共享功能。
	// 默认值： true。
	EnableExpandedDataSharing *bool `json:"enable_expanded_data_sharing,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
