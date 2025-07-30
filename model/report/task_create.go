package report

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TaskCreateRequest 创建异步报表任务 API Request
type TaskCreateRequest struct {
	// AdvertiserID advertiser_id 和advertiser_ids 必传其一。若您同时传入advertiser_id和advertiser_ids，将忽略advertiser_id字段。
	// 广告主ID。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserIDs advertiser_id 和advertiser_ids必传其一。若您同时传入advertiser_id和advertiser_ids，将忽略advertiser_id字段。
	// 广告主ID列表。
	// 最大数量：5。
	// 注意：
	// 本字段仅当 report_type为BASIC或AUDIENCE时生效。
	// advertiser_ids中指定的广告主应属于同一个TikTok for Business用户。否则会出错。使用/oauth2/advertiser/get/接口来获取同一个TikTok for Business用户的广告主ID列表。
	// 若您还同时传入了start_date 和 end_date ，会根据每个广告账户的时区拉取数据。
	// 若您传入多个广告主ID，您可在dimensions字段设置最多两个ID维度（advertiser_id和另一 ID 维度）。不过，很多指标不支持在advertiser_id 维度下使用。同时，推荐您在metrics的值中设置currency，因为不同广告账号的币种可能不同。
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
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
	// OrderField 排序字段。
	// t 所有支持的指标（不包括属性指标）均支持排序，默认不排序。
	OrderField string `json:"order_fieeld,omitempty"`
	// OrderType 排序方式。
	// 枚举值: ASC（升序），DESC（降序）。
	// 默认值: DESC。
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// EnableReportTitleTranslate 是否启用异步报表标题行（包括维度名称和指标名称）的翻译。
	// 若本字段设置为false，则将显示未翻译的标题（例如ad_id）而非翻译过的标题（例如AD ID）。
	// 默认值：true。
	// 注意：
	//
	// 本字段仅当report_type为BASIC或AUDIENCE时生效。
	// 自 2024 年 11 月 15 日起，将更新异步报表标题行中某些指标的翻译，以提供更好的用户体验。这一变更不会影响指标字段本身的名称。为了最大限度地减少潜在影响，强烈建议在您的代码中将 enable_report_title_translation 设置为 false。这将消除代码对指标翻译的依赖，避免指标翻译的定期更新带来的影响。虽然指标翻译可帮助您将指标映射到 TikTok 广告管理平台中的对应报表指标，但在代码中依赖指标翻译可能会在未来引发重大变更。通过禁用标题行翻译，您可以确保您的集成更加稳定、可靠。
	EnableReportTitleTranslate bool `json:"enable_report_title_translation,omitempty"`
	// OutputFormat 输出格式。 枚举值：CSV_STRING，CSV_DOWNLOAD，XLSX_DOWNLOAD。默认值：CSV_STRING。
	OutputFormat enum.OutputFormat `json:"output_format,omitempty"`
	// FileName 自定义下载文件名称。当output_format为CSV_DOWNLOAD或XLSX_DOWNLOAD时，可选择传入。最多可传入255个字符，文件名称不能包含以下特殊符号：['/', '\t', '\b', '@', '#', '$', '%', '^', '&', '*', '(', ')', '[', ']']。
	FileName string `json:"file_name,omitempty"`
	// Filtering 筛选条件。支持的筛选条件根据 service_type 和 data_level 有所不同。了解不同类型报表支持的筛选条件，请参考报表类型下的相应文章。
	// 注意:
	//
	// 当您传入filtering字段时, 请确保同时传入 field_name, filter_type和filter_value字段，作为列表中的对象。
	// 若您想使用country_code筛选分组，只需要在请求的dimensions字段中指定country_code。country_code支持基础报表（同步/异步）、受众分析报表（同步/异步）和试玩素材报表（仅同步）。
	// 异步模式下的基础报表，受众分析报表、DSA报表和合约广告报表仅支持campaign_ids，adgroup_ids和ad_ids为筛选条件。
	Filtering []IntegratedGetFilter `json:"filtering,omitempty"`
}

// Encode implements PostRequest interface
func (r *TaskCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TaskCreateResponse 创建异步报表任务 API Response
type TaskCreateResponse struct {
	model.BaseResponse
	Data struct {
		// TaskID 异步报表任务ID
		TaskID string `json:"task_id,omitempty"`
	} `json:"data"`
}
