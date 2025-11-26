package blockedword

import "github.com/bububa/tiktok-business/util"

// UpdateRequest 更新屏蔽词 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// OldWord 旧屏蔽词内容
	OldWord string `json:"old_word,omitempty"`
	// NewWord 新屏蔽词内容
	NewWord string `json:"new_word,omitempty"`
}

// Encode implements PostRequest
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
