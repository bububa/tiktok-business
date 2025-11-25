package dmp

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CustomAudienceLookalikeCreateRequest 创建相似受众 API Request
type CustomAudienceLookalikeCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceName 受众群体名称。限128字符内
	CustomAudienceName string `json:"custom_audience_name,omitempty"`
	// AudienceSubType 受众子类型，表明可以使用的广告类型。枚举值：NORMAL: 常规受众，可用于非覆盖和频次广告。REACH_FREQUENCY: 覆盖和频次广告受众，只可用于覆盖和频次广告。默认： NORMAL。
	AudienceSubType enum.AudienceSubType `json:"audience_sub_type,omitempty"`
	// LookalikeSpec 相似受众群体的参数
	LookalikeSpec *LookalikeSpec `json:"lookalike_spec,omitempty"`
}

// Encode implements PostRequest
func (r *CustomAudienceLookalikeCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CustomAudienceLookalikeCreateResponse 创建相似受众 API Response
type CustomAudienceLookalikeCreateResponse struct {
	model.BaseResponse
	Data struct {
		// CustomAudienceID 自定义受众 ID
		CustomAudienceID string `json:"custom_audience_id,omitempty"`
	} `json:"data"`
}
