package negativekeyword

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 更新否定关键词 API Request
type UpdateRequest struct {
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
	// OldKeywordID 您想更新的否定关键词的ID。
	// 若想获取否定关键词的ID，请使用/search_ad/negative_word/get/。
	// 注意: 若指定的否定关键词ID不存在，会出现报错提示。
	OldKeywordID string `json:"old_keyword_id,omitempty"`
	// Keyword 原否定关键词的更新信息
	Keyword *Keyword `json:"keyword,omitempty"`
}

// Encode implements PostRequest
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新否定关键词 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data struct {
		// NewKeywordID 否定关键词更新后对应的ID
		NewKeywordID string `json:"new_keyword_id,omitempty"`
	} `json:"data"`
}
