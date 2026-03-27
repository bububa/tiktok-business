package webhook

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// TCMVideosEntity the linking of a TikTok video to a TTO campaign
type TCMVideosEntity struct {
	// EventType Webhook Event type.
	// Supported value:
	// tto_campaign_webhook: TTO video linking event.
	EventType string `json:"event_type,omitempty"`
	// VideoID ID of the TikTok video that is linked to a TTO campaign.
	VideoID string `json:"video_id,omitempty"`
	// EventTime The time when the event occurred, in ISO 8601 format.
	// Example: 2024-11-18T21:57:00.070272Z.
	EventTime model.DateTime `json:"event_time,omitzero"`
	// HandleName Handle name of the creator.
	// Note: If the handle name of the creator cannot be matched, this field will not be returned.
	HandleName string `json:"handle_name,omitempty"`
	// TTOCampaignEventType The specific type of TTO video linking event.
	// Supported values:
	// tto_video_uploaded: A TikTok video has been uploaded and linked to a TTO campaign, but has not passed review.
	// tto_video_published: A TikTok video has been uploaded and linked to a TTO campaign, and has passed review.
	TTOCampaignEventType string `json:"tto_campaign_event_type,omitempty"`
	// CampaignID ID of the TTO campaign
	CampaignID string `json:"campaign_id,omitempty"`
	// TTOTCMAccountID ID of a TTO Creator Marketplace account.
	TTOTCMAccountID string `json:"tto_tcm_account_id,omitempty"`
	// LogID Log ID.
	LogID string `json:"log_id,omitempty"`
}

func (e TCMVideosEntity) Entity() enum.SubscribeEntity {
	return enum.SubscribeEntity_TCM_VIDEOS
}
