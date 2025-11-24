package tool

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InterestKeywordGetRequest 通过 ID 获取其他兴趣分类 API Request
type InterestKeywordGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// KeywordQuery 想要查询的其他兴趣分类信息。
	// 最大数量：50。
	KeywordQuery *Keyword `json:"keyword_query,omitempty"`
	// Filtering 筛选条件
	Filtering *InterestKeywordGetFilter `json:"filtering,omitempty"`
}

// InterestKeywordGetFilter 筛选条件
type InterestKeywordGetFilter struct {
	// AudienceType 想要筛选的受众类型。枚举值：
	// GENERAL_INTEREST: 大众兴趣。受众对某类别视频拥有长期兴趣。
	// PURCHASE_INTENTION: 购买意向。受众近期行为表明他们可能会购买某内容类别相关产品。
	// 默认值：GENERAL_INTEREST。
	AudienceType string `json:"audience_type,omitempty"`
}

// Encode implements GetRequest interface
func (r *InterestKeywordGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("keyword_query", string(util.JSONMarshal(r.KeywordQuery)))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// InterestKeywordGetResponse 通过 ID 获取其他兴趣分类 API Response
type InterestKeywordGetResponse struct {
	model.BaseResponse
	Data struct {
		Keywords []Keyword `json:"keywords,omitempty"`
	} `json:"data"`
}
