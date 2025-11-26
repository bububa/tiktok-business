package blockedword

import "github.com/bububa/tiktok-business/util"

// DeleteRequest 删除屏蔽词 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// BlockedWords 屏蔽词列表
	BlockedWords []string `json:"blocked_words,omitempty"`
}

// Encode implements PostRequest
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
