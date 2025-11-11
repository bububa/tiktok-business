package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// AdvertiserUnionpayInfoSubmitRequest 为营业执照提交银联验证 API Request
type AdvertiserUnionpayInfoSubmitRequest struct {
	// BcID 商务中心 ID。
	// 您需对商务中心有管理员权限。
	// 若想获取您有管理员权限的商务中心，可使用/bc/get/，需确保返回的 user_role 为 ADMIN
	BcID string `json:"bc_id,omitempty"`
	// AdvertiserID 广告账户 ID。
	// 您需指定与需要银联验证的营业执照绑定的广告账户 ID，且需对广告账户有管理员权限
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// RepresentativeName 营业执照上法人的中文名字
	RepresentativeName string `json:"representative_name,omitempty"`
	// RepresentativeDocumentType 法人证件类型。
	// 枚举值：
	// ID_CARD ：居民身份证。
	// PASSPORT：护照。
	// HK_MACAO_EXIT_ENTRY_PERMIT ：往来港澳通行证，又称港澳通行证。
	// TAIWAN_MAINLAND_TRAVEL_PERMIT ：台湾居民来往大陆通行证，又称台胞证。
	// HK_MACAO_MAINLAND_TRAVEL_PERMIT：港澳居民来往内地通行证，又称回乡证。
	// 默认值： ID_CARD
	RepresentativeDocumentType enum.RepresentativeDocumentType `json:"representative_document_type,omitempty"`
	// RepresentativeID 法人证件 ID。
	// representative_document_type 为 ID_CARD或未传入时，指定法人的身份证号，共 18 位。
	// 示例：234123190101012345。
	// representative_document_type 为 PASSPORT 时，指定包含字母和数字的护照号码。
	// 示例：EF1234567。
	// representative_document_type 为 HK_MACAO_EXIT_ENTRY_PERMIT 时，指定包含字母和数字的港澳通行证证件号码。
	// 示例：CA1234567。
	// representative_document_type 为 TAIWAN_MAINLAND_TRAVEL_PERMIT 时，指定仅包含数字的台胞证证件号码。
	// 示例： 12345678。
	// representative_document_type 为 HK_MACAO_MAINLAND_TRAVEL_PERMIT 时，指定包含字母和数字的回乡证证件号码。
	// 示例：H12345678。
	RepresentativeID string `json:"representative_id,omitempty"`
	// UnionpayAccount 银联账号，共 12 到 19 位
	UnionpayAccount string `json:"unionpay_account,omitempty"`
	// RepresentativePhoneNumber 法人的手机号码，共 11 位。
	RepresentativePhoneNumber string `json:"representative_phone_number,omitempty"`
}

// Encode implements PostRequest
func (r *AdvertiserUnionpayInfoSubmitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
