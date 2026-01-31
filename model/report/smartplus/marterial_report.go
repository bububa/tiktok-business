package smartplus

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/report"
	"github.com/bububa/tiktok-business/util"
)

// MaterialReportRequest Run an Upgraded Smart+ Creative Overview Report API Request
type MaterialReportRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Dimensions Dimension grouping conditions.
	// For example, ["campaign_id", "main_material_id"] indicates that both campaign_id and main_material_id are grouped.
	// To learn about available dimensions and supported dimension groupings, see Supported dimensions.
	Dimensions []string `json:"dimensions,omitempty"`
	// Metrics Metrics to query.
	// Default value: ["spend", "impressions"].
	// Note that the amount of currency metrics is based on the currency of the ad account (advertiser_id).
	// To learn about the supported metrics at the ad level, see Supported ad-level metrics for Upgraded Smart+ Creative Reports.
	Metrics []string `json:"metrics,omitempty"`
	// Filtering Filtering conditions.
	Filtering *MaterialReportOverviewFilter `json:"filtering,omitempty"`
	// QueryLifetime Whether to request lifetime metrics.
	// Supported values: true, false.
	// Default value: false.
	// If you set query_lifetime to true, the parameters start_date and end_date will be ignored.
	QueryLifetime bool `json:"query_lifetime,omitempty"`
	// StartDate Required when query_lifetime is false or not specified.
	// Query start date (closed interval) in the format of YYYY-MM-DD.
	// The date is based on the ad account time zone.
	StartDate string `json:"start_date,omitempty"`
	// EndDate Required when query_lifetime is false or not specified.
	// Query end date (closed interval) in the format of YYYY-MM-DD.
	// The date is based on the ad account time zone.
	// Note: The time difference between start_date and end_date must be equal to or less than 365 days.
	EndDate string `json:"end_date,omitempty"`
	// SortField Sorting field.
	// Default sorting field: spend.
	// All supported metrics (excluding attribute metrics) support sorting.
	SortField string `json:"sort_field,omitempty"`
	// SortType Sorting order.
	// Enum values: ASC (ascending), DESC(descending).
	// Default value: DESC.
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Page Current page number.
	// Value range: ≥1.
	// Default value: 1.
	Page int `json:"page,omitempty"`
	// PageSize Page size.
	// Value range: 1-100.
	// Default value: 10.
	PageSize int `json:"page_size,omitempty"`
}

type MaterialReportOverviewFilter struct {
	// CampaignIDs A list of campaign IDs.
	// Max size: 200.
	// To obtain campaign IDs, use /smart_plus/campaign/get/.
	// Filter combination rule: You cannot combine the filter campaign_ids with the filter adgroup_ids or smart_plus_ad_ids.
	// Filter and dimension combination rule: You can combine the filter with any of the following dimensions:
	// advertiser_id
	// campaign_id
	// adgroup_id
	// smart_plus_ad_id
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// AdgroupIDs A list of ad group IDs.
	// Max size: 200.
	// To obtain ad group IDs, use /smart_plus/adgroup/get/.
	// Filter combination rule: You cannot combine the filter adgroup_ids with the filter campaign_ids or smart_plus_ad_ids.
	// Filter and dimension combination limitation: You can combine the filter with any of the following dimensions:
	// advertiser_id
	// adgroup_id
	// smart_plus_ad_id
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// SmartPlusAdIDs A list of ad IDs.
	// Max size: 200.
	// To obtain ad IDs, use /smart_plus/ad/get/.
	// Filter combination rule: You cannot combine the filter smart_plus_ad_ids with the filter campaign_ids or adgroup_ids.
	// Filter and dimension combination rule: You can combine the filter with any of the following dimensions:
	// advertiser_id
	// smart_plus_ad_id
	SmartPlusAdIDs []string `json:"smart_plus_ad_ids,omitempty"`
	// MainMaterialTypes A list of main creative types.
	// Enum values:
	// VIDEO_NON_SPARK_ADS: Video (non-Spark Ads).
	// VIDEO_SPARK_ADS: Video (Spark Ads).
	// CAROUSEL_NON_SPARK_ADS: Carousel (non-Spark Ads).
	// CAROUSEL_SPARK_ADS: Carousel (Spark Ads).
	// SINGLE_IMAGE_NON_SPARK_ADS: Single image (non-Spark Ads).
	// CATALOG_VIDEO: Catalog video.
	// CATALOG_VIDEO_TEMPLATE: Catalog video template.
	// CATALOG_CAROUSEL: Catalog carousel.
	// CATALOG_MULTI_SHOW: Catalog Multi-Show. Multi-Show Experience is an auto-play video carousel experience designed to drive user exploration and engagement across a breadth of personally-relevant title offerings within your content library.
	//
	// Filter combination rule: The filter main_material_types can be used alone.
	// Filter and dimension combination rule: You can combine the filter with any of the following dimensions:
	// advertiser_id
	// campaign_id
	// adgroup_id
	// smart_plus_ad_id
	MainMaterialTypes []string `json:"main_material_types,omitempty"`
	// MainMaterialIDs A list of of main creative IDs.
	// Filter combination rule: When you use this filter, you can specify up to 400 IDs in total for the filters main_material_ids, ad_text_entity_ids, call_to_action_entity_ids, and interactive_add_on_entity_ids.
	// Filter and dimension combination rule: You can combine the filter with any of the following dimensions:
	// advertiser_id
	// campaign_id
	// adgroup_id
	// smart_plus_ad_id
	// Note: When you use the filter main_material_id, you need to specify the filter main_material_types simultaneously and filter only one type.
	MainMaterialIDs []string `json:"main_material_ids,omitempty"`
	// AdTextEntityIDs A list of ad text entity IDs.
	// The ID is a system-generated unique identifier for an ad text.
	// Filter combination rule: When you use this filter, you can specify up to 400 IDs in total for the filters main_material_ids, ad_text_entity_ids, call_to_action_entity_ids, and interactive_add_on_entity_ids.
	// Filter and dimension combination rule: You can combine the filter with any of the following dimensions:
	// advertiser_id
	// campaign_id
	// adgroup_id
	// smart_plus_ad_id
	AdTextEntityIDs []string `json:"ad_text_entity_ids,omitempty"`
	// CallToActionEntityIDs A list of call-to-action entity IDs.
	// The ID is a system-generated unique identifier for a call-to-action.
	// Filter combination rule: When you use this filter, you can specify up to 400 IDs in total for the filters main_material_ids, ad_text_entity_ids, call_to_action_entity_ids, and interactive_add_on_entity_ids.
	// Filter and dimension combination rule: You can combine the filter with any of the following dimensions:
	// advertiser_id
	// campaign_id
	// adgroup_id
	// smart_plus_ad_id
	CallToActionEntityIDs []string `json:"call_to_action_entity_ids,omitempty"`
	// InteractiveAddOnEntityIDs A list of interactive add-on entity IDs.
	// The ID is a system-generated unique identifier for an interactive add-on.
	// Filter combination rule: When you use this filter, you can specify up to 400 IDs in total for the filters main_material_ids, ad_text_entity_ids, call_to_action_entity_ids, and interactive_add_on_entity_ids.
	// Filter and dimension combination rule: You can combine the filter with any of the following dimensions:
	// advertiser_id
	// campaign_id
	// adgroup_id
	// smart_plus_ad_id
	InteractiveAddOnEntityIDs []string `json:"interactive_add_on_entity_ids,omitempty"`
}

// Encode impressions GetRequest interface
func (r *MaterialReportRequest) Encode() string {
	values := util.NewURLValues()
	if r.AdvertiserID != "" {
		values.Set("advertiser_id", r.AdvertiserID)
	}
	if len(r.Dimensions) > 0 {
		values.Set("dimensions", string(util.JSONMarshal(r.Dimensions)))
	}
	if len(r.Metrics) > 0 {
		values.Set("metrics", string(util.JSONMarshal(r.Metrics)))
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.QueryLifetime {
		values.Set("query_lifetime", "true")
	}
	if r.StartDate != "" {
		values.Set("start_date", r.StartDate)
	}
	if r.EndDate != "" {
		values.Set("end_date", r.EndDate)
	}
	if r.SortField != "" {
		values.Set("sort_field", r.SortField)
	}
	if r.SortType != "" {
		values.Set("sort_type", string(r.SortType))
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

// MaterialReportResponse Run an Upgraded Smart+ Creative Overview Report API Response
type MaterialReportResponse struct {
	model.BaseResponse
	Data *MaterialReportResult `json:"data,omitempty"`
}

type MaterialReportResult struct {
	// List 分页信息
	List []Stat `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type Stat struct {
	// Dimensions Dimension data
	Dimensions *Dimensions `json:"info,omitempty"`
	// Metrics 指标。
	Metrics *report.Metrics `json:"metrics,omitempty"`
}

type Dimensions struct {
	// AdvertiserID Returned when the parameter dimensions contains advertiser_id in the request.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID Returned when the parameter dimensions contains campaign_id in the request.
	CampaignID string `json:"campaign_id,omitempty"`
	// AdgroupID Returned when the parameter dimensions contains adgroup_id in the request.
	AdgroupID string `json:"adgroup_id,omitempty"`
	// SmartPlusAdID Returned when the parameter dimensions contains smart_plus_ad_idin the request.
	SmartPlusAdID string `json:"smart_plus_ad_id,omitempty"`
	// MainMaterialID Returned when the parameter dimensions contains main_material_id in the request.
	// Material ID.
	// When main_material_type is VIDEO_NON_SPARK_ADS, CATALOG_VIDEO, or SINGLE_IMAGE_NON_SPARK_ADS, this field represents the ID of a main creative.
	// When main_material_type is VIDEO_SPARK_ADS or CAROUSEL_SPARK_ADS, this field represents the ID of TikTok post.
	// When main_material_type is CATALOG_VIDEO_TEMPLATE, CATALOG_CAROUSEL, or CATALOG_MULTI_SHOW, this field represents the ID of a catalog.
	// When main_material_type is CAROUSEL_NON_SPARK_ADS, this field represents the ID of a carousel.
	MainMaterialID string `json:"main_material_id,omitempty"`
	// MainMaterialType Returned when the parameter dimensions contains main_material_id in the request.
	// Main creative type.
	// Enum values:
	// VIDEO_NON_SPARK_ADS: Video (non-Spark Ads).
	// VIDEO_SPARK_ADS: Video (Spark Ads).
	// CAROUSEL_NON_SPARK_ADS: Carousel (non-Spark Ads).
	// CAROUSEL_SPARK_ADS: Carousel (Spark Ads).
	// SINGLE_IMAGE_NON_SPARK_ADS: Single image (non-Spark Ads).
	// CATALOG_VIDEO: Catalog video.
	// CATALOG_VIDEO_TEMPLATE: Catalog video template.
	// CATALOG_CAROUSEL: Catalog carousel.
	// CATALOG_MULTI_SHOW: Catalog Multi-Show. Multi-Show Experience is an auto-play video carousel experience designed to drive user exploration and engagement across a breadth of personally-relevant title offerings within your content library.
	MainMaterialType string `json:"main_material_type,omitempty"`
	// AdTextEntityID Ad text entity ID.
	// The ID is a system-generated unique identifier for an ad text.
	AdTextEntityID string `json:"ad_text_entity_id,omitempty"`
	// AdTextEntityType Ad text entity type.
	// Enum value:
	// text: ad text.
	AdTextEntityType string `json:"ad_text_entity_type,omitempty"`
	// CallToActionEntityID Call-to-action entity ID.
	// The ID is a system-generated unique identifier for a call-to-action.
	CallToActionEntityID string `json:"call_to_action_entity_id,omitempty"`
	// CallToActionEntityType Call-to-action entity type.
	// Enum values:
	// manual_call_to_action: standard call to action.
	// dynamic_call_to_action: dynamic call to action.
	CallToActionEntityType string `json:"call_to_action_entity_type,omitempty"`
	// InteractiveAddOnEntityID Interactive add-on entity ID.
	// The ID is a system-generated unique identifier for an interactive add-on.
	InteractiveAddOnEntityID string `json:"interactive_add_on_entity_id,omitempty"`
	// InteractiveAddOnEntityType Interactive add-on entity type.
	// Enum value:
	// add_on: interactive add-on.
	InteractiveAddOnEntityType string `json:"interactive_add_on_entity_type,omitempty"`
}

type Metrics struct {
	report.Metrics
	// AdMaterialIDs An ad-specific material ID generated when a particular creative is used in an ad.
	// This ID differs from the creative ID you receive when uploading the creative to your ad account’s Creative Library.
	AdMaterialID []string `json:"ad_material_id,omitempty"`
	// MultiName A combination of setting names based on the setting dimensions (ad_text_entity_ids, call_to_action_entity_ids, and interactive_add_on_entity_ids) specified in the request.
	// For instance, if you specify the dimensions ad_text_entity_ids and call_to_action_entity_ids in the request, the multi_name will be {ad_text_name}_{call_to_action_name}.
	// Example: Video1_Adtext1_CTA1.
	MultiName string `json:"multi_name,omitempty"`
	// MainMaterialName Main creative name
	MainMaterialName string `json:"main_material_name,omitempty"`
	// AdText Ad text
	AdText string `json:"ad_text,omitempty"`
	// CallToAction Call-to-action
	CallToAction string `json:"call_to_action,omitempty"`
	// InteractiveAddOnName Interactive add-on name.
	InteractiveAddOnName string `json:"interactive_add_on_name,omitempty"`
	// InteractiveAddOnID Interactive add-on ID.
	// This ID is generated when you create the interactive add-on in the Creative Library and is different from interactive_add_on_entity_id.
	InteractiveAddOnID string `json:"interactive_add_on_id,omitempty"`
	// MaterialPrimaryStatus Returned when the parameter dimensions contains smart_plus_ad_id in the request.
	// The primary status of the creative.
	// For enum values, see List of values for material_primary_status.
	MaterialPrimaryStatus enum.PrimaryStatus `json:"material_primary_status,omitempty"`
	// MaterialSecondStatusList The list of secondary statuses of the creative.
	// For enum values, see List of values for material_second_status_list.
	// Note: When the parameter dimensions doesn't contain smart_plus_ad_id in the request, this field will not be returned.
	MaterialSecondStatusList []enum.SecondaryStatus `json:"material_secondary_status,omitempty"`
	// OperationStatus Operation status.
	// Enum values:
	// ENABLE : enabled (in 'ON' status).
	// DISABLE: disabled (in 'OFF' status).
	// FROZEN: terminated and cannot be enabled.
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
}
