package subscription

import "github.com/bububa/tiktok-business/enum"

// Subscription 订阅信息
type Subscription struct {
	// AppID 开发者应用 ID
	AppID string `json:"app_id,omitempty"`
	// SubscriptionID 订阅 ID。
	SubscriptionID string `json:"subscription_id,omitempty"`
	// SubscribeEntity 订阅对象。
	// 枚举值:
	// AD_ACCOUNT_SUSPENSION：广告账户的暂时停用状态。
	// LEAD：线索。
	// AD_GROUP：广告组的审核状态。
	// AD：广告的审核状态。
	// TCM_SPARK_ADS：上传至某个 TCM 工作流程 2.0 订单的视频的 Spark Ads 授权状态。
	// CREATIVE_FATIGUE：单个广告、广告组中所有广告或广告账户下所有广告的疲劳状态。
	SubscribeEntity string `json:"subscribe_entity,omitempty"`
	// CallbackURL 回调链接。
	CallbackURL string `json:"callback_url,omitempty"`
	// SubscriptionDetail 订阅内容
	SubscriptionDetail *SubscriptionDetail `json:"subscription_detail,omitempty"`
}

// SubscriptionDetail 订阅内容
type SubscriptionDetail struct {
	// AccessToken 授权访问令牌。
	// 可以通过/oauth2/access_token/接口获取。
	// 注意：若用于配置 Webhook 的access_token被撤销（无论是开发者还是广告主撤销），则 Webhook 订阅将失效，您将无法再接收到 Webhook 事件。此时，您需重新生成access_token，并使用新访问令牌重新配置 Webhook。
	AccessToken string `json:"access_token,omitempty"`
	// ProductFilters 仅当 subscribe_entity 为 API_SERVICE_STATUS 时生效。
	// 要订阅的 API 产品列表。
	// 枚举值：
	// BUSINESS_CENTER_API：商务中心 API。
	// CREATIVES_API：创意管理 API。
	// CATALOG_API：商品库 API。
	// TIKTOK_STORE_API：TikTok Store API。
	// CAMPAIGN_API：推广系列管理 API 。
	// REPORTING_API：报表 API。
	// AUDIENCE_API：受众 API。
	// STREAMING_API：流式 API。
	// EVENTS_API：事件 API。
	// ACCOUNTS_API：账号 API。
	// MENTIONS_API：提及 API。
	// TIKTOK_ONE_API：TikTok One API。
	// DISCOVERY_API：热门发现 API。
	// SPARK_RECOMMEND_API：Spark Ads 推荐 API。
	// BUSINESS_MESSAGING_API：Business Messaging API。
	// 如果未指定product_filters，将订阅所有 API 产品的更新。
	ProductFilters []enum.APIServiceType `json:"product_filters,omitempty"`
	// LeadSource 仅当 subscribe_entity 为 LEAD 时生效。
	// 订阅的线索来源。
	// 枚举值：
	// INSTANT_FORM：通过即时表单生成的线索。
	// DIRECT_MESSAGE：通过绑定的企业号的私信生成的线索。
	//
	// 默认值：INSTANT_FORM。
	LeadSource enum.LeadSource `json:"lead_source,omitempty"`
	// BcID 若 subscribe_entity 为 AD_ACCOUNT_SUSPENSION，您需传入 bc_id 和 advertiser_ids 其中之一。
	// 商务中心 ID。
	// 您将订阅该商务中心所拥有的所有广告账户的暂时停用状态。
	// 注意：您需对该广告账户有管理员权限。若想获取您有管理员权限的商务中心列表，可使用 /bc/get/ 并选择 user_role 为 ADMIN 的商务中心。
	BcID string `json:"bc_id,omitempty"`
	// AdvertiserIDs 若 subscribe_entity 为 AD_ACCOUNT_SUSPENSION，您需传入 bc_id 和 advertiser_ids 其中之一。
	// 广告账户 ID 列表。
	// 您将订阅这些广告账户的暂时停用状态，广告账户可来自不同的商务中心。
	// 最大数量：100。
	// 注意：您需对该广告账户有访问权限。若想获取您有访问权限的广告账户列表，可使用 /bc/asset/get/。将 asset_type 设置为 ADVERTISER 并选择 advertiser_role 为 ADMIN、OPERATOR 或 ANALYST 的广告账户。
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
	// AdvertiserID 当subscribe_entity为 AD_GROUP，AD，LEAD或CREATIVE_FATIGUE时生效。
	// 广告主 ID。
	// 订阅广告组或广告的审核状态时必填（subscribe_entity为AD_GROUP, 或 AD）。
	// 若想订阅线索（subscribe_entity为LEAD）：
	// 当 lead_source 为 INSTANT_FORM 或未传入时，需传入 advertiser_id（订阅与某一广告账户绑定且使用即时表单的线索广告生成的线索）或 library_id（订阅与某一表单库绑定且使用即时表单的线索广告生成的线索），page_id 为可选字段。
	// 当 lead_source 为 DIRECT_MESSAGE 时，需传入 advertiser_id（订阅与某一广告账户绑定的私信广告生成的线索）或 library_id（订阅与某一表单库绑定的私信广告生成的线索），不支持设置 page_id。
	// 若想订阅疲劳状态（subscribe_entity 为 CREATIVE_FATIGUE），advertiser_id必传，您可以选择传入adgroup_id和ad_id其中之一，或均不传入。若只传入advertiser_id，您将订阅广告账户下所有广告的疲劳状态。
	//
	// 注意：若 subscribe_entity 为 LEAD 且传入了 advertiser_id，您需对广告账户有管理员权限。若想获取您有管理员权限的广告账户列表，可使用 /bc/asset/get/。将 asset_type 设置为 ADVERTISER，并挑选返回的 advertiser_role 为 ADMIN 的广告账户。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// LibraryID 当subscribe_entity为 LEAD时生效。
	// 表单库 ID，该表单库存有你想订阅的线索。
	// 您可通过/page/library/get/返回的library_id获取表单库 ID。
	// 若想订阅即时表单线索（subscribe_entity 为 LEAD 且 lead_source 为 INSTANT_FORM 或未传入），需传入 advertiser_id 或 library_id，page_id 为可选参数。若此时只传入 advertiser_id，则将订阅广告主账户（advertiser_id）下的所有即时表单（page_id）。
	// 若想订阅私信线索（subscribe_entity 为 LEAD 且 lead_source 为 DIRECT_MESSAGE），需传入 advertiser_id 或 library_id。不支持设置 page_id。
	LibraryID string `json:"library_id,omitempty"`
	// PageID 当 subscribe_entity 为 LEAD 且 lead_source 为 INSTANT_FORM 或未传入时生效。
	// 当 subscribe_entity 为 LEAD 且 lead_source 为 DIRECT_MESSAGE 时不支持本字段。
	// 即时表单的页面 ID，该表单有您想订阅的线索。
	// 若想获取即时表单列表，可使用 /page/get/ 并将 business_type 设置为 LEAD_GEN。
	// 若想订阅即时表单线索（subscribe_entity 为LEAD 且 lead_source 为 INSTANT_FORM 或未传入lead_source），需传入advertiser_id或library_id，page_id为可选参数。若此时只传入advertiser_id，则将订阅广告主账户（advertiser_id）下的所有线索表单（page_id）。
	PageID string `json:"page_id,omitempty"`
	// AdgroupID 当subscribe_entity为 AD_GROUP或CREATIVE_FATIGUE时生效。
	// 广告组 ID。
	// subscribe_entity为 AD_GROUP时必填。
	// 若您将subscribe_entity设置为 CREATIVE_FATIGUE，且同时传入advertiser_id和adgroup_id，您将订阅所指定的广告组的疲劳状态。
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdID 当subscribe_entity为 AD或CREATIVE_FATIGUE时生效。
	// 广告ID 。
	// subscribe_entity为 AD时必填。
	// 若您将subscribe_entity设置为 CREATIVE_FATIGUE，且同时传入advertiser_id和ad_id，您将订阅所指定的广告的疲劳状态。
	AdID string `json:"ad_id,omitempty"`
	// TcmAccountID subscribe_entity为 TCM_SPARK_ADS时必填。
	// TikTok Creator Marketplace账户ID。
	TcmAccountID string `json:"tcm_account_id,omitempty"`
	// VideoID subscribe_entity为 TCM_SPARK_ADS时必填。
	// 上传至某一 TCM 工作流程 2.0 订单的视频的 ID。
	// 若想获取上传至某一 TCM 工作流程 2.0 订单的视频的 ID 列表，可使用/tcm/order/get/v2/或/tcm/report/get/v2/。
	VideoID string `json:"video_id,omitempty"`
}
