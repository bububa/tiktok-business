package webhook

import "github.com/bububa/tiktok-business/model"

// CreativeFatigureEntity 广告疲劳状态
type CreativeFatigureEntity struct {
	// AdvID 广告主 ID
	AdvID model.Uint64 `json:"adv_id,omitempty"`
	// AdgroupIDs 通过advertiser_id订阅疲劳状态时返回。
	// 广告组ID列表。
	AdgroupIDs []model.Uint64 `json:"adgroup_ids,omitempty"`
	// AdIDs 通过advertiser_id或adgroup_id订阅疲劳状态时返回。
	// 广告ID列表。
	AdIDs []model.Uint64 `json:"ad_ids,omitempty"`
	// AdgroupID 通过ad_id或adgroup_id订阅疲劳状态时返回。
	// 广告组ID。
	AdgroupID model.Uint64 `json:"adgroup_id,omitempty"`
	// AdID 通过ad_id订阅疲劳状态时返回。
	// 广告ID。
	AdID model.Uint64 `json:"ad_id,omitempty"`
	// Date 检测到所订阅的对象（广告、广告组或广告账号）发生创意疲劳的日期。
	// 格式：YYYY-MM-DD.
	Date string `json:"date,omitempty"`
}
