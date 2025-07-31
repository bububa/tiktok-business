package adgroup

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取广告组列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Fields 希望返回的字段集合。
	// 不传时，默认返回所有字段。
	// 传入时则仅返回指定字段。枚举值参见返回数据中list下的字段。
	Fields []string `json:"fields,omitempty"`
	// ExcludeFieldTypesInResponse 想要从返回中移除的字段种类。
	// 允许的枚举值：
	// NULL_FIELD: 值为null的字段。
	ExcludeFieldTypesInResponse []enum.ExcludeFieldType `json:"exclude_field_types_in_response,omitempty"`
	// Filtering 筛选条件。
	// 示例： filtering={"objective_type":"APP_PROMOTION"}
	Filtering *GetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1。
	// 取值范围：≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值: 10。
	// 取值范围：1-1,000。
	// 注意：若您在筛选字段 buying_types 使用了枚举值RESERVATION_TOP_VIEW，则本字段允许的最大值为 100。
	PageSize int `json:"page_size,omitempty"`
}

// GetFilter 筛选条件。
type GetFilter struct {
	// CampaignIDs 推广系列的id列表，支持筛选指定推广系列下的广告组，允许数量范围：1-100。
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
	// AdgroupName 广告组名字，支持广告组名字的模糊搜索
	AdgroupName string `json:"adgroup_name,omitempty"`
	// PrimaryStatus 一级状态。枚举值详见 枚举值 - 一级状态
	// 默认值：STATUS_NOT_DELETE，返回除STATUS_DELETE外所有状态的广告组。如果您想获得包括STATUS_DELETE在内的所有状态的广告组，请使用STATUS_ALL。
	PrimaryStatus enum.PrimaryStatus `json:"primary_status,omitempty"`
	// SecondaryStatus 推广系列二级状态。枚举值详见枚举值 - 二级状态 - 推广系列状态。
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
	// PromotionType 您想要筛选的推广对象类型（优化位置）。
	// 枚举值：
	// APP：应用。
	// WEBSITE：落地页。
	// INSTANT_FORM ：线索表单（即时表单）。
	// LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE：TikTok 私信。
	// LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE：社交媒体应用。
	// LEAD_GEN_CLICK_TO_CALL：电话通话。
	//
	// 注意：
	//
	// 本筛选字段不同于用于创建广告组的promotion_type字段，即本接口返回的promotion_type字段。promotion_type 作为筛选字段同样支持在同步基础报表和同步DSA报表中使用。
	// 请查看不同推广对象类型之间的映射 ，了解如何使用本字段筛选出用于广告组创建的推广对象类型。
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// BidStrategy 竞价策略。 枚举值：BID_STRATEGY_COST_CAP, BID_STRATEGY_BID_CAP, BID_STRATEGY_MAX_CONVERSION 和 BID_STRATEGY_LOWEST_COST
	BidStrategy enum.BidStrategy `json:"bid_strategy,omitempty"`
	// CreativeMaterialMode 创意投放方式。
	// 枚举值: CUSTOM（自定义），SMART_CREATIVE（智能创意）。
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// BillingEvents 计费事件，按照计费事件筛选。
	// 枚举值：详见枚举值-计费事件 。
	BillingEvents []enum.BillingEvent `json:"billing_events,omitempty"`
	// CreationFilterStartTime 广告组最早创建时间，格式：YYYY-MM-DD HH:MM:SS（UTC时区）。创建时间晚于此时间的广告组会被返回。
	// 建议：为了让任务成功执行及任务速度不受影响，创建时间的范围建议在6个月以内。
	CreationFilterStartTime model.DateTime `json:"creation_filter_start_time,omitzero"`
	// CreationFilterEndTime 广告组最晚创建时间，格式：YYYY-MM-DD HH:MM:SS（UTC时区）。创建时间先于此时间的广告组会被返回。
	// 建议：为了让任务成功执行及任务速度不受影响，创建时间的范围建议在6个月以内。
	CreationFilterEndTime model.DateTime `json:"creation_filter_end_time,omitzero"`
	// SplitTestEnabled 按广告组是否启用了拆分对比测试筛选。
	// true：仅获取已启用拆分对比测试的广告组，false：仅获取未启用拆分对比测试的广告组。
	// 注意： 若您不填写该字段，则将获取所有广告组。
	SplitTestEnabled *bool `json:"split_test_enabled,omitempty"`
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

// GetResponse 获取广告组列表 API Response
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
