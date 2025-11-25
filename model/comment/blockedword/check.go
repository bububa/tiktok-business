package blockedword

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CheckRequest 检查屏蔽词是否存在 API Request
type CheckRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// BlockedWords 屏蔽词列表
	BlockedWords []string `json:"blocked_words,omitempty"`
}

// Encode implements GetRequest
func (r *CheckRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("blocked_words", string(util.JSONMarshal(r.BlockedWords)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// CheckResponse 检查屏蔽词是否存在 API Response
type CheckResponse struct {
	model.BaseResponse
	Data struct {
		// Results 查询结果
		Results []CheckWord `json:"results,omitempty"`
	} `json:"data"`
}

// CheckWord 查询结果
type CheckWord struct {
	// Word 您希望检查的屏蔽词
	Word string `json:"word,omitempty"`
	// Blocked 是否已屏蔽。枚举值: True(已屏蔽)， False未屏蔽)
	Blocked bool `json:"blocked,omitempty"`
}
