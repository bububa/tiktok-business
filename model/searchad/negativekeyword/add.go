package negativekeyword

import (
	"github.com/bububa/tiktok-business/util"
)

// AddRequest 创建否定关键词 API Request
type AddRequest struct {
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
	// Replace 是否使用通过keywords字段指定的关键词替换掉推广系列或广告组目前对应的所有否定关键词。
	// 枚举值：
	// true：替换目前对应的所有否定关键词）。
	// false：仅增加否定关键词。
	// 默认值：false。
	Replace bool `json:"replace,omitempty"`
	// Keywords 您想创建的否定关键词列表。
	// 最大数量： 1,000。
	// 注意：单个广告组最多可设置 10,000 个否定关键词。
	Keywords []Keyword `json:"keywords,omitempty"`
}

// Encode implements PostRequest
func (r *AddRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
