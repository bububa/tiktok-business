package creative

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/report"
	"github.com/bububa/tiktok-business/util"
)

// ReportGetRequest 创意基础/视频洞察报表
type ReportGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserIDs 需传入 advertiser_id 和 advertiser_ids 其中之一。
	// 广告主 ID 列表，用于为列表中的所有广告主生成报表。
	// 最大数量：50。
	// 注意：
	// 您需对列表对应的所有广告账号均有访问权限。若想获取某一商务中心中您有访问权限的广告账号，可使用/bc/asset/get/。
	// 若广告账号的币种不同，则返回的指标默认币种为美元。
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
	// ReportType 报表类型。
	// 枚举值：
	// VIDEO_INSIGHT：为一个或多个广告账号生成的视频洞察报表。
	ReportType enum.ReportType `json:"report_type,omitempty"`
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
	// TikTokItemIDs TikTok Spark Ads 帖子的 ID 列表，用于筛选用于创建这些帖子的视频。
	// 最大数量：20。
	// 若想获取 TikTok 帖子 ID，可使用 /tt_video/info/ 或 /identity/video/get/。
	TikTokItemIDs []string `json:"tiktok_item_ids,omitempty"`
	// AdName 广告名称。支持模糊搜索
	AdName string `json:"ad_name,omitempty"`
	// AdID 您希望用作筛选条件的广告ID列表。支持批量请求
	AdID []string `json:"ad_id,omitempty"`
	// AdgroupName 您希望用作筛选条件的广告组名称。支持模糊搜索
	AdgroupName string `json:"adgroup_name,omitempty"`
	// AdgroupID 您希望用作筛选条件的广告组ID。支持批量请求
	AdgroupID []string `json:"adgroup_id,omitempty"`
	// CampaignName 推广系列名称
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignID 您希望用作筛选条件的推广系列ID。支持批量请求
	CampaignID []string `json:"campaign_id,omitempty"`
	// CountryCode 国家或地区代码。枚举值见：附录-地区代码。
	CountryCode []string `json:"country_code,omitempty"`
	// CreativeFormat 素材样式列表。
	// 枚举值：
	// VIDEO_16_9：视频（16：9）。
	// VIDEO_1_1：视频（1：1）。
	// VIDEO_9_16：视频（9：16）。
	// OTHER：其他。
	CreativeFormat []string `json:"creative_format,omitempty"`
	// VideoMaterialSources 视频素材来源列表。
	// 枚举值：
	// UPLOADED_TO_TIKTOK_ADS_MANAGER：上传至 TikTok 广告管理平台的视频。
	// UPLOADED_TO_CATALOG：上传至商品库的视频。
	// CREATIVE_TOOL_SMART_VIDEO：创意工具 - 微电影。
	// CREATIVE_TOOL_QUICK_OPTIMIZATION：创意工具 - 一键优化。
	// CREATIVE_TOOL_VIDEO_TEMPLATE：创意工具 - 模版视频。
	// CREATIVE_TOOL_SMART_VIDEO_SOUNDTRACK：创意工具 - 智能配乐。
	// CREATIVE_TOOL_TIKTOK_VIDEO_EDITOR：创意工具 - 视频编辑器。
	// TIKTOK_CREATIVE_EXCHANGE：TikTok Creative Exchange。
	// CATALOG_VIDEO_TEMPLATE：商品库视频模版。
	// DYNAMIC_VIDEO_EDITOR：动态视频编辑器。
	// CREATIVE_CHALLENGE：创意挑战。
	// AUTOMATED_CREATIVE_OPTIMIZATION：程序化创意。
	// OTHER：其他。
	// QUICK_GENERATION：快速生成。
	// CREATIVE_CENTER_VIDEO_UPLOAD：创意中心 - 视频上传。
	// CREATIVE_CENTER_TIKTOK_VIDEO_EDITOR：创意中心 - TikTok 视频编辑器。
	// CREATIVE_CENTER_VIDEO_TEMPLATE：创意中心 - 视频模板。
	// DYNAMIC_SCENE：动态场景。
	// SMART_OPTIMIZATION_TOOL：智能优化工具。
	VideoMaterialSources []enum.VideoMaterialSource `json:"video_material_sources,omitempty"`
	// Placement 投放版位。枚举值见：枚举值-版位。
	Placement []enum.Placement `json:"placement,omitempty"`
	// Identities 认证身份（广告发布身份）列表。
	// 枚举值：
	// SPARK_ADS：可用于 Spark Ads 的认证身份，包括授权用户、TikTok 用户 和添加到商务中心的 TikTok 用户身份。
	// NON_SPARK_ADS：不可用于 Spark Ads 的认证身份，即自定义用户身份。
	// 若想了解认证身份详情，请查看认证身份。
	Identities []string `json:"identities,omitempty"`
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
	if r.AdvertiserID != "" {
		values.Set("advertiser_id", r.AdvertiserID)
	}
	if len(r.AdvertiserIDs) > 0 {
		values.Set("advertiser_ids", string(util.JSONMarshal(r.AdvertiserIDs)))
	}
	if r.ReportType != "" {
		values.Set("report_type", string(r.ReportType))
	}
	if r.MaterialType != "" {
		values.Set("material_type", string(r.MaterialType))
	}
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

// ReportGetResponse 创意基础/视频洞察报表 API Response
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
	Info *Dimensions `json:"info,omitempty"`
	// Metrics 指标。具体指标详情请参考指标表。
	Metrics *report.Metrics `json:"metrics,omitempty"`
}

type Dimensions struct {
	// MaterialID 素材 ID
	MaterialID string `json:"material_id,omitempty"`
	// MaterialName 视频素材名称
	MaterialName string `json:"material_name,omitempty"`
	// VideoID 视频 ID
	VideoID string `json:"video_id,omitempty"`
	// ImageID 图片 ID
	ImageID string `json:"image_id,omitempty"`
	// CreateTime 视频创建的时间，采用 ISO-8601 格式。
	// 示例： "2024-01-01T00:00:01Z"。
	CreateTime string `json:"create_time,omitempty"`
	// TikTokItemIDs 仅当视频（material_id）已用于创建 Spark Ads 时返回。
	// 使用该视频创建过的 TikTok Spark Ads 帖子的 ID 列表。
	TikTokItemIDs []string `json:"tiktok_item_ids,omitempty"`
	// CountryCode 视频（material_id）投放的地区的地区代码列表。
	// 若想了解枚举值，请查看附录 - 地区代码。
	CountryCode []string `json:"country_code,omitempty"`
	// VideoMaterialSource 视频素材来源。
	// 枚举值：
	// UPLOADED_TO_TIKTOK_ADS_MANAGER：上传至 TikTok 广告管理平台的视频。
	// UPLOADED_TO_CATALOG：上传至商品库的视频。
	// CREATIVE_TOOL_SMART_VIDEO：创意工具 - 微电影。
	// CREATIVE_TOOL_QUICK_OPTIMIZATION：创意工具 - 一键优化。
	// CREATIVE_TOOL_VIDEO_TEMPLATE：创意工具 - 模版视频。
	// CREATIVE_TOOL_SMART_VIDEO_SOUNDTRACK：创意工具 - 智能配乐。
	// CREATIVE_TOOL_TIKTOK_VIDEO_EDITOR：创意工具 - 视频编辑器。
	// TIKTOK_CREATIVE_EXCHANGE：TikTok Creative Exchange。
	// CATALOG_VIDEO_TEMPLATE：商品库视频模版。
	// DYNAMIC_VIDEO_EDITOR：动态视频编辑器。
	// CREATIVE_CHALLENGE：创意挑战。
	// AUTOMATED_CREATIVE_OPTIMIZATION：程序化创意。
	// OTHER：其他。
	// QUICK_GENERATION：快速生成。
	// CREATIVE_CENTER_VIDEO_UPLOAD：创意中心 - 视频上传。
	// CREATIVE_CENTER_TIKTOK_VIDEO_EDITOR：创意中心 - TikTok 视频编辑器。
	// CREATIVE_CENTER_VIDEO_TEMPLATE：创意中心 - 视频模板。
	// DYNAMIC_SCENE：动态场景。
	// SMART_OPTIMIZATION_TOOL：智能优化工具。
	//
	// 注意：若视频的素材来源无法追踪，本字段可能返回 null 值。
	VideoMaterialSource enum.VideoMaterialSource `json:"video_material_source,omitempty"`
	// Source 素材来源
	Source enum.MaterialSource `json:"source,omitempty"`
	// Placement 视频投放的广告版位。
	// 枚举值：
	// PLACEMENT_TIKTOK：TikTok。
	// PLACEMENT_PANGLE：Pangle。
	// PLACEMENT_GLOBAL_APP_BUNDLE：全球化应用组合。
	Placement []enum.Placement `json:"placement,omitempty"`
	// Identity 视频投放时使用的认证身份。
	// 枚举值：
	// SPARK_ADS：可用于 Spark Ads 的认证身份，包括授权用户、TikTok 用户 和添加到商务中心的 TikTok 用户身份。
	// NON_SPARK_ADS：不可用于 Spark Ads 的认证身份，即自定义用户身份。
	// 若想了解认证身份详情，请查看认证身份。
	Identity string `json:"identity,omitempty"`
	// Currency metrics 中的指标数据使用的币种代码。
	// 示例：USD。
	// 若想了解币种代码对应的币种，请查看不同币种的预算校验比例和取值范围。
	Currency string `json:"currency,omitempty"`
	// NumAds 绑定的广告数量
	NumAds int64 `json:"num_ads,omitempty"`
	// TotalActiveDays 累计活跃天数。若spend > 0，则当日计入活跃天。此字段仅支持VIDEO类型素材
	TotalActiveDays int64 `json:"total_active_days,omitempty"`
	// RelatedAdIDs 关联的广告。此字段仅支持VIDEO类型素材
	RelatedAdIDs []string `json:"related_ad_ids,omitempty"`
	// AdgroupNumber 关联的广告组数量。此字段仅支持VIDEO类型素材
	AdgroupNumber int64 `json:"adgroup_number,omitempty"`
	// RelatedAdgroupIDs 关联的广告组。此字段仅支持VIDEO类型素材
	RelatedAdgroupIDs []string `json:"related_adgroup_ids,omitempty"`
	// PageID 落地页ID列表
	PageID []string `json:"page_id,omitempty"`
	// PageStaus 落地页状态。枚举值: DRAFT（草稿）, READY（可用）
	PageStatus []enum.PageStatus `json:"page_status,omitempty"`
	// PageBizType 落地页业务类型。枚举值: CUSTOM（定制）, APP_PROFILE（应用信息）, INSTANT_FORM（即时表单）
	PageBizType []enum.PageBizType `json:"page_biz_type,omitempty"`
	// PageTemplateType 落地页模板类型。在page_biz_type 中包含 CUSTOM时，枚举值为：
	// CUSTOMIZED（定制）
	// INTRODUCTION（介绍）
	// BRAND_SAFETY（品牌安全）
	// SALES_PRODUCT（产品销售）
	// MOVIE_TRAILER（电影预告）
	// 在page_biztype 包含 INSTANT_FORM时，枚举值为NORMAL_FORM（常规表单）和ADVANCED_FORM`（复杂表单）
	PageTemplateType []enum.PageTemplateType `json:"page_template_type,omitempty"`
	// PageThumbnail 页面缩略图链接
	PageThumbnail string `json:"page_thumbnail,omitempty"`
	// AppProfileCountry 落地页配置的国家代码列表。只在instant_page_biz_type包含 APP_PROFILE时有效
	AppProfileCountry []string `json:"app_profile_country,omitempty"`
	// AppProfileExternalAppID 落地页导向的应用的ID。只在instant_page_biz_type包含 APP_PROFILE时有效
	AppProfileExternalAppID []string `json:"app_profile_exterinal_app_id,omitempty"`
	// AppProfileIcon 应用图标
	AppProfileIcon string `json:"app_profile_icon,omitempty"`
	// CustomizedPageExternalAction 定制落地页特有的转化事件
	CustomizedPageExternalAction string `json:"customized_page_external_action,omitempty"`
}
