package discovery

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SearchRequest Get trending search keywords API Request
type SearchRequest struct {
	// BusinessID Application specific unique identifier for the TikTok account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// IsPersonalized Whether to return personalized search keywords for the Business Account.
	// Supported values:
	// true: To return trending search keywords that are personalized for the Business Account.
	// false: To return trending search keywords that are not personalized for the Business Account.
	// Default value: false.
	IsPersonalized bool `json:"is_personalized,omitempty"`
}

// Encode implements GetRequest
func (r *SearchRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	if r.IsPersonalized {
		values.Set("is_personalized", "true")
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// SearchResponse Get trending search keywords API Response
type SearchResponse struct {
	model.BaseResponse
	Data struct {
		// SearchKeywords A list of 20 trending search keywords.
		SearchKeywords []string `json:"search_keywords,omitempty"`
	} `json:"data"`
}
