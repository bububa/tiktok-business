package creative

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/report"
	"github.com/bububa/tiktok-business/util"
)

// ReportGetRequest 创意基础报表
type ReportGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// MaterialType 素材类型。 枚举值：VIDEO、 IMAGE, INSTANT_PAGE
	MaterialType enum.MaterialType `json:"material_type,omitempty"`
	// Lifetime 是否查询全部的数据。如果是，那么不需要指定 start_date 和 end_date。默认值：False。
	Lifetime bool `json:"lifetime,omitempty"`
	// StartDate 起始时间，闭区间。格式如：2020-01-01（广告主时区）。当lifetime为False时必填
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，闭区间。格式如：2020-01-01（广告主时区）。当lifetime为False时必填
	EndDate string `json:"end_date,omitempty"`
	// InfoFields 您需要的有关创意素材的信息。默认值： [material_id, video_id, image_id, page_id]。可选范围：返回值中info下的所有字段。
	InfoFields []string `json:"info_fields,omitempty"`
	// MetricsFields 您需要的指标或者维度数据。默认值：[impressions , spend]。可选范围：返回值中metrics下的所有字段。
	MetricsFields []string `json:"metrics_fields,omitempty"`
	// Filtering 筛选条件
	Filtering *ReportGetFilter `json:"filtering,omitempty"`
	// SortField 排序字段。支持按照素材创建时间以及所有的指标数据进行排序。默认不排序
	SortField string `json:"sort_field,omitempty"`
	// SortType 排序方式。枚举值：ASC, DESC。 默认值：DESC
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Page 当前页数。默认值：1，取值范围：≥ 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。默认值：10，取值范围：1-1000
	PageSize int `json:"page_size,omitempty"`
}

// ReportGetFilter 筛选条件
type ReportGetFilter struct {
	// MaterialID 素材ID。最多可传入10个ID
	MaterialID []string `json:"material_id,omitempty"`
	// MaterialName 素材名称。支持模糊搜索
	MaterialName string `json:"material_name,omitempty"`
	// AdName 广告名称。支持模糊搜索
	AdName string `json:"ad_name,omitempty"`
	// AdID 您希望用作筛选条件的广告ID列表。支持批量请求
	AdID []string `json:"ad_id,omitempty"`
	// AdgroupName 您希望用作筛选条件的广告组名称。支持模糊搜索
	AdgroupName string `json:"adgroup_name,omitempty"`
	// AdgroupID 您希望用作筛选条件的广告组ID。支持批量请求
	AdgroupID []string `json:"adgroup_id,omitempty"`
	// CampaignID 您希望用作筛选条件的推广系列ID。支持批量请求
	CampaignID []string `json:"campaign_id,omitempty"`
	// CountryCode 国家或地区代码。枚举值见：附录-地区代码。
	CountryCode []string `json:"country_code,omitempty"`
	// Placement 投放版位。枚举值见：枚举值-版位。
	Placement []enum.Placement `json:"placement,omitempty"`
	// PageID 落地页ID列表
	PageID []string `json:"page_id,omitempty"`
	// PageStatus 落地页状态。枚举: DRAFT（草稿）, READY（可用）
	PageStatus enum.PageStatus `json:"page_status,omitempty"`
	// PageBizType 落地页业务类型。枚举: CUSTOM（定制）, APP_PROFILE（应用信息）, INSTANT_FORM（即时表单）
	PageBizType enum.PageBizType `json:"page_biz_type,omitempty"`
	// PageTemplateType 落地页模板类型。在page_biz_type 中包含 CUSTOM时，枚举值为：
	// CUSTOMIZED（定制）
	// INTRODUCTION（介绍）
	// BRAND_SAFETY（品牌安全）
	// SALES_PRODUCT（产品销售）
	// MOVIE_TRAILER（电影预告）
	// 在page_biz_type 包含 INSTANT_FORM时，枚举值为NORMAL_FORM（常规表单），ADVANCED_FORM（复杂表单）
	PageTemplateType enum.PageTemplateType `json:"page_template_type,omitempty"`
	// AppProfileCountry 落地页配置的国家或地区代码列表。只在instant_page_biz_type包含 APP_PROFILE时有效
	AppProfileCountry []string `json:"app_profile_country,omitempty"`
	// AppProfileExternalAppID 落地页导向的应用的ID。只在instant_page_biz_type包含 APP_PROFILE时有效
	AppProfileExternalAppID []string `json:"app_profile_exterinal_app_id,omitempty"`
	// CustomizedPageExternalAction 定制落地页特有的转化事件
	CustomizedPageExternalAction []string `json:"customized_page_external_action,omitempty"`
	// Spend 消耗范围
	Spend *Range `json:"spend,omitempty"`
	// Conversion 转化值
	Conversion *Range `json:"conversion,omitempty"`
}

// Range 范围
type Range struct {
	// Min 最小值，开区间。不填表示不限制最小值
	Min float64 `json:"min,omitempty"`
	// Max 最大值，闭区间。不填表示不限制最大值
	Max float64 `json:"max,omitempty"`
}

// Encode impressions GetRequest interface
func (r *ReportGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("material_type", string(r.MaterialType))
	if r.Lifetime {
		values.Set("lifetime", "true")
	}
	if r.StartDate != "" {
		values.Set("start_date", r.StartDate)
	}
	if r.EndDate != "" {
		values.Set("end_date", r.EndDate)
	}
	if len(r.InfoFields) > 0 {
		values.Set("info_fields", string(util.JSONMarshal(r.InfoFields)))
	}
	if len(r.MetricsFields) > 0 {
		values.Set("metrics_fields", string(util.JSONMarshal(r.MetricsFields)))
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

// ReportGetResponse 创意基础报表 API Response
type ReportGetResponse struct {
	model.BaseResponse
	Data *ReportGetResult `json:"data,omitempty"`
}

type ReportGetResult struct {
	// List 分页信息
	List []Stat `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type Stat struct {
	// Info 素材信息
	Info *report.Dimensions `json:"info,omitempty"`
	// Metrics 指标。具体指标详情请参考指标表。
	Metrics *report.Metrics `json:"metrics,omitempty"`
}
