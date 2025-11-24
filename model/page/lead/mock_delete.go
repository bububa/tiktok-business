package lead

import "github.com/bububa/tiktok-business/util"

// MockDeleteRequest 删除测试线索 API Request
type MockDeleteRequest struct {
	// AdvertiserID 当测试线索在广告账户中创建时，本字段必填。
	// 广告账户 ID。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// LibraryID 当测试线索在商务中心中的表单库中创建时，本字段必填。
	// 表单库 ID
	LibraryID string `json:"library_id,omitempty"`
	// PageID 想要删除的测试线索的 ID。
	// 若想获取当前的测试线索的 ID，可使用 /page/lead/mock/get/。
	PageID string `json:"page_id,omitempty"`
}

// Encode implements PostRequest
func (r *MockDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
