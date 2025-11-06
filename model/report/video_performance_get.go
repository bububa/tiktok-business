package report

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// VideoPerformanceGetRequest 获取秒级表现数据 API Request
type VideoPerformanceGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// ReportType 报表类型。
	// 枚举值：
	// AD即将废弃：提供广告秒级表现数据的报表。
	// VIDEO：视频素材的视频洞察报表。
	// 默认值：AD。
	// 注意：下个版本中将废弃枚举值 AD ，建议您使用 VIDEO 拉取对应报表数据。
	ReportType enum.ReportType `json:"report_type,omitempty"`
	// MetricsFields 你想获取的指标。默认返回所有指标。支持的指标可见支持指标表格。
	MetricsFields []string `json:"metrics_fields,omitempty"`
	// Filtering 筛选条件。你只能传入三个条件中的一个条件，且必须传入一个条件。
	Filtering *VideoPerformanceGetFilter `json:"filtering,omitempty"`
	// SortField 排序依据的字段。默认根据创建时间(CREATE_TIME)排序
	SortField string `json:"sort_field,omitempty"`
	// SortType 排序类型。枚举值: ASC (升序，从最早创建的开始), DES (降序，从最晚创建的开始)
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Page 页号
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

type VideoPerformanceGetFilter struct {
	// AdIDs 若report_type设置为AD或未传入，需传入ad_ids ，adgroup_ids，和 campaign_ids 其中之一。
	// 您想获取对应秒级表现数据的广告ID列表。
	// 注意：目前，传入的ad_ids数量不能超过200。
	AdIDs []string `json:"ad_ids,omitempty"`
	// AdgroupIDs 若report_type设置为AD或未传入，需传入ad_ids ，adgroup_ids，和 campaign_ids 其中之一。
	// 您想获取对应秒级表现数据的广告组ID列表。
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// CampaignIDs 若report_type设置为AD或未传入，需传入ad_ids ，adgroup_ids，和 campaign_ids 其中之一。
	// 您想获取对应秒级表现数据的推广系列ID列表。
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// MaterialIDs 若report_type设置为VIDEO，需传入material_ids或video_ids。
	// 您想获取对应视频洞察数据的视频素材ID列表。
	// 最大数量：1。
	MaterialIDs []string `json:"material_ids,omitempty"`
	// VideoIDs 若report_type设置为VIDEO，需传入material_ids或video_ids。
	// 您想获取对应视频洞察数据的视频ID列表。
	// 最大数量：1。
	VideoIDs []string `json:"video_ids,omitempty"`
	// StartTime 仅当report_type设置为VIDEO时生效。
	// 若您想获取指定时间区间的视频洞察数据，您需同时传入start_time和end_time，并将lifetime设置为false。
	// 查询起始时间（闭区间），格式为YYYY-MM-DD hh:mm:ss (UTC+0时间)。
	StartTime string `json:"start_time,omitempty"`
	// EndTime 仅当report_type设置为VIDEO时生效。
	// 若您想获取指定时间区间的视频洞察数据，您需同时传入start_time和end_time，并将lifetime设置为false。
	// 查询结束时间（闭区间），格式为YYYY-MM-DD hh:mm:ss (UTC+0时间)。
	EndTime string `json:"end_time,omitempty"`
	// LifeTime 仅当report_type设置为VIDEO时生效。
	// 若您想获取指定时间区间的视频洞察数据，您需同时传入start_time和end_time，并将lifetime设置为false。
	// 是否请求指标的Lifetime（所有）数据。
	// 默认值：true。
	LifeTime *bool `json:"life_time,omitempty"`
}

// Encode implements GetRequest
func (r *VideoPerformanceGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if r.ReportType != "" {
		values.Set("report_type", string(r.ReportType))
	}
	if len(r.MetricsFields) > 0 {
		values.Set("metrics_fields", string(util.JSONMarshal(r.MetricsFields)))
	}
	values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// VideoPerformanceGetResponse 获取秒级表现数据 API Response
type VideoPerformanceGetResponse struct {
	model.BaseResponse
	Data *VideoPerformanceGetResult `json:"data,omitempty"`
}

type VideoPerformanceGetResult struct {
	// List 结果列表
	List []VideoPerformance `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type VideoPerformance struct {
	// Dimensions 所有请求的维度组合数据
	Dimensions *VideoPerformanceInfo `json:"dimensions,omitempty"`
	// Metrics 所有请求的指标数据
	Metrics *VideoPerformanceMetrics `json:"metrics,omitempty"`
}

type VideoPerformanceInfo struct {
	// AdID 广告ID。当dimensions包含 ad_id 时返回
	AdID string `json:"ad_id,omitempty"`
	// VideoID 视频 ID
	VideoID string `json:"video_id,omitempty"`
	// Duration 视频时长，以秒计。
	Duration int64 `json:"duration,omitempty"`
}

type VideoPerformanceMetrics struct {
	// ViewCount 该广告在这一秒内的浏览量
	ViewCount []int64 `json:"view_count,omitempty"`
	// ClickCount 该广告在这一秒内的点击量
	ClickCount []int64 `json:"click_count,omitempty"`
	// SlideCount 该广告的受众在这一秒内从广告滑动到主页的次数
	SlideCount []int64 `json:"slide_count,omitempty"`
	// ConvertCount 该广告的受众在这一秒内完成的转化事件数
	ConvertCount []int64 `json:"convert_count,omitempty"`
	// BreakCount 在这一秒内停止观看该广告的观众数
	BreakCount []int64 `json:"break_count,omitempty"`
	// LikeCount 该广告在这一秒内获得的点赞数
	LikeCount []int64 `json:"like_count,omitempty"`
	// ShareCount 该广告在这一秒内获得的分享数
	ShareCount []int64 `json:"share_count,omitempty"`
	// CommentCount 该广告在这一秒内获得的评论数
	CommentCount []int64 `json:"comment_count,omitempty"`
	// Clicks 该视频在这一秒内获得的点击次数
	Clicks []int64 `json:"clicks,omitempty"`
	// Conversions 该视频在这一秒内完成的转化事件数
	Conversions []int64 `json:"conversions,omitempty"`
	// DropOff 在这一秒内停止观看该视频的观众数
	DropOff []int64 `json:"drop_off,omitempty"`
	// Retain 在这一秒内继续观看该视频的观众数
	Retain []int64 `json:"retain,omitempty"`
	// CTR 在这一秒内视频观看产生了点击的次数占比
	CTR []float64 `json:"ctr,omitempty"`
	// CVR 在这一秒内视频点击产生了转化的次数占比
	CVR []float64 `json:"cvr,omitempty"`
}
