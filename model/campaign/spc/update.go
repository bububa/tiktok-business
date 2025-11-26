package spc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/ad/aco"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 更新 Smart+ 推广系列 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName 推广系列名称
	CampaignName string `json:"campaign_name,omitempty"`
	// SpecialIndustries 特殊广告分类。
	// 枚举值:
	// HOUSING：房地产，房屋保险，抵押贷款或相关的广告。
	// EMPLOYMENT：工作机会，实习机会，职业认证项目或相关的广告。
	// CREDIT：信用卡申请，汽车贷款，长期融资或相关的广告。
	// 注意：
	// 在推广系列商品来源为商品库和推广系列商品来源非商品库的购物广告中使用特殊广告分类目前为单独的白名单功能。如需使用这些功能，请联系您的TikTok销售代表。
	// 本字段对于注册地为美国或加拿大的广告主已全量。若您是注册地非美国或加拿大的广告主且想在定向美国或加拿大的广告中使用特殊广告分类，您需申请额外的白名单。
	SpecialIndustries []enum.SpecialIndustry `json:"special_industries,omitempty"`
	// ProductSpecificType 选择商品的维度。
	//  枚举值：
	// ALL：允许 TikTok 从所有商品中动态选择。
	// PRODUCT_SET：请选择一个商品系列。TikTok 将动态选择该商品系列中的商品。
	// CUSTOMIZED_PRODUCTS：选择自定义数量的特定商品
	ProductSpecificType enum.ProductSpecificType `json:"product_specific_type,omitempty"`
	// ProductSetID 商品系列 ID
	ProductSetID string `json:"prodduct_seit_id,omitempty"`
	// ProductIDs 商品库商品的商品 ID 列表
	ProductIDs []string `json:"product_ids,omitempty"`
	// BidPrice 出价。（成本上限策略下）想尽可能达到的平均单次成效成本
	BidPrice float64 `json:"bid_price,omitempty"`
	// ConversionBidPrice oCPM转化出价
	ConversionBidPrice float64 `json:"conversion_bid_price,omitempty"`
	// RoasBid 用于价值优化的 ROAS 目标值
	RoasBid float64 `json:"roas_bid,omitempty"`
	// LocationIDs 定向地域 ID。使用tool/region 获取定向地域ID
	LocationIDs []string `json:"location_ids,omitempty"`
	// Gender 定向受众性别
	// 枚举值: GENDER_FEMALE,GENDER_MALE,GENDER_UNLIMITED。
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// ExcludeAgeUnderEighteen 是否排除十八岁以下受众
	ExcludeAgeUnderEighteen *bool `json:"exclude_age_under_eighteen,omitempty"`
	// ExcludeAudienceIDs 排除受众 ID 列表
	ExcludeAudienceIDs []string `json:"exclude_audience_ids,omitempty"`
	// CommentDisabled 是否允许用户在TikTok上评论您的广告
	CommentDisabled *bool `json:"comment_disabled,omitempty"`
	// ShareDisabled 本广告组中的广告是否禁止分享到第三方平台
	ShareDisabled *bool `json:"share_disabled,omitempty"`
	// BlockedPangleAppIDs Pangle 媒体屏蔽列表 ID
	BlockedPangleAppIDs []string `json:"blocked_pangle_app_ids,omitempty"`
	// Budget 预算。
	// 当开启了推广系列预算优化(budget_optimize_on)时，返回0.0。
	// 对于 TopView 广告，本字段代表预估折后预算。
	Budget float64 `json:"budget,omitempty"`
	// ScheduleType 广告投放时间类型。
	// 枚举值: SCHEDULE_FROM_NOW，SCHEDULE_START_END。如果您选择了SCHEDULE_START_END，您需要明确投放开始和结束时间。 如果您选择了SCHEDULE_FROM_NOW，您只需要明确投放开始时间。
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime 广告投放起始时间(UTC+0)。
	// 格式：YYYY-MM-DD HH:MM:SS。
	ScheduleStartTime model.DateTime `json:"schedule_start_time,omitzero"`
	// ScheduleEndTime 广告投放结束时间(UTC+0)。
	// 格式： YYYY-MM-DD HH:MM:SS
	ScheduleEndTime model.DateTime `json:"schedule_end_time,omitzero"`
	// Dapparting 广告投放安排。格式为48x7的字符串。字符只能为0或者1。0代表不投放，1代表投放。每个字符对应一个30分钟的时间段。第一个字符对应的是周一的凌晨0:01-0:30，第二个字符对应0:31-1:00，依次类推。最后一个字符代表周日23:31-0:00。
	// 注意：
	// 不设置，全部设置为0，或者全部都设置为1都代表了要进行全时投放。
	Dayparting *string `json:"dayparting,omitempty"`
	// IdentityID 不使用 Spark Ads 时的认证身份 ID
	IdentityID *string `json:"identity_id,omitempty"`
	// MediaInfoList 媒体信息列表
	MediaInfoList []aco.MediaInfo `json:"media_info_list,omitempty"`
	// CatalogCreativeToggle 是否启用商品库素材的自动选择。
	//  支持的值：true，false。
	// 若本字段设置为 true，系统将使用可用的广告样式（商品库视频、商品库轮播和单视频）自动生成广告。
	CatalogCreativeToggle *bool `json:"catalog_creative_toggle,omitempty"`
	// TitleList 广告文案列表。广告文案作为您的广告素材一部分向受众展示，传递您想传达的信息
	TitleList []aco.Title `json:"title_list,omitempty"`
	// CallToActionList 行动引导文案列表
	CallToActionList []aco.CallToAction `json:"call_to_action,omitempty"`
	// CallToActionID 行动引导文案素材包 ID
	CallToActionID *string `json:"call_to_action_id,omitempty"`
	// CardList 卡片ID列表。长度范围：[0,1]
	CardList []aco.Card `json:"card_list,omitempty"`
	// PageList 页面 ID 列表
	PageList []aco.Page `json:"page_list,omitempty"`
	// DeeplinkType 深度链接类型
	DeeplinkType enum.DeeplinkType `json:"deeplink_type,omitempty"`
	// Deeplink 深度链接，用户下载应用后点击链接跳转至的特定页面
	Deeplink *string `json:"deeplink,omitempty"`
	// LandingPageURLs 落地页URL列表
	LandingPageURLs []aco.LandingPageURL `json:"landing_page_urls,omitempty"`
	// ImpressionTrackingURL 默认展示监测 URL
	ImpressionTrackingURL *string `json:"impression_tracking_url,omitempty"`
	// ClickTrackingURL 点击监测 URL
	ClickTrackingURL *string `json:"click_tracking_url,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新 Smart+ 推广系列 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *Campaign `json:"data,omitempty"`
}
