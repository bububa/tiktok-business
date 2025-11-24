package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// HashtagRecommendRequest 搜索定向话题标签 API Request
type HashtagRecommendRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Keywords 关键词，用于生成一组相关的话题标签。
	// 如果您传入了多个不相关的关键词且将 operator 设置为 AND，可能不会返回有效推荐。
	// 最大数量：10。
	Keywords []string `json:"keywords,omitempty"`
	// Operator 关键词的运算符。
	// 枚举值：
	// AND：与。为 keywords 中指定的所有关键词整体生成推荐话题标签列表。
	// OR：或。为 keywords 中指定的每个关键词分别生成推荐话题标签列表。
	// 默认值：AND。
	Operator enum.FilterSetOperator `json:"operator,omitempty"`
}

// Encode implements GetRequest interface
func (r *HashtagRecommendRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("keywords", string(util.JSONMarshal(r.Keywords)))
	if r.Operator != "" {
		values.Set("operator", string(r.Operator))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// HashtagRecommendResponse 搜索定向话题标签 API Response
type HashtagRecommendResponse struct {
	model.BaseResponse
	Data struct {
		// RecommendedKeywords 推荐话题标签列表
		RecommendedKeywords []RecommendedKeyword `json:"recommended_keywords,omitempty"`
	} `json:"data"`
}
