package savedaudience

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取已保存受众详情 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// SavedAudienceIDs 已保存受众 ID 列表。
	// 最大数量：100。
	// 若想创建已保存受众并获取已保存受众 ID，可使用 /dmp/saved_audience/create/。
	SavedAudienceIDs []string `json:"saved_audience_ids,omitempty"`
}

// Encode implements GetRequest
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.SavedAudienceIDs) > 0 {
		values.Set("saved_audience_ids", string(util.JSONMarshal(r.SavedAudienceIDs)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ListResponse 获取已保存受众详情 API Response
type ListResponse struct {
	model.BaseResponse
	Data struct {
		// SavedAudiences 已保存受众列表
		SavedAudiences []Audience `json:"saved_audiences,omitempty"`
	} `json:"data"`
}

type Audience struct {
	// SavedAudienceID 已保存受众ID
	SavedAudienceID string `json:"saved_audience_id,omitempty"`
	// SavedAudienceName 已保存受众名称。
	// 名称不能超过512字符，不能包含emoji。
	SavedAudienceName string `json:"saved_audience_name,omitempty"`
	// LocationIDs 您想要定向的地域 ID 列表。
	// 您可以使用/tool/region/或/tool/targeting/search/接口，查询您可以投放的地区以及相应ID。
	// 注意：
	// 定向地域不允许重叠。例如，不支持同时定向美国和加利福尼亚州。
	LocationIDs []string `json:"location_ids,omitempty"`
	// Gender 定向受众性别。
	// 枚举值: GENDER_FEMALE,GENDER_MALE,GENDER_UNLIMITED
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// AgeGroups 定向受众年龄。 枚举值：详见枚举值-受众年龄区间
	AgeGroups []enum.Age `json:"age_groups,omitempty"`
	// Languages 受众语言。 您可通过/tool/language/获取语言代码。若您不想限制受众语言，可为本字段传空值或不传入本字段
	Languages []string `json:"languages,omitempty"`
	// AudienceIDs 受众ID列表。
	// 通过/dmp/custom_audience/list/接口获取受众ID。
	AudienceIDs []string `json:"audience_ids,omitempty"`
	// ExcludedAudienceIDs 要排除的受众ID列表。
	// 通过/dmp/custom_audience/list/接口获取受众ID。
	ExcludedAudienceIDs []string `json:"excluded_audience_ids,omitempty"`
	// InterestCategoryIDs 兴趣分类。
	// 使用/tool/target_recommend_tags/接口根据您的定向地区和行业获取推荐兴趣分类，或使用/tool/interest_category/接口获取所有支持的兴趣分类。如果指定了兴趣，不支持兴趣定向的版位上的用户将被限制投放，若要确保所有版位可投，请勿传该字段。
	InterestCategoryIDs []string `json:"interest_category_ids,omitempty"`
	// InterestKeywordIDs 用于定向的兴趣关键词ID列表。
	// 您可以使用/tool/interest_keyword/recommend/接口获取推荐兴趣关键词。
	InterestKeywordIDs []string `json:"interest_keyword_ids,omitempty"`
	// Actions 行为分类ID列表
	Actions []Action `json:"actions,omitempty"`
	// OperatingSystems 受众操作系统。 枚举值：ANDROID， IOS。只允许传入一个值
	OperatingSystems []enum.OperatingSystem `json:"operating_systems,omitempty"`
	// MinAndroidVersion 受众最低Android版本。枚举值：详见枚举值-最低Android版本
	MinAndroidVersion string `json:"min_android_version,omitempty"`
	// MinIosVersion 受众最低iOS版本。
	// 枚举值：详见枚举值-最低iOS版本。
	MinIosVersion string `json:"min_ios_version,omitempty"`
	// DeviceModelIDs 设备机型ID列表
	DeviceModelIDs []string `json:"device_model_ids,omitempty"`
	// NetworkTypes 网络类型。默认值：unlimited。 详见枚举值-网络类型
	// NetworkTypes 网络类型。
	// 详见枚举值-网络类型。
	NetworkTypes []enum.NetworkType `json:"network_types,omitempty"`
	// CarrierIDs 运营商 ID列表
	CarrierIDs []string `json:"carrier_ids,omitempty"`
	// DevicePriceRange 设备价格区间（10000代表1000+）。
	// 该数字必须是50的倍数。
	// 重要提示：最终的实际上限将在您设定的上限基础上增加50，您可以在TikTok广告管理平台的广告组设置中看到实际上限。例如，如果您设置的价格区间是[0, 250]，实际区间则为[0, 300]。
	DevicePriceRanges []int `json:"device_price_ranges,omitempty"`
}

// Action 行为分类
type Action struct {
	// ActionCategoryIDs 您想要定向的行为分类ID列表
	ActionCategoryIDs []string `json:"action_category_ids,omitempty"`
	// ActionScene 行为场景。枚举值：VIDEO_RELATED(视频行为)，CREATOR_RELATED（创作者互动行为），HASHTAG_RELATED（视频标签相关）
	ActionScene enum.ActionScene `json:"action_scene,omitempty"`
	// ActionPeriod 选择n天内发生的行为。
	// 枚举值: 0，7，15。如果action_scene为CREATOR_RELATED或HASHTAG_RELATED，不论您传入何值，系统将默认使用0。0表示不设置行为发生的明确时间范围。
	ActionPeriod *int `json:"action_period,omitempty"`
	// VideoUserActions 所选用户行为种类下想要定向的的具体用户行为。
	// 若action_scene为VIDEO_RELATED，枚举值为：WATCHED_TO_END（完播），LIKED（点赞），COMMENTED（评论），SHARED（分享）。
	// 若action_scene为CREATOR_RELATED，枚举值为：FOLLOWING（关注），VIEW_HOMEPAGE（浏览主页）。
	// 若action_scene为HASHTAG_RELATED，枚举值为： VIEW_HASHTAG（浏览话题标签）。
	VideoUserActions []enum.VideoUserAction `json:"video_user_actions,omitempty"`
}
