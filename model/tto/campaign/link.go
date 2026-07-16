package campaign

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

type LinkRequest struct {
	AccountID  string             `json:"tto_tcm_account_id,omitempty"`
	CampaignID string             `json:"campaign_id,omitempty"`
	VideoID    string             `json:"video_id,omitempty"`
	Action     enum.TTOLinkAction `json:"action,omitempty"`
}

func (r *LinkRequest) Encode() []byte { return util.JSONMarshal(r) }

type LinkInfo struct {
	VideoID               string               `json:"video_id,omitempty"`
	Status                enum.TTOLinkStatus   `json:"status,omitempty"`
	CampaignID            string               `json:"campaign_id,omitempty"`
	LastRemindedTimestamp model.Int64          `json:"last_reminded_timestamp,omitempty"`
	NumberOfRequests      model.Int64          `json:"number_of_requests,omitempty"`
	CampaignType          enum.TTOCampaignType `json:"campaign_type,omitempty"`
	LinkRequestID         string               `json:"link_request_id,omitempty"`
}
type LinkResponse struct {
	model.BaseResponse
	Data *LinkInfo `json:"data,omitempty"`
}

type LinkStatusGetRequest struct {
	AccountID    string               `json:"tto_tcm_account_id,omitempty"`
	CampaignIDs  []string             `json:"campaign_ids,omitempty"`
	HandleNames  []string             `json:"handle_names,omitempty"`
	CampaignType enum.TTOCampaignType `json:"campaign_type,omitempty"`
	Page         int                  `json:"page,omitempty"`
	PageSize     int                  `json:"page_size,omitempty"`
}

func (r *LinkStatusGetRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("tto_tcm_account_id", r.AccountID)
	if len(r.CampaignIDs) > 0 {
		v.Set("campaign_ids", string(util.JSONMarshal(r.CampaignIDs)))
	}
	if len(r.HandleNames) > 0 {
		v.Set("handle_names", string(util.JSONMarshal(r.HandleNames)))
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

type LinkStatusGetResponse struct {
	model.BaseResponse
	Data *LinkStatusGetResult `json:"data,omitempty"`
}

type LinkStatusGetResult struct {
	VideoInfos []LinkInfo      `json:"video_infos,omitempty"`
	PageInfo   *model.PageInfo `json:"page_info,omitempty"`
}
