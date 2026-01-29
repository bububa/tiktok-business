package smartplus

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest Get Upgraded Smart+ Ad Groups API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Fields Fields that you want to get.
	// When this field is not specified, all fields are returned by default.
	// For allowed fields, see the fields under list in the Response section
	Fields []string `json:"fields,omitempty"`
	// Filtering Filtering conditions
	Filtering *GetFilter `json:"filtering,omitempty"`
	// Page Current page number.
	// Default value: 1.
	// Value range: ≥ 1.
	Page int `json:"page,omitempty"`
	// PageSize Page size.
	// Default value: 10.
	// Value range: 1-1,000.
	PageSize int `json:"page_size,omitempty"`
}

type GetFilter struct {
	// CampaignIDs Filter by campaign IDs.
	// Max size: 100.
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// AdgroupIDs Filter by ad group IDs.
	// Max size: 100
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// AdgroupName Ad group name
	AdgroupName string `json:"adgroup_name,omitempty"`
	// PrimaryStatus Primary status.
	// For enum values, see Enumeration-Primary Status.
	// The default value is STATUS_NOT_DELETE, which returns ad groups in all statuses excluding STATUS_DELETE.
	// If you want to get ad groups in all statuses including STATUS_DELETE, use STATUS_ALL instead.
	PrimaryStatus enum.PrimaryStatus `json:"primary_status,omitempty"`
	// SecondaryStatus Ad group secondary status.
	// For enum values, see Enumeration - Ad Group Status - Secondary Status.
	SecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// ObjectiveType Advertising objective.
	// Currently, we support APP_PROMOTION, WEB_CONVERSIONS, and LEAD_GENERATION.
	// For detailed explanation of enum values, see Enumeration-Advertising Objective.
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// SalesDestination Sales destination, the destination where you want to drive your sales.
	// Enum values:
	// WEBSITE: Website. Drive sales on your website.
	// APP: App. Drive sales on your app (product catalog required).
	// WEB_AND_APP: Website and app. Drive sales on both your website and your app.
	SalesDestination enum.SalesDestination `json:"sales_destination,omitempty"`
	// PromotionType Promotion type (Optimization location). You can decide where you'd like to promote your products using this field.
	// Currently, we support APP_ANDROID, APP_IOS, WEBSITE, and LEAD_GENERATION.
	// When objective_type is APP_PROMOTION, set this field to APP_ANDROID or APP_IOS.
	// When objective_type is WEB_CONVERSIONS, set this field to WEBSITE.
	// When objective_type is LEAD_GENERATION, set this field to LEAD_GENERATION.
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// OptimizationGoal The measurable results you'd like to drive with your campaigns.
	// Currently, we support CLICK, INSTALL, IN_APP_EVENT, VALUE, CONVERT, TRAFFIC_LANDING_PAGE_VIEW, and LEAD_GENERATION.
	// To find the detailed description for each optimization goal, see Enumeration - Optimization Goal.
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
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

// GetResponse Get Upgraded Smart+ Ad Groups API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// List 符合条件的广告组列表，请求参数未设置fields字段或者fields中包含以下子字段时会返回相关字段
	List []Adgroup `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
