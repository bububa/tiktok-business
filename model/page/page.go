package page

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Page 页面
type Page struct {
	// PageID 页面 ID
	PageID string `json:"page_id,omitempty"`
	// Status 页面状态。
	// 枚举值: EDITED（草稿）,PUBLISHED（可投放）
	Status enum.PageStatus `json:"status,omitempty"`
	// CreateTime 创建时间，格式为YYYY-MM-DD HH:MM:SS（UTC 时间）。
	// 示例：2024-01-01 00:00:00。
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// UpdateTime 更新时间，格式为YYYY-MM-DD HH:MM:SS（UTC 时间）。
	// 示例：2024-01-01 00:00:00。
	UpdateTime model.DateTime `json:"update_time,omitzero"`
	// PublishTime 发布时间，格式为YYYY-MM-DD HH:MM:SS（UTC 时间）。
	// 示例：2024-01-01 00:00:00。
	PublishTime model.DateTime `json:"publish_time,omitzero"`
	// Title 标题
	Title string `json:"title,omitempty"`
	// PreviewURL 预览链接（可投放时返回，草稿状态时为空）
	PreviewURL string `json:"preview_url,omitempty"`
	// Thumbnail 页面预览图片
	Thumbnail string `json:"thumbnail,omitempty"`
	// AppID 仅为所有应用介绍页和指定了应用的自定义 TikTok 即时体验页面返回本字段。
	// 页面中指定的应用的 ID。
	// 若想筛选应用介绍页，可在请求中将 business_type 设置为 APP_PROFILE_PAGE 或将 business_types 设置为 ["APP_PROFILE_PAGE"]。
	AppID string `json:"app_id,omitempty"`
	// HasCpp 仅为自定义 TikTok 即时体验页面返回本字段。
	// 自定义 TikTok 即时体验页面中是否包含自定产品页面。
	// 枚举值：true，false。
	// 若想筛选自定义 TikTok 即时体验页面，可在请求中将 business_type 设置为 TIKTOK_INSTANT_PAGE 或将 business_types 设置为["TIKTOK_INSTANT_PAGE"]。
	HasCpp bool `json:"has_cpp,omitempty"`
	// DestinationURLs 当business_type 为LEAD_GEN（线索表单），APP_PROFILE_PAGE（应用介绍页）或TIKTOK_INSTANT_PAGE ( 自定义页面 (自定义 TikTok 即时体验页面 ) ) 时，返回此字段。
	// 目标页 URL。
	DestinationURLs []string `json:"destination_urls,omitempty"`
	// MessageAppType 与自定义 TikTok 即时体验页面绑定的即时通讯应用类型。
	// 枚举值：
	// MESSENGER：Messenger。
	// WHATSAPP：WhatsApp。
	// IM_URL：自定义 TikTok 即时体验页面没有绑定即时通讯应用。
	MessageAppType enum.MessageAppType `json:"message_app_type,omitempty"`
	// MessageAppAccountID 与自定义 TikTok 即时体验页面绑定的即时通讯应用账号 ID。
	// messaging_app_type 为 MESSENGER 时，本字段代表 Facebook 公共主页编号。
	// messaging_app_type 为 WHATSAPP 时，本字段代表 WhatsApp 电话号码。
	// messaging_app_type 为 IM_URL 时，本字段将为空字符（""）。
	MessageAppAccountID string `json:"message_app_account_id,omitempty"`
	// TransferStatus 页面是否转移到商务中心。
	// 枚举值: UNSET,TRANSFERRED。
	TransferStatus string `json:"transfer_status,omitempty"`
	// TemplateID 页面模板 ID
	TemplateID string `json:"template_id,omitempty"`
	// UserID 创建页面的用户 ID
	UserID string `json:"user_id,omitempty"`
	// IsAssociated 商品橱窗页是否已被关联至广告（商品橱窗页不能重复使用，每个页面只能关联一个广告）
	IsAssociated bool `json:"is_associated,omitempty"`
	// DuplicateID 复制源页面 ID
	DuplicateID string `json:"duplicate_id,omitempty"`
}
