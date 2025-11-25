package blockedword

import "github.com/bububa/tiktok-business/util"

// CreateRequest 创建屏蔽词 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// BlockedWords 屏蔽词列表
	BlockedWords []string `json:"blocked_words,omitempty"`
}

// Encode implements PostRequest
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
