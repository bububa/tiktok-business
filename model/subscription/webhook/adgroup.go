package webhook

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// AdgroupEntity 广告审核状态
type AdgroupEntity struct {
	// LogID 日志 ID
	LogID string `json:"log_id,omitempty"`
	// AdvID 广告主 ID
	AdvID model.Uint64 `json:"adv_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName 推广系列名称
	CampaignName string `json:"campaign_name,omitempty"`
	// AdgroupID 广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName 广告组名称
	AdgroupName string `json:"adgroup_name,omitempty"`
	// RejectReason 拒审原因
	RejectReason string `json:"reject_reason,omitempty"`
}

func (e AdgroupEntity) Entity() enum.SubscribeEntity {
	return enum.SubscribeEntity_AD_GROUP
}

type AdEntity struct {
	AdgroupEntity
	// AdID 广告 ID
	AdID string `json:"ad_id,omitempty"`
	// AdName 广告名称
	AdName string `json:"ad_name,omitempty"`
}

func (e AdEntity) Entity() enum.SubscribeEntity {
	return enum.SubscribeEntity_AD
}
