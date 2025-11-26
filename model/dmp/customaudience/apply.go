package customaudience

import "github.com/bububa/tiktok-business/util"

// ApplyRequest 将受众应用到多个广告组 API Request
type ApplyRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceID 自定义受众ID。您只能传入一个自定义受众ID
	CustomAudienceID string `json:"custom_audience_id,omitempty"`
	// AdgroupIDs 广告组ID列表。
	// 注意：
	// adgroup_ids 和 custom_audience_id 必须属于同一个广告主账号。否则将会报错。
	// 不能将相似受众用于覆盖和频次广告，否则会出错。详情见下。
	// 1. 如果custom_audience_id为相似受众，且adgroup_ids为覆盖和频次广告组，系统会报错。
	// 2. 如果custom_audience_id为相似受众，且受众子类型为REACH_FREQUENCY，系统会报错
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// ActionMode 要对受众进行的特定操作。枚举值：Apply（应用），Disconnect（解除应用）
	ActionMode string `json:"action_mode,omitempty"`
	// UsageMode 使用方式。选择将该受众在广告组中作为定向使用或作为排除使用。当action_mode为Apply时必填。枚举值： Include（作为定向使用）， Exclude（作为排除使用）
	UsageMode string `json:"usage_mode,omitempty"`
}

// Encode implements PostRequest
func (r *ApplyRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
