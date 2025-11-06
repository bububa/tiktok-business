package creator

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SimilarRequest 获取相似创作者 API Request
type SimilarRequest struct {
	// TCMAccountID TikTok Creator Marketplace 账户 ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// HandleName 您想要用于搜索相似创作者的的 TTCM 创作者用户名
	HandleName string `json:"handle_name,omitempty"`
}

// Encode implements GetRequest
func (r *SimilarRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("tcm_account_id", r.TCMAccountID)
	values.Set("handle_name", r.HandleName)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// SimilarResponse 获取相似创作者 API Response
type SimilarResponse struct {
	model.BaseResponse
	Data struct {
		// Creators 返回的相似创作者
		Creators []Creator `json:"creators,omitempty"`
	} `json:"data"`
}
