package tool

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// HashtagGetRequest 通过 ID 获取定向话题标签 API Request
type HashtagGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// KeywordIDs 话题标签ID列表
	KeywordIDs []string `json:"keyword_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r *HashtagGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("keyword_ids", string(util.JSONMarshal(r.KeywordIDs)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// HashtagGetResponse 通过 ID 获取定向话题标签 API Response
type HashtagGetResponse struct {
	model.BaseResponse
	Data struct {
		// KeywordsStatus 视频话题标签列表
		KeywordsStatus []Keyword `json:"keywords_status,omitempty"`
	} `json:"data"`
}
