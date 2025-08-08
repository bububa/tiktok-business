package ad

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取广告 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Fields 返回字段。查看返回中list下的字段，获取您可以指定的字段。如果您不传入该字段，系统将默认返回所有字段。
	Fields []string `json:"fields,omitempty"`
	// ExcludeFieldTypesInResponse 想要从返回中移除的字段种类。
	// 允许的枚举值：
	// NULL_FIELD: 值为null的字段。
	ExcludeFieldTypesInResponse []enum.ExcludeFieldType `json:"exclude_field_types_in_response,omitempty"`
	// Filtering 筛选条件
	Filtering *GetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1。
	// 取值范围：≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值: 10。
	// 取值范围：1-1,000。
	PageSize int `json:"page_size,omitempty"`
}

// GetFilter 筛选条件
type GetFilter struct {
	// CampaignIDs 推广系列 ID，允许数量范围: 1-100
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// CampaignSystemOrigins 推广系列来源。
	// 枚举值：
	// PROMOTE：该推广系列为通过 TikTok 移动应用创建的内容加热推广系列。
	// TT_ADS_PLATFORM：该推广系列为通过 API 或TikTok 广告管理平台创建的非内容加热推广系列。
	// 默认值：["TT_ADS_PLATFORM"] 。
	//
	// 仅支持获取内容加热推广系列的以下信息：
	// advertiser_id
	// campaign_id
	// create_time
	// campaign_name
	// operation_status
	// secondary_status
	//
	// 若想了解支持的内容加热广告功能及推广系列筛选条件，请查看内容加热推广系列。
	CampaignSystemOrigins []enum.CampaignSystemOrigin `json:"campaign_system_origins,omitempty"`
	// AdgroupIDs 广告组id列表，支持筛选指定的广告组，允许数量范围：1-100
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// AdIDs 广告 ID 列表，允许数量范围: 1-100
	AdIDs []string `json:"ad_ids,omitempty"`
	// PrimaryStatus 一级状态。枚举值详见 枚举值 - 一级状态
	// 默认值：STATUS_NOT_DELETE，返回除STATUS_DELETE外所有状态的广告组。如果您想获得包括STATUS_DELETE在内的所有状态的广告组，请使用STATUS_ALL。
	PrimaryStatus enum.PrimaryStatus `json:"primary_status,omitempty"`
	// SecondaryStatus 二级状态。枚举值详见枚举值 - 二级状态 - 推广系列状态。
	// 注意：
	// 沙箱环境下本字段不返回，因为推广系列未实际投放。
	// 请查看一级状态下所支持的二级状态 ，了解对于所指定的一级状态，您可以传入的secondary_status 有效值。
	SecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// ObjectiveType 推广目标，见 枚举值-推广目标
	// 注意：若您选择WEB_CONVERSIONS或CONVERSIONS的推广目标作为筛选条件，目前会同时返回以WEB_CONVERSIONS或CONVERSIONS为推广目标的数据。
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// BuyingTypes 购买类型。
	// 枚举值：
	// AUCTION：竞价广告。
	// RESERVATION_RF：合约覆盖和频次广告以及 TikTok Pulse 广告。
	// RESERVATION_TOP_VIEW：合约 TopView 广告。
	//
	// 默认值：["AUCTION", "RESERVATION_RF"]。
	//
	// 若想了解 API 中支持的 TopView 广告功能，请查看 TopView。
	//
	// 注意： 枚举值 RESERVATION_TOP_VIEW 不支持与其他值同时使用。若想筛选 TopView 广告，可将本字段设置为["RESERVATION_TOP_VIEW"]。
	BuyingTypes []enum.BuyingType `json:"buying_types,omitempty"`
	// OptimizationGoal 优化目标。枚举值见枚举值-优化目标。
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// CreativeMaterialMode 创意投放方式。
	// 枚举值: CUSTOM（自定义）、DYNAMIC（程序化）、SMART_CREATIVE（智能创意）
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// Destination 目标页类型。
	// 枚举值：
	// APP：所推广 App 在Google Play 或 Apple Store 的下载页面。广告对应的广告组中promotion_type 设置为APP_ANDROID 或 APP_IOS。
	// TIKTOK_INSTANT_PAGE: TikTok 即时体验页面，包括自定义页面、线索表单（即时体验表单）、应用介绍页和即时体验商品页面。广告中指定了TikTok即时体验页面的页面 ID（page_id）。
	// WEBSITE：网站页面。广告中指定了落地页（landing_page_url）。
	// SOCIAL_MEDIA_APP：社交媒体 URL。若想了解支持将目标页面设置为社交媒体 URL 的广告类型，请查看创建推广对象类型为即时通讯应用的线索广告。
	// PHONE_CALL：电话号码。若想了解支持将目标页面设置为电话号码的广告类型，请查看创建推广对象类型为电话通话的线索广告。
	//
	// 若您同时指定了objective_type ，请查看推广目标下支持筛选的目标页面类型，了解本字段可以设置的值。
	Destination enum.Destination `json:"destination,omitempty"`
	// CreationFilterStartTime 筛选创建时间晚于某一时间的广告，格式为 YYYY-MM-DD HH:MM:SS（UTC 时区）。
	// 建议：为确保任务的成功执行及速度，创建时间的筛选范围建议不超过 6 个月。
	CreationFilterStartTime model.DateTime `json:"creation_filter_start_time,omitzero"`
	// CreationFilterEndTime 筛选创建时间早于某一时间的广告，格式为 YYYY-MM-DD HH:MM:SS（UTC 时区）。
	// 建议：为确保任务的成功执行及速度，创建时间的筛选范围建议不超过 6 个月。
	CreationFilterEndTime model.DateTime `json:"creation_filter_end_time,omitzero"`
	// ModifiedAfter 筛选修改时间晚于某一时间的广告，格式为 YYYY-MM-DD HH:MM:SS（UTC 时区）。
	// 建议：为确保任务的成功执行及速度，修改时间的筛选范围建议不超过 6 个月。
	ModifiedAfter model.DateTime `json:"modified_after,omitzero"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if len(r.ExcludeFieldTypesInResponse) > 0 {
		values.Set("exclude_field_types_in_response", string(util.JSONMarshal(r.ExcludeFieldTypesInResponse)))
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

// GetResponse 获取广告 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// List 符合条件的广告列表，请求参数未设置fields字段或者fields中包含以下子字段时会返回相关字段
	// 注意：返回中不包含通过 TikTok Shop 创建的商品总交易额最大化广告（Product GMV max ads）和直播总交易额最大化广告（LIVE GMV max ads）。通过 API 或 TikTok 广告管理平台创建的购物广告则不受影响。
	List []Ad `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
