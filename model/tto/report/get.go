package report

import (
	"net/http"
	"strconv"

	"github.com/bububa/tiktok-business/model"
	tcmreport "github.com/bububa/tiktok-business/model/tcm/report"
	"github.com/bububa/tiktok-business/util"
)

type Metrics = tcmreport.Metrics
type VideoInfo = tcmreport.VideoInfo
type CreatorInfo = tcmreport.CreatorInfo
type GenderDistribution = tcmreport.GenderDistribution
type AgeDistribution = tcmreport.AgeDistribution
type CountryDistribution = tcmreport.CountryDistribution
type DeviceDistribution = tcmreport.DeviceDistribution
type LocaleDistribution = tcmreport.LocaleDistribution
type LanguageDistribution = tcmreport.LanguageDistribution
type AudienceInterestDistribution = tcmreport.AudienceInterestDistribution
type VideoViewsBySource = tcmreport.VideoViewsBySource
type DailyStat = tcmreport.DailyStat
type AnchorMetrics = tcmreport.AnchorMetrics

type GetRequest struct {
	AccountID   string `json:"tto_tcm_account_id,omitempty"`
	CampaignID  string `json:"campaign_id,omitempty"`
	CountryCode string `json:"-"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	Page        int    `json:"page,omitempty"`
	PageSize    int    `json:"page_size,omitempty"`
}

func (r *GetRequest) PreRequest(req *http.Request) error {
	if r.CountryCode != "" {
		req.Header.Set("Country-Code", r.CountryCode)
	}
	return nil
}

func (r *GetRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("tto_tcm_account_id", r.AccountID)
	v.Set("campaign_id", r.CampaignID)
	if r.StartDate != "" {
		v.Set("start_date", r.StartDate)
	}
	if r.EndDate != "" {
		v.Set("end_date", r.EndDate)
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
	AccountID           string              `json:"tto_tcm_account_id,omitempty"`
	CampaignID          string              `json:"campaign_id,omitempty"`
	CountryCode         string              `json:"country_code,omitempty"`
	CampaignType        string              `json:"campaign_type,omitempty"`
	CampaignChannelType string              `json:"campaign_channel_type,omitempty"`
	PageInfo            *model.PageInfo     `json:"page_info,omitempty"`
	Videos              []tcmreport.Metrics `json:"videos,omitempty"`
}
