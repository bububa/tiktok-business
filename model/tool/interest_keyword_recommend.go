package tool

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InterestKeywordRecommendRequest 搜索其他兴趣分类 API Request
type InterestKeywordRecommendRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Keywords keywords和keyword必填其一。
	// 若同时传入keywords 和keyword，将忽略 keyword。
	// 种子关键词列表，用来搜索其他兴趣分类。
	// 列表中的每个种子关键词都将单独用于生成对应的其他兴趣分类。
	// 最大数量：10。
	Keywords []string `json:"keywords,omitempty"`
	// Keyword keywords和keyword必填其一。
	// 若同时传入keywords 和keyword，将忽略 keyword。
	// 单个种子关键词，用来搜索其他兴趣分类。
	Keyword string `json:"keyword,omitempty"`
	// Mode 搜索模式。
	// 枚举值: FUZZ_MATCH(模糊匹配)，SEMANTIC_RECOMMEND(相似搜索)。
	// 默认值: FUZZ_MATCH。
	Mode string `json:"mode,omitempty"`
	// Language 关键词语言。
	// 枚举值：fr, id, it, ja, ms, ar, vi, en, ru, es, th, tr, hi, zh, de, ko。
	// 默认: en (英语)。
	// 需要注意的是，定向受众语言不受关键词语言限制。例如，定向非英语地区受众可以使用英语分类。
	Language string `json:"language,omitempty"`
	// Limit 你想获得的推荐的其他兴趣分类数。
	// 默认值: 50。
	// 取值范围: 1-50。
	Limit int `json:"limit,omitempty"`
	// AudienceType 受众类型。
	// 枚举值:
	// GENERAL_INTEREST:  大众兴趣。受众对某类别视频拥有长期兴趣。
	// PURCHASE_INTENTION（已废弃）: 购买意向。受众近期行为表明他们可能会购买某内容类别相关产品。
	// 默认值:GENERAL_INTEREST。
	AudienceType string `json:"audience_type,omitempty"`
}

// Encode implements GetRequest interface
func (r *InterestKeywordRecommendRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.Keywords) > 0 {
		values.Set("keywords", string(util.JSONMarshal(r.Keywords)))
	}
	if r.Keyword != "" {
		values.Set("keyword", r.Keyword)
	}
	if r.Mode != "" {
		values.Set("mode", r.Mode)
	}
	if r.Language != "" {
		values.Set("language", r.Language)
	}
	if r.Limit > 0 {
		values.Set("limit", strconv.Itoa(r.Limit))
	}
	if r.AudienceType != "" {
		values.Set("audience_type", r.AudienceType)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// InterestKeywordRecommendResponse 搜索其他兴趣分类 API Response
type InterestKeywordRecommendResponse struct {
	model.BaseResponse
	Data struct {
		RecommendedKeywords []RecommendedKeyword `json:"recommended_keywords,omitempty"`
	} `json:"data"`
}

type RecommendedKeyword struct {
	// InputKeyword 仅当请求中传入 keywords 且 keywords 中包含多个关键词时返回。
	// 其他兴趣分类对应的种子关键词。
	// 例如，若种子关键词为"test"，对应的推荐兴趣关键词为"Tester"，则返回的input_keyword 将为"test"，keyword 将为 "Tester"
	InputKeyword string `json:"input_keyword,omitempty"`
	Keyword
}

type Keyword struct {
	// Keyword 其他兴趣分类
	Keyword string `json:"keyword,omitempty"`
	// KeywordID 其他兴趣分类 ID
	KeywordID string `json:"keyword_id,omitempty"`
	// Language 其他兴趣分类语言。
	// 枚举值：fr, id, it, ja, ms, ar, vi, en, ru, es, th, tr, hi, zh, de, ko
	Language string `json:"language,omitempty"`
	// Status 其他兴趣分类状态，即分类是否有效。
	// 枚举值:
	// EFFECTIVE：有效。有效其他兴趣分类指一个分类能覆盖的受众人数足够大。
	// INEFFECTIVE：无效。无效其他兴趣分类指一个分类能覆盖的受众人数太少。
	// 注意: 仅有效状态的其他兴趣分类可在广告组中用于定向受众。
	Status enum.KeywordStatus `json:"status,omitempty"`
	// KeywordStatus 标签状态。
	// 枚举值: ONLINE (可用), OFFLINE (不可用)。
	KeywordStatus enum.OnlineOffline `json:"keyword_status,omitempty"`
}
