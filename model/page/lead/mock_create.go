package lead

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// MockCreateRequest 创建测试线索 API Request
type MockCreateRequest struct {
	// LeadSource 线索来源。
	// 枚举值：
	// INSTANT_FORM：通过即时表单生成的线索。
	// DIRECT_MESSAGE：通过绑定的企业号的私信生成的线索。
	// 默认值：INSTANT_FORM。
	LeadSource enum.LeadSource `json:"lead_source,omitempty"`
	// AdvertiserID 当您想在广告账户中创建测试线索时，本字段必填。
	// 广告账户 ID。
	// 注意：若您传入本字段，您需要有该广告账户的管理员权限。若想获取您有管理员权限的广告账户列表，可使用 /bc/asset/get/。将 asset_type 设置为 ADVERTISER，并挑选返回的 advertiser_role 为 ADMIN 的广告账户。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// LibraryID 当您想在商务中心中的表单库中创建测试线索时，本字段必填。
	// 表单库 ID。
	// 若想获取某一商务中心中的表单库列表，可使用 /page/library/get/
	LibraryID string `json:"library_id,omitempty"`
	// PageID 当 lead_source 为 INSTANT_FORM 或未传入时必填。
	// 当 lead_source 为 DIRECT_MESSAGE 时不支持本字段。
	// 即时表单的页面 ID。
	// 若想获取即时表单列表，可使用 /page/get/ 并将 business_type 设置为 LEAD_GEN。
	PageID string `json:"page_id,omitempty"`
}

// Encode implements PostRequest
func (r *MockCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// MockCreateResponse 创建测试线索 API Response
type MockCreateResponse struct {
	model.BaseResponse
	Data *Lead `json:"data,omitempty"`
}
