package smartplus

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest Get Upgraded Smart+ Campaigns API Request
type GetRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Fields Fields that you want to get.
	// When not specified, all fields are returned by default.
	// For allowed fields, see the fields under list in the Response section.
	Fields []string `json:"fields,omitempty"`
	// Page Current page number.
	// Default value: 1.
	// Value range: ≥ 1.
	Page int `json:"page,omitempty"`
	// PageSize Page size.
	// Default value: 10.
	// Value range: 1-1,000.
	PageSize int `json:"page_size,omitempty"`
	// Filtering Filtering conditions.
	Filtering *GetFilter `json:"filtering,omitempty"`
}

type GetFilter struct {
	// CampaignIDs Filter by campaign IDs.
	// Max size: 100.
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// CampaignName Filter by campaign name.
	// Fuzzy match is supported.
	CampaignName string `json:"campaign_name,omitempty"`
	// PrimaryStatus Primary status.
	// For enum values, see Enumeration-Primary Status.
	PrimaryStatus enum.PrimaryStatus `json:"primary_status,omitempty"`
	// SecondaryStatus Filter by campaign secondary status.
	// For enum values, see Enumeration- Campaign Status - Secondary Status.
	SecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// ObjectiveType Filter by advertising objective.
	// For enum values, see Enumeration-Advertising Objective.
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// SalesDestination Sales destination, the destination where you want to drive your sales.
	// Enum values:
	// WEBSITE: Website. Drive sales on your website.
	// APP: App. Drive sales on your app (product catalog required).
	// WEB_AND_APP: Website and app. Drive sales on both your website and your app.
	SalesDestination enum.SalesDestination `json:"sales_destination,omitempty"`
	// CreationFilterStartTime The earliest campaign creation time, in the format of YYYY-MM-DD HH:MM:SS (UTC time). This filter will retrieve campaigns created after this time.
	// Example: 2025-01-01 00:00:01.
	// If creation_filter_start_time and creation_filter_end_time are not specified, the results will include campaigns created at any time.
	// Note: To ensure task efficiency and speed, consider setting a time range within 6 months for the creation time filters.
	CreationFilterStartTime model.DateTime `json:"creation_filter_start_time,omitzero"`
	// CreationFilterEndTime The latest campaign creation time, in the format of YYYY-MM-DD HH:MM:SS (UTC time). This filter will retrieve campaigns created before this time.
	// Example: 2025-02-01 00:00:01.
	// If creation_filter_start_time and creation_filter_end_time are not specified, the results will include campaigns created at any time.
	// Note: To ensure task efficiency and speed, consider setting a time range within 6 months for the creation time filters.
	CreationFilterEndTime model.DateTime `json:"creation_filter_end_time,omitzero"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse Get Upgraded Smart+ Campaigns API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// List 推广系列列表
	List []Campaign `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
