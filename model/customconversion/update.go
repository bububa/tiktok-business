package customconversion

import "github.com/bububa/tiktok-business/util"

// UpdateRequest 更新自定义转化 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomConversionID 自定义转化的 ID
	CustomConversionID string `json:"custom_conversion_id,omitempty"`
	// Name 自定义转化的新名称。
	// 名称必须以字母开头，可以包含字母、数字、空格、下划线或横线。
	//
	// 长度限制: 50 个字符，不能包含表情符号。
	//
	// 注意:
	//
	// 不要指定标准事件的名称。
	// 不允许在广告账户内重复使用自定义转化名称。
	Name string `json:"name,omitempty"`
	// Description 自定义转化的新描述。
	// 长度限制: 500 个字符，不能包含表情符号。
	Description string `json:"description,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
