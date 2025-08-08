package spc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/report"
	"github.com/bububa/tiktok-business/util"
)

// ReportGetRequest 生成 Smart+ 推广系列报表 API Request
type ReportGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignIDs 推广系列 ID 列表。
	// 最大数量：1。
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// MaterialType 推广系列的创意素材类型。
	// 枚举值：SPC_VIDEO，SPC_AD_TEXT，SPC_SPARK，SPC_CATALOG_VIDEO，SPC_CATALOG_CAROUSEL。
	// 若material_type为SPC_VIDEO，则可以获取对应视频的数据。您可在metrics字段中指定基础数据指标、视频播放数据指标、互动情况指标、应用内事件指标、应用内事件指标（SKAN）、网页事件指标、时间指标、日期指标和视频素材详情指标。
	// 若material_type为SPC_AD_TEXT，则可以获取广告文案的数据。您仅可在metrics字段中指定基础数据指标、互动情况指标、应用内事件指标、应用内事件指标（SKAN）、网页事件指标、时间指标、日期指标和视频素材详情指标。
	// 若material_type为SPC_SPARK，则可以获取 Spark Ads 帖子的数据。您仅可在metrics字段中指定基础数据指标、互动情况指标、应用内事件指标、应用内事件指标（SKAN）、网页事件指标和时间指标。
	// 若 material_type 为 SPC_CATALOG_VIDEO，则可以获取商品库视频的数据。您仅可在metrics字段中指定基础数据指标、互动情况指标、应用内事件指标、应用内事件指标（SKAN）、网页事件指 标和日期指标。
	// 若 material_type 为 SPC_CATALOG_CAROUSEL，则可以获取商品库轮播的数据。您仅可在metrics字段中指定基础数据指标、互动情况指标、应用内事件指标、应用内事件指标（SKAN）、网页事件指标和日期指标。
	//
	// 若想了解各个指标类别中包含的具体指标，可查看 Smart+ 推广系列报表指标。
	//
	// 注意：
	//
	// 当 material_type 设置为 SPC_SPARK时，确保 campaign_ids 字段仅包含 Spark Ads 的推广系列。如果在 campaign_ids 中同时包含 Spark Ads 和非 Spark Ads，则将返回空列表。
	// SPC_CATALOG_VIDEO 和 SPC_CATALOG_CAROUSEL 类型的 Smart+ 推广系列报表仅支持 Smart+ 商品库广告。
	MaterialType enum.MaterialType `json:"material_type,omitempty"`
	// IsIncludeRemovedMaterial 仅当 material_type 为 SPC_VIDEO，SPC_AD_TEXT 或 SPC_SPARK 时生效。
	// 是否查询已删除的视频、文案或帖子的数据。
	// 枚举值：
	// true：返回中包含已删除的视频、文案或帖子的数据。
	// false：返回中不包含已删除的视频、文案或帖子的数据。
	// 默认值：true
	IsIncludeRemovedMaterial bool `json:"is_include_removed_material,omitempty"`
	// Metrics 要查询的数据指标。
	// 默认值：["spend", "impressions"]。
	// 具体指标请参阅 Smart+ 推广系列报表支持的所有指标。
	Metrics []string `json:"metrics,omitempty"`
	// QueryLifetime 是否请求lifetime指标。
	// 默认值：false。
	// 若query_lifetime = true，则忽略start_date和end_date参数。
	QueryLifetime bool `json:"query_lifetime,omitempty"`
	// StartDate 当query_lifetime设置为false或不填时，该字段为必填。
	// 查询起始日期（闭区间），格式为YYYY-MM-DD（UTC+0）。日期根据广告主时区设定。
	StartDate string `json:"start_date,omitempty"`
	// EndDate 当query_lifetime设置为false或不填时，该字段为必填。
	// 查询结束时间（闭区间），格式为YYYY-MM-DD（UTC+0）。日期根据广告主时区设定。
	// 注意：start_date和end_date的间隔时间不能超过365天。
	EndDate string `json:"end_date,omitempty"`
	// Page 当前页数。
	// 默认值：1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 取值范围：1-100。
	// 默认值：10。
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r *ReportGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("campaign_id", string(util.JSONMarshal(r.CampaignIDs)))
	values.Set("material_type", string(r.MaterialType))
	if !r.IsIncludeRemovedMaterial {
		values.Set("is_include_removed_material", "false")
	}
	if len(r.Metrics) > 0 {
		values.Set("metrics", string(util.JSONMarshal(r.Metrics)))
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

// ReportGetResponse 生成 Smart+ 推广系列报表 API Response
type ReportGetResponse struct {
	model.BaseResponse
	Data *ReportGetResult `json:"data,omitempty"`
}

type ReportGetResult struct {
	// List 返回的数据列表
	Lisit []report.Stat `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
