package report

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// IntegratedGetRequest 请求同步报表 API Request
type IntegratedGetRequest struct {
	// AdvertiserID 您需传入 advertiser_id，advertiser_ids和 bc_id 其中之一。
	// 若 report_type 设置为 BASIC 或 AUDIENCE ，需传入 advertiser_id 或advertiser_ids其中一个。若同时传入advertiser_id 和advertiser_ids，advertiser_id会被忽略。
	// 若 report_type 设置为 PLAYABLE_MATERIAL，CATALOG 或 TT_SHOP， 需传入 advertiser_id。
	// 若 report_type 设置为 BC，需传入 bc_id 。
	// 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserIDs 当report_type设置为 BASIC 或 AUDIENCE 时，需传入advertiser_id或advertiser_ids其中一个。如果您同时传入advertiser_id和advertiser_ids，advertiser_id会被忽略。
	// 仅在report_type为BASIC或AUDIENCE时生效。
	//
	// 广告主ID列表。
	// 最大数量：5。
	//
	// 当使用此字段时，请注意以下信息：
	// 当使用advertiser_ids时，将返回最多20,000个ID。ID的截取逻辑基于广告系列、广告组或广告的创建时间，按降序排列，但无法保证每个广告主都有均匀分布的ID数量。
	// 当使用advertiser_ids时，返回的metrics数据将包括currency，timezone，和advertiser_id。
	// 不同的广告主可能有不同的时区。可以使用multi_adv_report_in_utc_time来决定将所有广告主返回的数据设为 UTC+0 时区，或设为各个广告主当地时区。
	// 不同的广告主可能有不同的货币，我们将在返回的metrics中为每个广告主返回对应的货币。这些指标数据没有小数点。例如，“100美元”的花费将显示为100，没有小数点符号，如“.000”。
	// 如果您想查看所有广告主ID的汇总指标，请在请求中开启enable_total_metrics字段。返回中的total_metrics将包括currency，timezone和 advertiser_id，值都为-。
	//
	// 注意：
	//
	// 若使用advertiser_ids，但dimensions字段中不包括任何ID维度，则返回的metrics中currency，timezone和 advertiser_id的值将返回0，与成本相关的指标（例如spend）的将返回-。
	// advertiser_ids中指定的广告主应属于同一个TikTok for Business用户。否则会出错。使用/oauth2/advertiser/get/接口来获取同一个TikTok for Business用户的广告主ID列表。
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
	// BcID 您需传入 advertiser_id/advertiser_ids 和 bc_id 其中之一。
	// 若 report_type 设置为 BASIC 或 AUDIENCE ，需传入 advertiser_id或advertiser_ids 。
	// 若 report_type 设置为 PLAYABLE_MATERIAL，CATALOG 或 TT_SHOP， 需传入 advertiser_id。
	// 若 report_type 设置为 BC，需传入 bc_id 。
	//
	// 您有权访问的商务中心的 ID。
	//
	// 若想获取您有权访问的商务中心的 ID 列表，需使用/bc/get/。
	BcID string `json:"bc_id,omitempty"`
	// ServiceType report_type 设置为 BC 时不支持本字段。
	// 广告服务类型。
	// 枚举值：
	// AUCTION：竞价广告，或竞价和合约广告。
	// RESERVATION（已废弃）：合约广告。
	// report_type 未设置为 BC 时的默认值: AUCTION。
	// 注意：若想拉取 TopView 广告的报表数据，可使用基础报表或受众分析报表，并在筛选条件 buying_type 中指定 RESERVATION_TOP_VIEW 。
	ServiceType enum.ServiceType `json:"service_type,omitempty"`
	// ReportType 报表类型。
	// 枚举值：
	// BASIC：基础报表。
	// AUDIENCE：受众分析报表。
	// PLAYABLE_MATERIAL ：试玩素材报表。
	// CATALOG：DSA 报表。
	// BC ：商务中心报表。
	// TT_SHOP：最大成交额广告报表。
	//
	// 当 service_type 为 AUCTION时，支持BASIC，AUDIENCE，PLAYABLE_MATERIAL，CATALOG和 TT_SHOP 报表。
	// 当 report_type 为 BC 时，无需传入 service_type。
	ReportType enum.ReportType `json:"report_type,omitempty"`
	// DataLevel report_type 设置为 BC 时不支持本字段。
	// 当report_type为 BASIC、AUDIENCE 或 CATALOG 时必填。
	//
	// 您想要查询的报表数据层级。
	//
	// 枚举值：
	// AUCTION_AD：竞价广告或竞价和合约广告，广告层级。
	// AUCTION_ADGROUP：竞价广告或竞价和合约广告，广告组层级。
	// AUCTION_CAMPAIGN：竞价广告或竞价和合约广告，推广系列层级。
	// AUCTION_ADVERTISER：竞价广告或竞价和合约广告，广告主层级。
	// RESERVATION_AD（已废弃）：合约广告，广告层级。
	// RESERVATION_ADGROUP（已废弃）：合约广告，广告组层级。
	// RESERVATION_CAMPAIGN（已废弃）：合约广告，推广系列层级。
	// RESERVATION_ADVERTISER（已废弃）：合约广告，广告主层级。
	//
	// 注意：
	//
	// 若您将本字段设置为AUCTION_ADVERTISER，则会同时返回广告主账号下的所有竞价广告和合约广告数据。
	// 当report_type为BASIC，AUDIENCE或CATALOG时：
	// 若本字段设置为 AUCTION_CAMPAIGN，会返回非删除状态的竞价推广系列下所有广告的数据，因为默认应用了campaign_status为STATUS_NOT_DELETE的筛选条件。
	// 若本字段设置为AUCTION_ADGROUP，会返回非删除状态的竞价广告组下所有广告的数据，因为默认应用了adgroup_status为STATUS_NOT_DELETE的筛选条件。
	// 若本字段设置为 AUCTION_AD，会返回所有竞价推广系列下非删除状态的广告的数据，因为默认应用了ad_status为STATUS_NOT_DELETE的筛选条件。
	// 若您想获取所有竞价推广系列下所有广告的数据，需将筛选条件 ad_status设置为STATUS_ALL （filtering=[{"field_name":"ad_status","filter_type":"IN","filter_value":"[\"STATUS_ALL\"]"}] ）。
	DataLevel enum.ReportDataLevel `json:"data_level,omitempty"`
	// Dimensions 维度组合。示例：["campaign_id", "stat_time_day"] 表示组合campaign_id和stat_time_day维度。
	// 不同报表类型支持的维度不同。
	// 查看基础报表中支持的维度。
	// 查看受众分析报表中支持的维度。
	// 查看试玩素材报表中支持的维度。
	// 查看 DSA 报表中支持的维度。
	// 查看商务中心报表中支持的维度。
	// 查看最大成交额广告报表中支持的维度。
	//
	// 注意：
	// 若 report_type 未设置为BC：
	// 当dimensions包含 stat_time_day 时，查询时间范围不能超过 30 天，
	// 当dimensions包含 stat_time_hour 时，查询时间范围不能超过 1 天。
	// 若您查询的dimensions中不只包含 xx_id 且查询时间范围内没有广告报表，会返回空列表。
	Dimensions []string `json:"dimensions,omitempty"`
	// Metrics 要查询的指标。
	// 不同报表类型支持的指标不同。欲了解不同类型报表类型支持的指标，请参考报表类型下的相应文章。
	// 若report_type设置为 BASIC，AUDIENCE，PLAYABLE_MATERIAL 或 CATALOG，默认值为["spend", "impressions"]。
	// 若report_type设置为BC，则默认值为通过dimensions指定的维度组合下支持的所有指标。
	// 若report_type设置为TT_SHOP，默认值为["spend", "billed_cost"]。
	// 注意：与金额有关的货币数据指标会基于广告账户所设定的币种显示。
	Metrics []string `json:"metrics,omitempty"`
	// EnableTotalMetrics 是否开启所请求指标的汇总数据。
	// 当启用 enable_total_metrics 时，我们将在您查询不同页面时提供所有页面的汇总数据。在这种情况下，您只需在请求第一页数据时指定此字段。
	// 由于20,000 ID截断限制，启用 enable_total_metrics 时，我们将最多提供20,000个ID的汇总数据。
	// 注意:
	// 本字段仅支持竞价广告基础报表和最大成交额广告报表。
	// 当在dimensions中只指定了一个 ID 维度和(或)stat_time_day参数时，本字段才生效。例如：["campaign_id"], ["campaign_id", "stat_time_day"]。
	// 不能在沙箱环境中使用本字段。
	EnableTotalMetrics bool `json:"enable_total_metrics,omitempty"`
	// StartDate 查询起始日期（闭区间），格式为YYYY-MM-DD。日期根据广告主时区设定。
	// 若query_lifetime 未传入或设置为false，则本字段必填。
	// 注意：
	// 若report_type未设置为BC且dimensions的值包含 stat_time_day，则允许通过start_date 和 end_date 指定的最长查询范围为 30 天。
	// 若report_type未设置为BC且dimensions的值不包含 stat_time_day，则允许通过start_date 和 end_date 指定的最长查询范围为 365 天。
	// 若report_type设置为BC，则允许通过start_date 和 end_date 指定的最长查询范围为 365 天。
	StartDate string `json:"start_date,omitempty"`
	// EndDate 查询结束日期（闭区间），格式为YYYY-MM-DD。日期根据广告主时区设定。
	// 若query_lifetime 未传入或设置为false，则本字段必填。
	// 注意：
	// 若report_type未设置为BC且dimensions的值包含 stat_time_day，则允许通过start_date 和 end_date 指定的最长查询范围为 30 天。
	// 若report_type未设置为BC且dimensions的值不包含 stat_time_day，则允许通过start_date 和 end_date 指定的最长查询范围为 365 天。
	// 若report_type设置为BC，则允许通过start_date 和 end_date 指定的最长查询范围为 365 天。
	EndDate string `json:"end_date,omitempty"`
	// QueryLifeTime 仅当 report_type 设置为 BASIC 或 PLAYABLE_MATERIAL 时支持本字段。
	// 是否请求 Lifetime 指标。
	// 默认值：false。
	// 若 query_lifetime = true，将忽略 start_date 和 end_date 参数。Lifetime 指标名称和普通指标一致。
	// 注意：
	// 请求 Lifetime 指标时不支持时间细分，或受众维度细分。
	// 请求 Lifetime 指标时，不支持按地域代码 (country_code)细分。
	// 受众分析报表、DSA 报表、商务中心报表和最大成交额广告报表不支持 Lifetime 指标。
	QueryLifeTime bool `json:"query_lifetime,omitempty"`
	// MultiAdvReportInUTCTime 当请求中使用了advertiser_ids时可传入该字段。是否将返回的指标设置为各个广告主的当地时区。默认值：false。如果设置为true，则将所有广告主返回的指标设置为UTC+0时区。
	MultiAdvReportInUTCTime bool `json:"multi_adv_report_in_utc_time,omitempty"`
	// OrderField 排序字段。
	// t 所有支持的指标（不包括属性指标）均支持排序，默认不排序。
	OrderField string `json:"order_fieeld,omitempty"`
	// OrderType 排序方式。
	// 枚举值: ASC（升序），DESC（降序）。
	// 默认值: DESC。
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Filtering 仅当 report_type 设置为 BASIC，AUDIENCE，PLAYABLE_MATERIAL，CATALOG 或 BC 时支持本字段。
	// report_type 设置为 TT_SHOP 时不支持本字段。
	// 筛选条件。
	// 支持的筛选条件根据 service_type 和 data_level 有所不同。了解不同类型报表支持的筛选条件，请参考报表类型下的相应文章。
	// 注意:
	// 当您传入filtering字段时, 请确保同时传入 field_name, filter_type和filter_value字段，作为列表中的对象。
	// 若您想使用country_code筛选分组，只需要在请求的dimensions字段中指定country_code。country_code支持基础报表（同步/异步）、受众分析报表（同步/异步）、试玩素材报表（仅同步）和商务中心报表（仅同步）。
	Filtering []IntegratedGetFilter `json:"filtering,omitempty"`
	// QueryMode report_type 设置为 BC 时不支持本字段。
	// 数据请求模式。
	// 枚举值：REGULAR，CHUNK。
	// 默认值：REGULAR。
	// 开启CHUNK模式时，获取报表数据的响应时间会缩短，并且请求更加稳定，同时不支持分页。
	// 注意：
	// CHUNK模式即将废弃。为保证流畅的 API 体验，不推荐使用该模式。
	// CHUNK模式仅支持在同步基础报表中使用。
	QueryMode enum.ReportQueryMode `json:"query_mode,omitempty"`
	// Page 当前页数。
	// 默认值: 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 取值范围：1-1000。
	// 默认值: 10。
	PageSize int `json:"page_size,omitempty"`
}

// IntegratedGetFilter 筛选条件
type IntegratedGetFilter struct {
	// FieldName 筛选字段名称。
	FieldName string `json:"field_name,omitempty"`
	// FilterType 筛选类型。
	// 枚举值：
	// IN：包含，当筛选类型为此项时，筛选的值必须是合法的JSON数组字符串。
	// MATCH：模糊匹配，相当于 MySQL 中的LIKE 操作。
	// GREATER_EQUAL：大于等于（简称GE）。
	// GREATER_THAN：大于（简称GT）。
	// LOWER_EQUAL：小于等于（简称LE）。
	// LOWER_THAN：小于（简称LT）。
	// BETWEEN： 在……之间。当筛选类型为此项时，筛选的值需要是合法的、包含两个元素的JSON数组字符串。
	FilterType enum.FilterType `json:"filter_type,omitempty"`
	// FilterValue 要筛选的值。
	// 当 filter_type为IN 时，filter_value 必须是合法的 JSON 数组字符串。
	FilterValue string `json:"filter_value,omitempty"`
}

// Encode implements GetRequest interface
func (r *IntegratedGetRequest) Encode() string {
	values := util.NewURLValues()
	if v := r.AdvertiserID; v != "" {
		values.Set("advertiser_id", v)
	}
	if len(r.AdvertiserIDs) > 0 {
		values.Set("advertiser_ids", string(util.JSONMarshal(r.AdvertiserIDs)))
	}
	if v := r.BcID; v != "" {
		values.Set("bc_id", v)
	}
	if v := string(r.ServiceType); v != "" {
		values.Set("service_type", v)
	}
	if v := string(r.ReportType); v != "" {
		values.Set("report_type", v)
	}
	if v := string(r.DataLevel); v != "" {
		values.Set("data_level", v)
	}
	if len(r.Dimensions) > 0 {
		values.Set("dimensions", string(util.JSONMarshal(r.Dimensions)))
	}
	if len(r.Metrics) > 0 {
		values.Set("metrics", string(util.JSONMarshal(r.Metrics)))
	}
	if r.EnableTotalMetrics {
		values.Set("enable_total_metrics", "true")
	}
	if v := r.StartDate; v != "" {
		values.Set("start_date", v)
	}
	if v := r.EndDate; v != "" {
		values.Set("end_date", v)
	}
	if r.QueryLifeTime {
		values.Set("query_lifetime", "true")
	}
	if r.MultiAdvReportInUTCTime {
		values.Set("multi_adv_report_in_utc_time", "true")
	}
	if v := r.OrderField; v != "" {
		values.Set("order_field", v)
	}
	if v := string(r.OrderType); v != "" {
		values.Set("order_type", v)
	}
	if len(r.Filtering) > 0 {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if v := string(r.QueryMode); v != "" {
		values.Set("query_mode", v)
	}
	if v := r.Page; v > 0 {
		values.Set("page", strconv.Itoa(v))
	}
	if v := r.PageSize; v > 0 {
		values.Set("page_size", strconv.Itoa(v))
	}

	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// IntegratedGetResponse 请求同步报表 API Response
type IntegratedGetResponse struct {
	model.BaseResponse
	Data *IntegratedGetResult `json:"data,omitempty"`
}

type IntegratedGetResult struct {
	// TotalMetrics 所有请求指标的汇总数据。 当您在请求中开启 enable_total_metrics 参数时，返回此对象。
	TotalMetrics *Metrics `json:"total_metrics,omitempty"`
	// List 数据列表
	List []Stat `json:"report,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type Stat struct {
	// Dimensions 所有请求的维度组合数据
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	// Metrics 所有请求的指标数据
	Metrics *Metrics `json:"metrics,omitempty"`
}
