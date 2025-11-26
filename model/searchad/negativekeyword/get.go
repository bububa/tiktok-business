package negativekeyword

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取否定关键词 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// 	ObjectType 对象类型。请通过本字段指定您想获取推广系列还是广告组对应的否定关键词列表。
	// 枚举值：
	// CAMPAIGN：推广系列。
	// ADGROUP：广告组。
	ObjectType string `json:"object_type,omitempty"`
	// ObjectID 对象ID。请通过本字段传入推广系列ID或广告组ID，从而获取对应的否定关键词列表。
	// 您需根据 object_type字段传入本字段：
	// 若 object_type设置为CAMPAIGN，您需通过本字段传入推广系列ID。
	// 若 object_type 设置为 ADGROUP , 您需通过本字段传入广告组ID。
	ObjectID string `json:"object_id,omitempty"`
	// Page 请通过本字段指定结果显示的页数。
	// 默认值： 1。
	// 取值范围： ≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值： 10。
	// 取值范围： [1,50]。
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("object_type", r.ObjectType)
	values.Set("object_id", r.ObjectID)
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 获取否定关键词 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Keywords 否定关键词相关信息。
	// 最大数量：1,000。
	Keywords []Keyword `json:"keywords,omitempty"`
}

// Keyword 否定关键词
type Keyword struct {
	// KeywordID 否定关键词的ID
	KeywordID string `json:"keyword_id,omitempty"`
	// Name 否定关键词的名称
	Name string `json:"name,omitempty"`
	// MatchType 否定关键词的匹配类型。
	// 枚举值：
	// PRECISE_WORD ：精确匹配。
	// PHRASE_WORD：词组匹配。
	// BROAD_WORD：广泛匹配。
	MatchType enum.KeywordMatchType `json:"match_type,omitempty"`
}
