package gmvmax

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ReportGetRequest 生成 GMV Max 推广系列报表 API Request
type ReportGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreIDs TikTok Shop ID 列表。
	// 最大数量：1。
	// 若想获取可用于 GMV Max 推广系列的 TikTok Shop，可使用 /gmv_max/store/list/ 并确认返回的 is_gmv_max_available 为 true
	StoreIDs []string `json:"store_ids,omitempty"`
	// StartDate 查询起始日期（闭区间），格式为 YYYY-MM-DD。日期基于广告账户时区
	StartDate string `json:"start_date,omitempty"`
	// EndDate 查询终止日期（闭区间），格式为 YYYY-MM-DD。日期基于广告账户时区。
	// 示例：2025-02-01。
	// 若 dimensions 字段值不包含 stat_time_day 和 stat_time_hour，则通过 start_date 和 end_date 指定的时间区间最大为 365 天。
	// 若 dimensions 字段值包含 stat_time_day，则通过 start_date 和 end_date 指定的时间区间最大为 30 天。
	// 若 dimensions 字段值包含 stat_time_hour，则通过 start_date 和 end_date 指定的时间区间最大为 1 天
	EndDate string `json:"end_date,omitempty"`
	// Metrics 要查询的指标。
	// 若想了解可为所有 GMV Max 推广系列、商品 GMV Max 推广系列和直播 GMV Max 推广系列分别请求的指标类型，可查看 GMV Max 推广系列报表指标
	Metrics []string `json:"metrics,omitempty"`
	// EnableTotalMetrics 是否开启所请求指标的汇总数据。
	// 当启用 enable_total_metrics 时，我们将在您查询不同页面时提供所有页面的汇总数据。
	// 支持的值：true，false。
	// 默认值：false
	EnableTotalMetrics bool `json:"enable_total_metrics,omitempty"`
	// Dimensions 维度组合，用于对结果进行分组。
	// 枚举值：
	// advertiser_id：按广告主 ID 分组。
	// campaign_id：按推广系列 ID 分组。
	// stat_time_day：按天分组。
	// stat_time_hour：按小时分组。
	// item_group_id：按 SPU ID 分组。
	// item_id：按帖子 ID 分组。
	// room_id：按直播间 ID 分组。
	// duration：按时长分组。
	// 例如，["campaign_id", "stat_time_day"] 表示组合 campaign_id 和 stat_time_day （天）维度。
	// 可用的维度组合取决于您想查询的指标类型。若想了解您想请求的指标类型支持的维度组合，可查看 GMV Max 推广系列报表指标
	Dimensions []string `json:"dimensions,omitempty"`
	// Filtering 筛选条件。
	// 若想了解您想请求的指标类型支持的筛选条件，可查看 GMV Max 推广系列报表指标
	Filtering *ReportGetFilter `json:"filtering,omitempty"`
	// SortField 排序字段。
	// 默认不进行排序。
	// 推荐您使用排序字段，确保结果按顺序返回。若想查看可选的字段，可查看支持的排序字段
	SortField string `json:"sort_field,omitempty"`
	// SortType 排序方式。
	// 枚举值：ASC（升序）， DESC（降序）。
	// 默认值：DESC
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Page 当前页数。
	// 取值范围：≥1。
	// 默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 取值范围：1-1,000。
	// 默认值：10
	PageSize int `json:"page_size,omitempty"`
}

type ReportGetFilter struct {
	// GMVMaxPromotionTypes 按 GMV Max 推广系列类型筛选。
	// 枚举值：
	// PRODUCT：商品 GMV Max 推广系列。
	// LIVE：直播 GMV Max 推广系列
	GMVMaxPromotionTypes []enum.GMVMaxPromotionType `json:"gmv_max_promotion_types,omitempty"`
	// CampaignIDs 按 GMV Max 推广系列 ID 列表筛选。
	// 最大数量：100
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// CampaignName 按 GMV Max 推广系列名称筛选。
	// 支持模糊搜索
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignStatuses 按推广系列状态筛选。
	// 枚举值：
	// STATUS_DELIVERY_OK：投放中。
	// STATUS_DISABLE：未投放。
	// STATUS_DELETE：已删除
	CampaignStatuses []enum.PrimaryStatus `json:"campaign_statuses,omitempty"`
	// ItemGroupIDs 按商品 SPU ID 列表筛选。
	// 最大数量：100
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
	// CreativeTypes 使用 campaign_ids 和 item_group_ids 筛选单个 SPU ID 时必填。
	// 在 campaign_ids 或 item_group_ids 中指定多个 ID，或同时指定多个 ID 维度时不支持本字段。
	// 按创意类型筛选。
	// 枚举值：
	// ADS_AND_ORGANIC：广告和原生。此类型包含已获授权并用于广告的所有创意素材。此类型中视频的总收入包括广告和原生视频产生的销售额。
	// ORGANIC：仅原生。此类型包含仅用于自然流量的视频。此类型中视频的总收入仅包括原生视频产生的销售额。
	// REMOVED：已移除。此类型包含自动选择的、但已从您的推广系列中手动移除的视频。
	// 最大数量：1。
	CreativeTypes []string `json:"creative_types,omitempty"`
	// CreativeDeliveryStatuses 此筛选条件仅支持在 campaign_ids 或 item_group_ids 中指定多个 ID，或同时指定多个 ID 维度时使用。
	// 此筛选条件不可与 creative_types 一起使用。
	// 按创意状态筛选。
	// 枚举值：
	// IN_QUEUE：排队中。作品已准备好进行转化驱动潜力测试。
	// LEARNING：学习中。作品正在接受转化驱动潜力测试。
	// DELIVERING：投放中。作品已作为推广系列的一部分进行定期推广。
	// NOT_DELIVERING：未投放。作品在测试过程中未表现出强大的转化潜力，不会作为广告计划的一部分进行定期推广。
	// AUTHORIZATION_NEEDED：需要授权。此视频尚未被授权用于广告。您可以从 TikTok 账号授权自己的视频，或向其他达人申请授权。在授权后 20 分钟内，此处的状态将会更新。
	// EXCLUDED：已排除。作品已被手动删除，无法作为广告计划的一部分进行定期推广。您可以随时将此作品重新添加到你的广告计划中。
	// UNAVAILABLE：不可用。私密视频、被拒视频、已删除视频或来自已禁用 TikTok 账号的视频不能用于广告。在某些情况下，您可能会看到来自不可用视频的消耗，这是因为在推广系列投放期间被 TikTok 内容审核流程拒绝的视频仍可能会投放一段时间。
	// REJECTED：拒绝。视频已被内容审查团队拒绝。如果您已提交申诉，那么在申诉审核完成前，视频状态将保持为“已拒绝”。
	// NOT_ACTIVE：不活跃。此视频已被降级处理，因为它上传超过 30 天且近 30 天内未产生任何收入。
	CreativeDeliveryStatuses []enum.CreativeDeliveryStatus `json:"creative_delivery_statuses,omitempty"`
	// SearchWord 按搜索关键词筛选，例如视频标题、TikTok 帖子 ID 或 TikTok 账号名称。
	// 支持模糊搜索
	SearchWord string `json:"search_word,omitempty"`
	// RoomIDs 按直播间 ID 列表筛选。
	// 最大数量：100
	RoomIDs []string `json:"room_ids,omitempty"`
}

// Encode implements GetRequest
func (r *ReportGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("store_ids", string(util.JSONMarshal(r.StoreIDs)))
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	values.Set("metrics", string(util.JSONMarshal(r.Metrics)))
	values.Set("dimensions", string(util.JSONMarshal(r.Dimensions)))
	if r.EnableTotalMetrics {
		values.Set("enable_total_metrics", "true")
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// ReportGetResponse 生成 GMV Max 推广系列报表 API Response
type ReportGetResponse struct {
	model.BaseResponse
	Data *ReportGetResult `json:"data,omitempty"`
}

type ReportGetResult struct {
	// TotalMetrics 所有请求指标的汇总数据。 当您在请求中开启 enable_total_metrics 参数时，返回此对象。
	TotalMetrics *Metrics `json:"total_metrics,omitempty"`
	// List 数据列表
	List []Stat `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type Stat struct {
	// Dimensions 所有请求的维度组合数据
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	// Metrics 所有请求的指标数据
	Metrics *Metrics `json:"metrics,omitempty"`
}

type Dimensions struct {
	// AdvertiserID 广告主ID。当dimensions包含 advertiser_id 时返回
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列ID。当dimensions包含 campaign_id 时返回
	CampaignID string `json:"campaign_id,omitempty"`
	// StatTimeDay 消耗发生的时间（天）。格式：2020-01-01 00:00:00
	StatTimeDay model.DateTime `json:"stat_time_day,omitzero"`
	// StatTimeHour 消耗发生的时间（小时）。格式：2020-01-01 10:00:00
	StatTimeHour model.DateTime `json:"stat_time_hour,omitzero"`
	// ItemGroupID SPU ID
	ItemGroupID string `json:"item_group_id,omitempty"`
	// ItemID 帖子 ID
	ItemID string `json:"item_id,omitempty"`
	// RoomID 直播间 ID
	RoomID string `json:"room_id,omitempty"`
	// Duration 时长
	Duration string `json:"duration,omitempty"`
}

type Metrics struct {
	// CampaignID 推广系列ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName 推广系列名称
	CampaignName string `json:"campaign_name,omitempty"`
	// OperationStatus 推广系列开/关状态
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// ScheduleType 排期类型
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime 排期起始时间
	ScheduleStartTime model.DateTime `json:"schedule_start_time,omitzero"`
	// ScheduleEndTime 排期终止时间
	ScheduleEndTime model.DateTime `json:"schedule_end_time,omitzero"`
	// TargetROIBudget 目标投资回报率 (ROI) 预算
	TargetROIBudget model.Float64 `json:"target_roi_budget,omitempty"`
	// BidType 优化模式
	BidType enum.BidType `json:"bid_type,omitempty"`
	// MaxDeliveryBudget 最大投放量预算
	MaxDeliveryBudget model.Float64 `json:"max_delivery_budget,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"product_name,omitempty"`
	// ItemGroupID 商品 SPU ID
	ItemGroupID string `json:"item_group_id,omitempty"`
	// ProductImageURL 商品图片 URL
	ProductImageURL string `json:"product_image_url,omitempty"`
	// Title 创意素材名称
	Title string `json:"title,omitempty"`
	// ItemID TikTok 帖子 ID
	ItemID string `json:"item_id,omitempty"`
	// TTAccountName TikTok 账号名称
	TTAccountName string `json:"tt_account_name,omitempty"`
	// TTAccountProfileImageURL TikTok 账号头像图片 URL
	TTAccountProfileImageURL string `json:"tt_account_profile_image_url,omitempty"`
	// TTAccountAuthorizationType 授权类型
	TTAccountAuthorizationType string `json:"tt_account_authorization_type,omitempty"`
	// ShopContentType 店铺内容类型
	ShopContentType string `json:"shop_content_type,omitempty"`
	// IdentityID 认证身份 ID （TikTok 账号 ID）
	IdentityID string `json:"identity_id,omitempty"`
	// LiveName 直播名称
	LiveName string `json:"live_name,omitempty"`
	// LiveStatus 直播状态
	LiveStatus string `json:"live_status,omitempty"`
	// LiveLaunchedTime 直播开始时间
	LiveLaunchedTime model.DateTime `json:"live_launched_time,omitzero"`
	// LiveDuration 直播时长
	LiveDuration model.Int64 `json:"live_duration,omitempty"`
	//
	// Cost 成本
	Cost model.Float64 `json:"cost,omitempty"`
	// Orders 订单数（SKU）
	Orders model.Int64 `json:"orders,omitempty"`
	// CostPerOrder 平均下单成本
	CostPerOrder model.Float64 `json:"cost_per_order,omitempty"`
	// GrossRevenue 总收入
	GrossRevenue model.Float64 `json:"gross_revenue,omitempty"`
	// ROI ROI
	ROI model.Float64 `json:"roi,omitempty"`
	// NetCost 净成本
	NetCost model.Float64 `json:"net_cost,omitempty"`
	// RoasBid 目标 ROI
	RoasBid model.Float64 `json:"roas_bid,omitempty"`
	// ProductStatus 商品状态
	ProductStatus string `json:"product_status,omitempty"`
	// CreativeDeliveryStatus #N/A
	CreativeDeliveryStatus enum.CreativeDeliveryStatus `json:"creative_delivery_status,omitempty"`
	// ProductImpressions 商品广告曝光数
	ProductImpressions model.Int64 `json:"product_impressions,omitempty"`
	// ProductClicks 商品广告点击数
	ProductClicks model.Int64 `json:"product_clicks,omitempty"`
	// ProductClickRate 商品广告点击率
	ProductClickRate model.Float64 `json:"product_click_rate,omitempty"`
	// AllShopsCostPerOrder 平均下单成本
	AllShopsCostPerOrder model.Float64 `json:"all_shops_cost_per_order,omitempty"`
	// AllShopsOrders 订单数
	AllShopsOrders model.Int64 `json:"all_shops_orders,omitempty"`
	// AllShopsGrossRevenue 总收入
	AllShopsGrossRevenue model.Float64 `json:"all_shops_gross_revenue,omitempty"`
	// AllShopsROI ROI
	AllShopsROI model.Float64 `json:"all_shops_roi,omitempty"`
	// AdClickRate 广告点击率
	AdClickRate model.Float64 `json:"ad_click_rate,omitempty"`
	// AdConversionRate 广告转化率
	AdConversionRate model.Float64 `json:"ad_conversion_rate,omitempty"`
	// AdVideoViewRate2s 2 秒广告视频播放率
	AdVideoViewRate2s model.Float64 `json:"ad_video_view_rate_2s,omitempty"`
	// AdVideoViewRate6s 6 秒广告视频播放率
	AdVideoViewRate6s model.Float64 `json:"ad_video_view_rate_6s,omitempty"`
	// AdVideoViewRateP25 25% 广告视频播放率
	AdVideoViewRateP25 model.Float64 `json:"ad_video_view_rate_p25,omitempty"`
	// AdVideoViewRateP50 50% 广告视频播放率
	AdVideoViewRateP50 model.Float64 `json:"ad_video_view_rate_p50,omitempty"`
	// AdVideoViewRateP75 75% 广告视频播放率
	AdVideoViewRateP75 model.Float64 `json:"ad_video_view_rate_p75,omitempty"`
	// AdVideoViewRateP100 100% 广告视频播放率
	AdVideoViewRateP100 model.Float64 `json:"ad_video_view_rate_p100,omitempty"`
	// LiveViews 直播播放量
	LiveViews model.Int64 `json:"live_views,omitempty"`
	// CostPerLiveView 直播播放平均成本
	CostPerLiveView model.Float64 `json:"cost_per_live_view,omitempty"`
	// TenSecondLiveViews 直播播放达 10 秒播放量
	TenSecondLiveViews model.Int64 `json:"10_second_live_views,omitempty"`
	// CostPer10SecondLiveViews 直播播放达 10 秒平均成本
	CostPer10SecondLiveViews model.Float64 `json:"cost_per_10_second_live_views,omitempty"`
	// LiveFollows 直播关注数
	LiveFollows model.Int64 `json:"live_follows,omitempty"`
}
