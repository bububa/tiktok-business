package advertiser

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Advertiser 广告账号（广告主账号）信息
type Advertiser struct {
	// AdvertiserID 广告账号 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// OwnerBcID 广告账户所属的商务中心的 ID。
	// 注意：仅当以下条件均满足时返回本字段：
	//
	// 对于拥有该广告账户的商务中心，用户有管理员权限，
	// 用户有对该广告账户的管理员（Admin）/操作员（Operator）/ 分析师（Analyst））权限，且该权限从拥有该广告账户的商务中心，而非合作伙伴商务中心处获得。
	OwnerBcID string `json:"owner_bc_id,omitempty"`
	// Status 广告账号状态。
	// 枚举值详见枚举值-广告主状态。
	// 示例：STATUS_ENABLE。
	Status enum.AdvertiserStatus `json:"status,omitempty"`
	// Role 广告账号角色。
	// 枚举值详见枚举值-广告主角色。
	// 示例：ROLE_ADVERTISER。
	Role enum.Role `json:"role,omitempty"`
	// RejectionReason 审核拒绝原因
	RejectionReason string `json:"rejection_reason,omitempty"`
	// Name 广告账号名称
	Name string `json:"name,omitempty"`
	// Timezone 广告账号与 GMT 偏移的时区信息，或广告账号所在时区在时区数据库中的名字，格式为“地域/城市”。
	// 示例：Etc/GMT，Europe/London
	Timezone *model.TimeLocation `json:"timezone,omitempty"`
	// DisplayTimezone 广告账号所在时区在时区数据库中的名字。
	// 示例：Europe/London。
	DisplayTimezone *model.TimeLocation `json:"display_timezone,omitempty"`
	// Company 广告账号公司名
	Company string `json:"company,omitempty"`
	// CompanyNameEditable 该广告账号的公司名称是否支持通过 API 修改。
	// 若本字段值为 true，则支持通过/advertiser/update/中的 company 字段修改公司名称。
	// 对于非自助广告主（即由 TikTok 代表或直客销售直接管理的广告主），本字段值将为 false。此时若想修改公司名称，需联系 TikTok 代表。
	CompanyNameEditable *bool `json:"company_name_editable,omitempty"`
	// Industry 广告账号行业类别代码。
	// 枚举值详见行业信息。
	// 示例：290303。
	Industry string `json:"industry,omitempty"`
	// Address 广告账号地址信息
	Address string `json:"address,omitempty"`
	// Country 广告账号注册地代码
	Country string `json:"country,omitempty"`
	// AdvertiserAccountType 广告账户类型。
	// 枚举值：RESERVATION （合约广告账户），AUCTION（竞价广告账户）。
	AdvertiserAccountType enum.ServiceType `json:"advertiser_account_type,omitempty"`
	// Currency 广告账号使用的货币类型，格式为 ISO 4217 代码。
	// 示例：EUR。
	Currency string `json:"currency,omitempty"`
	// Contacter 联系人姓名。详细信息已脱敏。
	// 示例：Te********************nt。
	Contacter string `json:"contacter,omitempty"`
	// Email 广告账号联系人邮箱。详细信息已脱敏。
	// 示例：l***********@*************。
	Email string `json:"email,omitempty"`
	// CellphoneNumber 联系手机号码。详细信息已脱敏
	CellphoneNumber string `json:"cellphone_number,omitempty"`
	// TelephoneNumber 固定电话号码。详细信息已脱敏
	TelephoneNumber string `json:"telephone_number,omitempty"`
	// Language 广告账号所使用的语言的代码
	Langauge string `json:"language,omitempty"`
	// LicenseNo 营业执照编号
	LicenseNo string `json:"license_no,omitempty"`
	// LicenseURL 营业执照预览 URL。
	// 链接默认 1 小时内有效。
	LicenseURL string `json:"license_url,omitempty"`
	// Description 品牌描t述
	Description string `json:"description,omitempty"`
	// Balance 广告账户可用余额。
	// 单位与广告账户货币类型（currency）相关。
	// 示例：5325.43。
	Balance model.Float64 `json:"balance,omitempty"`
	// CreateTime 广告账号创建时间，格式为 Epoch/Unix 时间戳，单位为秒。
	// 示例：1510740064。
	CreateTime int64 `json:"create_time,omitempty"`
}
