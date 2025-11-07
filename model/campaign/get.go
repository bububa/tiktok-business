package campaign

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取推广系列列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Fields 希望返回的字段集合。
	// 不传时，默认返回所有字段。
	// 传入时则仅返回指定字段。枚举值参见返回中list下的字段。
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
	PageSize int `json:"page_size,omitempty"`
}

// GetFilter 筛选条件。
type GetFilter struct {
	// CampaignIDs 推广系列 ID，允许数量范围: 1-100
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// CampaignName 推广系列名称，支持模糊匹配
	CampaignName string `json:"campaign_name,omitempty"`
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
	// CampaignType 推广系列类型。枚举值: REGULAR_CAMPAIGN（普通推广系列）、IOS14_CAMPAIGN（iOS14推广系列）。更多关于iOS14推广系列的信息，可参考 创建iOS14 推广系列。
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// ObjectiveType 推广目标，见 枚举值-推广目标
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
	// PrimaryStatus 一级状态。枚举值详见 枚举值 - 一级状态
	PrimaryStatus enum.PrimaryStatus `json:"primary_status,omitempty"`
	// SecondaryStatus 推广系列二级状态。枚举值详见枚举值 - 二级状态 - 推广系列状态。
	// 注意：
	// 沙箱环境下本字段不返回，因为推广系列未实际投放。
	// 请查看一级状态下所支持的二级状态 ，了解对于所指定的一级状态，您可以传入的secondary_status 有效值。
	SecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// CreationFilterStartTime 推广系列最早创建时间，格式：YYYY-MM-DD HH:MM:SS（UTC时区）。创建时间晚于此时间的推广系列会被返回。
	// 建议：为了让任务成功执行及任务速度不受影响，创建时间的范围建议在6个月以内。
	CreationFilterStartTime model.DateTime `json:"creation_filter_start_time,omitzero"`
	// CreationFilterEndTime 推广系列最晚创建时间，格式：YYYY-MM-DD HH:MM:SS（UTC时区）。创建时间先于此时间的推广系列会被返回。
	// 建议：为了让任务成功执行及任务速度不受影响，创建时间的范围建议在6个月以内。
	CreationFilterEndTime model.DateTime `json:"creation_filter_end_time,omitzero"`
	// IsSmartPerformanceCampaign 按是否为自动化类型的推广系列筛选。
	// 支持的值：
	// true：推广系列为 Smart+ 推广系列或智能效果网站推广系列。
	// false：推广系列为普通类型的推广系列。
	//
	// 注意：若您不填写该字段，则将获取所有推广系列。
	IsSmartPerformanceCampaign *bool `json:"is_smart_performance_campaign,omitempty"`
	// SplitTestEnabled 按推广系列是否启用了拆分对比测试筛选。
	// true：仅获取已启用拆分对比测试的推广系列，false：仅获取未启用拆分对比测试的推广系列。
	// 注意：若您不填写该字段，则将获取所有推广系列。
	SplitTestEnabled *bool `json:"split_test_enabled,omitempty"`
	// CampaignProductSource 按推广系列的商品来源筛选。
	// 枚举值：
	// CATALOG ：商品库。
	// STORE：TikTok Shop 店铺，或 TikTok 橱窗。
	// 若您传入本字段，返回结果仅包含推广目标为商品销量的推广系列（objective_type=PRODUCT_SALES）。若您同时传入筛选字段objective_type，则objective_type仅可设置为PRODUCT_SALES。
	CamtpaignProductSource enum.ProductSource `json:"campaign_product_source,omitempty"`
	// OptimizationGoal 优化目标。枚举值见枚举值-优化目标。
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// GMVMaxPromotionTypes GMV Max 推广系列类型。
	// 枚举值：
	// PRODUCT_GMV_MAX：商品 GMV Max 推广系列。
	// LIVE_GMV_MAX：直播 GMV Max 推广系列。
	// 注意：若想筛选 GMV Max 推广系列，需使用/gmv_max/campaign/get/。
	GMVMaxPromotionTypes []enum.GMVMaxPromotionType `json:"gmv_max_promotion_types,omitempty"`
	// StoreIDs 仅当 gmv_max_promotion_types 的值中包含 LIVE_GMV_MAX 或PRODUCT_GMV_MAX 时本字段生效。
	// TikTok Shop ID 列表。
	// 最大数量：10。
	// 注意：若想筛选 GMV Max 推广系列，需使用/gmv_max/campaign/get/。
	StoreIDs []string `json:"store_ids,omitempty"`
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

// GetResponse 获取推广系列列表 API Response
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
