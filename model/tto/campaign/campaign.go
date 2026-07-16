package campaign

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

type Campaign struct {
	CampaignID                         string               `json:"campaign_id,omitempty"`
	CampaignType                       enum.TTOCampaignType `json:"campaign_type,omitempty"`
	BrandProfileID                     string               `json:"brand_profile_id,omitempty"`
	BrandName                          string               `json:"brand_name,omitempty"`
	CampaignName                       string               `json:"campaign_name,omitempty"`
	HandleNames                        []string             `json:"handle_names,omitempty"`
	CampaignDescription                string               `json:"campaign_description,omitempty"`
	AnchorID                           string               `json:"anchor_id,omitempty"`
	VideoIDs                           []string             `json:"video_ids,omitempty"`
	AdvertiserIDs                      []string             `json:"advertiser_ids,omitempty"`
	SparkAdsRequestedAuthorizationDays int                  `json:"spark_ads_requested_authorization_days,omitempty"`
	InviteLink                         string               `json:"invite_link,omitempty"`
	CreateTime                         model.DateTime       `json:"create_time,omitzero"`
	UpdateTime                         model.DateTime       `json:"update_time,omitzero"`
	CountryCodes                       []string             `json:"country_codes,omitempty"`
	FailedHandleNames                  []string             `json:"failed_handle_names,omitempty"`
	DMInfo                             *DMInfo              `json:"dm_info,omitempty"`
}
type DMInfo struct {
	BusinessAccountHandle string `json:"business_account_handle,omitempty"`
	DMAllowed             bool   `json:"dm_allowed,omitempty"`
}

type CreateRequest struct {
	AccountID                                 string               `json:"tto_tcm_account_id,omitempty"`
	CampaignType                              enum.TTOCampaignType `json:"campaign_type,omitempty"`
	BrandProfileID                            string               `json:"brand_profile_id,omitempty"`
	BrandName                                 string               `json:"brand_name,omitempty"`
	CampaignName                              string               `json:"campaign_name,omitempty"`
	CampaignDescription                       string               `json:"campaign_description,omitempty"`
	AnchorID                                  string               `json:"anchor_id,omitempty"`
	AdvertiserIDs                             []string             `json:"advertiser_ids,omitempty"`
	DefaultSparkAdsRequestedAuthorizationDays int                  `json:"default_spark_ads_requested_authorization_days,omitempty"`
	HandleNames                               []string             `json:"handle_names,omitempty"`
	SendNotification                          bool                 `json:"send_notification,omitempty"`
	CampaignID                                string               `json:"campaign_id,omitempty"`
	BusinessAccountHandle                     string               `json:"business_account_handle,omitempty"`
}

func (r *CreateRequest) Encode() []byte { return util.JSONMarshal(r) }

type UpdateRequest struct {
	AccountID        string               `json:"tto_tcm_account_id,omitempty"`
	CampaignID       string               `json:"campaign_id,omitempty"`
	CampaignType     enum.TTOCampaignType `json:"campaign_type,omitempty"`
	HandleNames      []string             `json:"handle_names,omitempty"`
	SendNotification bool                 `json:"send_notification,omitempty"`
	AdvertiserIDs    []string             `json:"advertiser_ids,omitempty"`
}

func (r *UpdateRequest) Encode() []byte { return util.JSONMarshal(r) }

type CampaignResponse struct {
	model.BaseResponse
	Data *Campaign `json:"data,omitempty"`
}

type GetRequest struct {
	AccountID    string               `json:"tto_tcm_account_id,omitempty"`
	CampaignIDs  []string             `json:"campaign_ids,omitempty"`
	CampaignType enum.TTOCampaignType `json:"campaign_type,omitempty"`
	Page         int                  `json:"page,omitempty"`
	PageSize     int                  `json:"page_size,omitempty"`
}

func (r *GetRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("tto_tcm_account_id", r.AccountID)
	if len(r.CampaignIDs) > 0 {
		v.Set("campaign_ids", string(util.JSONMarshal(r.CampaignIDs)))
	}
	if r.CampaignType != "" {
		v.Set("campaign_type", string(r.CampaignType))
	}
	if r.Page > 0 {
		v.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		v.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	Campaigns []Campaign      `json:"campaigns,omitempty"`
	PageInfo  *model.PageInfo `json:"page_info,omitempty"`
	DMInfo    *DMInfo         `json:"dm_info,omitempty"`
}
