package dmp

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SegmentAudienceRequest 创建/删除受众细分群 API Request
type SegmentAudienceRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Action 需要执行的行动。枚举值：
	// create: 创建受众细分群，需在请求中同时传入custom_audience_name 。
	// delete: 删除受众细分群，需请求中同时传入delete_audience_id。
	// 不传时默认创建受众细分群，默认值为create。
	Action string `json:"action,omitempty"`
	// CustomAudienceName 新建受众细分群的名称。请注意，如果名称超过70个字符，名称将会被缩减到67个字符并以“...”结束。当action值为 create 时必填
	CustomAudienceName string `json:"custom_audience_name,omitempty"`
	// DeleteAudienceID 需要删除的受众ID。当action值为 delete 时必填
	DeleteAudienceID string `json:"delete_audience_id,omitempty"`
}

// Encode implements PostRequest
func (r *SegmentAudienceRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// SegmentAudienceResponse 创建/删除受众细分群 API Response
type SegmentAudienceResponse struct {
	model.BaseResponse
	Data struct {
		// AudienceID 受众ID。您在受众目标映射的时候需要传入受众ID的值
		AudienceID string `json:"audience_id,omitempty"`
	}
}
