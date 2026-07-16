package creator

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

type CampaignJoinRequest struct {
	CreatorID  string `json:"creator_id,omitempty"`
	InviteLink string `json:"tto_invite_link,omitempty"`
}

func (r *CampaignJoinRequest) Encode() []byte { return util.JSONMarshal(r) }

type CampaignVideoLinkRequest struct {
	CreatorID  string `json:"creator_id,omitempty"`
	InviteLink string `json:"tto_invite_link,omitempty"`
	VideoID    string `json:"video_id,omitempty"`
}

func (r *CampaignVideoLinkRequest) Encode() []byte { return util.JSONMarshal(r) }

type LinkRequestGetRequest struct {
	CreatorID  string `json:"creator_id,omitempty"`
	InviteLink string `json:"tto_invite_link,omitempty"`
}

func (r *LinkRequestGetRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("creator_id", r.CreatorID)
	v.Set("tto_invite_link", r.InviteLink)
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type LinkRequest struct {
	VideoID       string `json:"video_id,omitempty"`
	LinkRequestID string `json:"link_request_id,omitempty"`
}
type LinkRequestGetResponse struct {
	model.BaseResponse
	Data struct {
		CampaignID             string        `json:"campaign_id,omitempty"`
		SparkAuthorizationDays int           `json:"spark_authorization_days,omitempty"`
		LinkRequests           []LinkRequest `json:"link_requests,omitempty"`
	} `json:"data"`
}

type LinkRequestConfirmRequest struct {
	CreatorID     string                    `json:"creator_id,omitempty"`
	LinkRequestID string                    `json:"link_request_id,omitempty"`
	Action        enum.TTOLinkConfirmAction `json:"action,omitempty"`
}

func (r *LinkRequestConfirmRequest) Encode() []byte { return util.JSONMarshal(r) }

type EmptyResponse struct {
	model.BaseResponse
	Data struct{} `json:"data"`
}
