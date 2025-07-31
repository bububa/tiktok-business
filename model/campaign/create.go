package campaign

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建推广系列 API Request
type CreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// ObjectiveType 推广目标，获取枚举值，参阅 枚举值-推广目标。
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// AppPromotionType 应用推广类型。枚举值：
	// APP_INSTALL：应用安装。让新用户安装您的应用。
	// APP_RETARGETING：应用再营销。重新吸引现有用户在您的应用中执行操作。
	// APP_PREREGISTRATION：应用预注册。让新用户在您的应用发布前完成预注册。
	// APP_POSTS_PROMOTION：应用帖子推广。推广 TikTok 帖子，提升应用认知度并度量转化量。
	// 注意：目前不支持通过 API 创建应用帖子推广类型的推广系列，仅支持通过 API 更新或检索此类推广系列。
	AppPromotionType enum.AppPromotionType `json:"aepp_promotion_type,omitempty"`
	// VirtualObjectiveType 新推广目标类型。

	// 枚举值：SALES（销量）。
	//
	// 使用本字段创建推广目标为销量的推广系列，该目标合并了网站转化量目标和商品销量目标。了解销量目标的详情，参见销量目标。
	VirtetualObjectiveType enum.VirtualObjectiveType `json:"virtual_objective_type,omitempty"`
	// SalesDestination 销售目标页面，即想要推动销售的渠道。
	// 枚举值：
	// TIKTOK_SHOP：TikTok Shop。推动 TikTok Shop 上的销售。
	// WEBSITE：网站。推动网站上的销售。
	// APP：应用。推动应用上的销售（需要商品库）。
	SalesDestination enum.SalesDestination `json:"sales_destination,omitempty"`
	// CampaignType 推广系列类型。枚举值: REGULAR_CAMPAIGN（普通推广系列）、IOS14_CAMPAIGN（iOS 14专属推广系列）。
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// AppID 推广的App的ID。您可通过/app/list/获取app_id。
	AppID string `json:"app_id,omitempty"`
	// IsAdvancedDedicatedCampaign 该推广系列是否为高级专属推广系列。高级专属推广系列采用捕捉实时信号的高级投放模型。
	// 枚举值：true，false。
	IsAdvancedDedicatedCampaign bool `json:"is_advanced_dedicated_campaign,omitempty"`
	// DisableSkanCampaign 是否关闭 SKAN（SKAdNetwork）归因（Apple 针对 iOS 推广系列的转化归因解决方案）。
	// 枚举值：
	// true：关闭 SKAN 归因。推广系列将不受专属推广系列配额限制的约束，您可为此推广系列拉取自归因平台（SAN）指标数据。但是，不支持为此推广系列拉取 SKAN 指标数据。了解自归因平台集成的详情。
	// false：启用 SKAN 归因。推广系列将受专属推广系列配额限制的约束，您可为此推广系列拉取 SKAN 指标数据。
	DisableSkanCampaign *bool `json:"disable_skan_campaign,omitempty"`
	// CampaignAppProfilePageState 下载中间页使用情况。
	// 枚举值： ON , OFF。
	// 默认值：OFF。
	// 本字段仅可在objective_type为APP_PROMOTION的 iOS 14专属推广系列中使用，否则会出现报错。
	CampaignAppProfilePageState enum.AppProfilePageState `json:"campaign_app_profile_page_state,omitempty"`
	// RfCampaignType 仅当推广目标（objective_type）设置为RF_REACH时，可根据本字段判断具体的合约推广系列类型。
	// 推广目标非RF_REACH时，本字段无需传入。

	// 枚举值： STANDARD （普通覆盖和频次推广系列）， PULSE（TikTok Pulse推广系列）。
	//
	// 注意:
	//
	// 本字段目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	// 推广系列创建后，本字段不可修改。
	// 若您将rf_campaign_type 设置为PULSE，则使用/adgroup/rf/create/创建广告组时，需在contextual_tag_ids字段传入 PREMIUM类型的内容相关定向标签，创建的广告组的CPM将是固定的。
	// 若您将rf_campaign_type 设置为PULSE，则使用/adgroup/rf/update/更新广告组时，不支持更新定向地域。
	// 若您将rf_campaign_type 设置为PULSE，则使用/adgroup/rf/create/创建广告组时，feed_type不可设置为 TOP_FEED。
	RfCampaignType enum.RfCampaignType `json:"rf_campaign_type,omitempty"`
	// CampaignProductSource 仅当推广系列对应的objective_type为PRODUCT_SALES，且设置了商品来源时返回。
	// 推广系列的商品来源。
	// 枚举值：
	// CATALOG ：商品库。
	// STORE：TikTok Shop 店铺，或 TikTok 橱窗。
	CampaignProductSource enum.ProductSource `json:"campaign_product_source,omitempty"`
	// CampaignName 推广系列名称
	// 推广系列名称不能超过512字符，不能包含emoji。
	// 中文或日文每个字占用2个字符，每个英语字符占用1个字符。
	CampaignName string `json:"campaign_name,omitempty"`
	// RequestID 请求ID，用于创建同名推广计划。该ID支持接口幂等，避免重复请求。如果您传入相同的request ID重试多次请求，只有一次请求会成功。
	// 该ID和返回参数中的 request_id 不同。返回的request_id用于唯一标识一次 HTTP 请求。
	//
	// 本字段的值需为字符串格式的64位整数值。
	RequestID string `json:"request_id,omitempty"`
	// SpecialIndustries 特殊广告分类。
	//为推广系列选择住房、就业或信贷分类，即表示您保证不会利用 TikTok 在与住房、就业或信贷相关的广告中歧视具有受保护特征的人群，且遵守 TikTok 的特殊广告分类政策。
	//
	// 您使用字段表明您的推广系列推广特殊广告分类的机会后，广告组层级的定位选项将受限，从而帮助您遵守这些类别相关的反歧视法。若想了解相关限制，请查看使用特殊广告分类的推广系列中广告组的定向要求。
	//
	// 枚举值:
	// HOUSING：房地产，房屋保险，抵押贷款或相关的广告。
	// EMPLOYMENT：工作机会，实习机会，职业认证项目或相关的广告。
	// CREDIT：信用卡申请，汽车贷款，长期融资或相关的广告。
	//
	// 注意：
	//
	// 在推广系列商品来源为商品库和推广系列商品来源非商品库的购物广告中使用特殊广告分类目前为单独的白名单功能。如需使用这些功能，请联系您的TikTok销售代表。
	// 本字段对于注册地为美国或加拿大的广告主已全量。若您是注册地非美国或加拿大的广告主且想在定向美国或加拿大的广告中使用特殊广告分类，您需申请额外的白名单。
	// 您为推广系列设置特殊广告分类后，允许清除特殊广告分类，但不支持更改特殊广告分类的类别，且若创建时推广系列未设置特殊广告分类，不允许通过更新为推广系列重新启用特殊广告分类。
	SpecialIndustries []enum.SpecialIndustry `json:"special_industries,omitempty"`
	// BudgetOptimizeOn 仅为开启了推广系列预算优化的推广系列返回本字段。
	// 是否开启推广系列预算优化。
	// 枚举值：true（开启）。
	// 详见推广系列预算优化。
	// 注意：对于 Smart+ 推广系列或智能效果网站推广系列，本字段值将为 true，因为这些推广系列均默认启用推广系列预算优化。
	BudgetOptimizeOn bool `json:"budget_optimize_on,omitempty"`
	// BudgetMode 预算类型。
	// 枚举值：
	// BUDGET_MODE_INFINITE：不限预算。
	// BUDGET_MODE_TOTAL：总预算。
	// BUDGET_MODE_DAY：日预算。
	// BUDGET_MODE_DYNAMIC_DAILY_BUDGET：动态日预算，即一周七日的平均预算。单日实际消耗不超过设定的七日平均预算的125%。周消耗不会超过七日平均预算*7。
	// 当objective_type为非RF_REACH时，budget_mode必填。
	// 当objective_type为RF_REACH时，budget_mode仅支持BUDGET_MODE_INFINITE。
	//
	// 预算类型。
	//
	// 开启推广系列预算优化时，推广系列层级的budget_mode支持以下枚举值：
	// BUDGET_MODE_TOTAL：总预算。
	// BUDGET_MODE_DYNAMIC_DAILY_BUDGET：动态日预算，即一周七日的平均预算。单日实际消耗不超过设定的七日平均预算的125%。周消耗不会超过七日平均预算*7。
	// objective_type设置为TRAFFIC，APP_PROMOTION，WEB_CONVERSIONS，LEAD_GENERATION，PRODUCT_SALES（仅视频购物广告），REACH，VIDEO_VIEWS 或 ENGAGEMENT且您不想设置总预算时，将本字段设置为BUDGET_MODE_DYNAMIC_DAILY_BUDGET。
	// 注意：为推广目标为REACH，VIDEO_VIEWS 或 ENGAGEMENT 的推广系列启用 BUDGET_MODE_DYNAMIC_DAILY_BUDGET 目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	//
	// BUDGET_MODE_DAY：日预算。
	// objective_type 设置为REACH，VIDEO_VIEWS 或 ENGAGEMENT时，推荐您使用动态日预算（BUDGET_MODE_DYNAMIC_DAILY_BUDGET）而非日预算（BUDGET_MODE_DAY）。
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 推广系列预算
	// 当 budget_mode 为 BUDGET_MODE_DAY 或 BUDGET_MODE_DYNAMIC_DAILY_BUDGET或BUDGET_MODE_TOTAL 时必填。
	// 欲了解最低预算以及如何设置预算类型，请参阅预算。欲直接获取某一币种对应的每日预算取值范围，请参阅币种-每日预算取值范围。
	Budget float64 `json:"budget,omitempty"`
	// RtaID 实时 API ID，即实时 API 策略标识符
	// 若想获取与您的广告账户绑定的实时 API ID，请联系您的 TikTok 销售代表。
	// 若想了解为广告启用实时 API 的前提及启用步骤，请查看实时 API。
	//
	// 注意：为广告启用实时 API 目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。您需首先完成实时 API 的入门流程，具备接收发送至您的系统的实时 API 请求并进行回复的能力。
	RtaID string `json:"rta_id,omitempty"`
	// RtaBidEnabled 是否使用实时 API 出价。
	// 支持的值：
	// true：启用。
	// false：不启用。
	// 默认值：false。
	//
	// 仅在同时满足以下条件时支持将本字段设置为 true：
	// objective_type 为 APP_PROMOTION、WEB_CONVERSIONS 或 PRODUCT_SALES。
	// 传入了 rta_id。
	//
	// 注意: 启用实时 API 出价目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	RtaBidEnabled bool `json:"rta_bid_enabled,omitempty"`
	// RtaProductSelectionEnabled 是否使用实时 API 自动选择商品。
	// 枚举值：
	// true：使用。
	// false：不使用。
	// 默认值： false。
	//
	// 仅在满足以下条件时支持将本字段设置为true：
	// objective_type 设置为 PRODUCT_SALES。
	// campaign_product_source 设置为 CATALOG。
	// 传入了 rta_id 。
	RtaProductSelectionEnabled bool `json:"rta_product_selection_enabled,omitempty"`
	// OperationStatus 推广系列的操作状态。
	// ENABLE：推广系列处于启用（“开”）状态。
	// DISABLE：推广系列处于未启用（“关”）状态。
	// 对于覆盖和频次推广系列，需将operation_status设置为ENABLE，或不传入operation_status，请不要将operation_status设置为DISABLE。
	// 默认值：ENABLE。
	// 若想在推广系列创建后更新其启用/关闭状态，请使用/campaign/status/update/接口。
	OpeterationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// PostbackWindowMode 选定的模式决定您想要确保能接收到的 SKAN (SKAdNetwork) 4.0 回传。
	// 本字段支持在创建暂停状态的推广系列时指定，或在通过/campaign/status/update/接口暂停现有推广系列时指定。推荐做法为在暂停现有推广系列时指定本字段。
	// 选定的模式决定您想要确保能接收到的 SKAN (SKAdNetwork) 4.0 回传。对应的窗口期较长的回传需要更多时间接收回传，推广系列配额释放的时间也相应较长。
	//
	// 枚举值：
	// POSTBACK_WINDOW_MODE1：此模式确保能接收到第一次回传，对应的归因窗口期为 0-2 天。回传数据最长需要 4 天时间接收，推广系列配额释放的时间最长也为 4 天。
	// POSTBACK_WINDOW_MODE2：此模式确保能接收到第一次及第二次回传，对应的归因窗口期为 3-7 天。回传数据最长需要 13 天时间接收，推广系列配额释放的时间最长也为 13 天。
	// POSTBACK_WINDOW_MODE3：此模式确保能接收到三次回传，对应的归因窗口期为 8-35 天。回传数据最长需要 41 天时间接收，推广系列配额释放的时间最长也为 41 天。
	//
	// 若想了解 SKAN 4.0 归因窗口期和推广系列配额详情，请查看专属推广系列。
	//
	// 注意：
	//
	// 若您已为应用启用 Adjust，Airbridge，Appsflyer，Branch，Kochava，或 Singular 移动合作伙伴 (MMP) 跟踪，且您的 MMP SDK 版本已更新至支持 SKAN 4.0 的版本，您可在事件管理平台将应用迁移至 SKAN 4.0 。若想了解为应用启用 MMP 跟踪的步骤，请查看如何在 TikTok 广告管理平台中设置应用归因。若想了解将应用迁移至 SKAN 4.0 的步骤，请查看关于 SKAN 4.0 和 TikTok 以及如何迁移到 SKAN 4.0。
	// 本字段设置后不允许修改。
	// 若operation_status设置为ENABLE或未传入，将忽略本字段。
	// 若您传入了本字段，但指定的应用（app_id）未启用 SKAN 4.0 ，将出现报错。
	// 若campaign_type 设置为IOS14_CAMPAIGN，operation_status 设置为 DISABLE，且您已为应用启用了 SKAN 4.0，但您未传入本字段，则本字段值默认设置为POSTBACK_WINDOW_MODE1。
	// 若您已为应用启用了 SKAN 4.0，需定向 iOS 16.1 及以上版本的设备，从而确保能够接收 SKAN 4.0 的回传。若您想仅定向 iOS 16.1 及以上版本的设备，可在广告组层级将 min_ios_version 设置为 16.1 。
	PostbackWindowMode enum.PostbackWindowMode `json:"postback_window_mode,omitempty"`
}

// Encode implements PostRequest interface
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建推广系列 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Campaign `json:"data,omitempty"`
}
