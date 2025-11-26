package audienceinsight

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InfoRequest 获取潜在受众详情 API Request
type InfoRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceID 当请求中未指定locations时必填。
	// 自定义受众ID。
	// 注意：
	// 当指定了custom_audience_id时，不能使用selected_audience或locations。否则会报错。
	// 您输入的受众ID必须属于您指定的广告商。
	CustomAudienceID string `json:"custom_audience_id,omitempty"`
	// Locations 当请求中未指定custom_audience_id时必填。
	// 地域。
	// 最大数量：5。
	// 注意：当指定了locations，不能使用custom_audience_id。否则会报错。
	Locations []Location `json:"locations,omitempty"`
	// Dimensions 使用该字段控制返回响应的维度。
	// 枚举值：age，gender，country，top_regions，operating_system，operating_system_version，device_price， top_interests，bottom_interests，engagement，ad_interest_categories，top_hashtags。
	// 注意：目前，每次API调用只能输入一个维度。
	Dimensions []string `json:"dimensions,omitempty"`
	// SelectedAudience 您想筛选的受众
	SelectedAudience *SelectedAudience `json:"selected_audience,omitempty"`
}

// Location 地域
type Location struct {
	// CountryCode 国家代码。查看附录-地域定向获取详情
	CountryCode string `json:"country_code,omitempty"`
	// LocationIDs 地域ID。
	// 默认情况下，将选择指定国家代码下的所有地域ID。您输入的地域ID必须在指定的国家代码范围内
	LocationIDs []string `json:"location_ids,omitempty"`
}

// SelectedAudience 您想筛选的受众
type SelectedAudience struct {
	// TopRegionsCountryCode 若您在dimensions字段中选择了top_regions并指定了locations，则该字段为必填。 当您在dimensions字段中选择了top_regions时，相应的返回字段将包含您在locations中指定国家的子区域
	TopRegionsCountryCode []string `json:"top_regions_country_code,omitempty"`
	// LocationIDs 您想要定向的地域 ID 列表。
	LocationIDs []string `json:"location_ids,omitempty"`
	// AgeGroups 定向受众年龄。 枚举值：详见枚举值-受众年龄区间
	AgeGroups []enum.Age `json:"age_groups,omitempty"`
	// Gender 定向受众性别。
	// 枚举值: GENDER_FEMALE,GENDER_MALE,GENDER_UNLIMITED
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// DevicePriceRange 设备价格区间（10000代表1000+）。
	// 该数字必须是50的倍数。
	DevicePriceRanges []int `json:"device_price_ranges,omitempty"`
	// MinIosVersion 受众最低iOS版本。
	// 枚举值：详见枚举值-最低iOS版本。
	MinIosVersion string `json:"min_ios_version,omitempty"`
	// MinAndroidVersion 受众最低Android版本。枚举值：详见枚举值-最低Android版本
	MinAndroidVersion string `json:"min_android_version,omitempty"`
	// OperatingSystems 受众操作系统。 枚举值：ANDROID， IOS。只允许传入一个值
	OperatingSystems []enum.OperatingSystem `json:"operating_systems,omitempty"`
	// InterestCategoryIDs 兴趣分类。
	InterestCategoryIDs []string `json:"interest_category_ids,omitempty"`
	// VideoActionCategoryIDs 视频动作类别ID。
	// 您可以使用/tool/action_category/获取可用的视频互动类别列表（响应中的action_scene = VIDEO_RELATED）。
	// 您可以使用/tool/targeting_category/recommend/获取推荐的视频互动类别列表（响应中的scene = VIDEO_RELATED）。
	VideoActionCategoryIDs []string `json:"video_action_category_ids,omitempty"`
	// CreatorActionCategoryIDs 创作者动作类别ID。
	// 您可以使用/tool/action_category/获取可用的创作者互动类别列表（响应中的action_scene = CREATOR_RELATED）。您可以使用/tool/targeting_category/recommend/获取推荐的创作者互动类别列表（响应中的scene = CREATOR_RELATED）。
	CreatorActionCategoryIDs []string `json:"creator_action_category_ids,omitempty"`
	// HashtagActionCategoryIDs 话题ID。
	// 您可以使用/tool/hashtag/recommend/获取推荐的话题ID列表。您也可以使用/tool/hashtag/get/通过ID获取话题。
	HashtagActionCategoryIDs []string `json:"hashtag_action_category_ids,omitempty"`
}

// Encode implements PostRequest
func (r *InfoRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// InfoResponse 获取潜在受众详情 API Response
type InfoResponse struct {
	model.BaseResponse
	Data *Info `json:"data,omitempty"`
}

// Info 获取潜在受众详情
type Info struct {
	// Age 年龄。如果您在请求中的dimensions里输入了age，则会返回此字段。
	// 在下面列出的all和selected对象中，有两个字段传达了具体信息。详情见下方的代码示例。
	// label：年龄组标签。它的数据类型是string。
	// rate：标签所占百分比。它的数据类型是number。
	Age *InfoItem[LabelRate] `json:"age,omitempty"`
	// Gender 性别。如果您在请求中的dimensions里输入了gender，则会返回此字段。
	// 在下面列出的all和selected对象中，有两个字段传达了具体信息。
	// label：性别标签。它的数据类型是string。
	// rate：标签所占百分比。它的数据类型是number。
	Gender *InfoItem[LabelRate] `json:"gender,omitempty"`
	// Country 国家。如果您在请求中的dimensions里输入了country，则会返回此字段。
	// 在下面列出的all和selected对象中，有两个字段传达了具体信息。
	// label：国家名。它的数据类型是string。
	// rate：国家所占百分比。它的数据类型是number。
	Country *InfoItem[LabelRate] `json:"country,omitempty"`
	// TopRegion 国家。如果您在请求中的dimensions里输入了country，则会返回此字段。
	// 在下面列出的all和selected对象中，有几个字段传达了具体信息。
	// label：地区名。它的数据类型是string。
	// rate：地区所占百分比。它的数据类型是number。
	// region_id：地区ID。它的数据类型是string。
	// country_code：国家代码。它的数据类型是string。
	TopRegion *InfoItem[RegionLabelRate] `json:"top_region,omitempty"`
	// OperatingSystem 操作系统。如果您在请求中的dimensions里输入了operating_system，则会返回此字段。
	// 在下面列出的all和selected对象中，有两个字段传达了具体信息。
	// label：操作系统。它的数据类型是string。
	// rate：操作系统所占百分比。它的数据类型是number。
	OperatingSystem *InfoItem[LabelRate] `json:"operating_system,omitempty"`
	// OperatingSystemVersion 操作系统版本。如果您在请求中的dimensions里输入了operating_system_version，则会返回此字段。
	// 在下面列出的all和selected对象中，有两个字段传达了具体信息。
	// label：操作系统版本。它的数据类型是string。
	// rate：操作系统版本所占百分比。它的数据类型是number。
	OperatingSystemVersion *InfoItem[LabelRate] `json:"operating_system_version,omitempty"`
	// DevicePrice 设备价格。如果您在请求中的dimensions里输入了device_price，则会返回此字段。
	// 在下面列出的all和selected对象中，有两个字段传达了具体信息。
	// label：设备价格。它的数据类型是string。
	// rate：设备价格所占百分比。它的数据类型是number。
	DevicePrice *InfoItem[LabelRate] `json:"device_price,omitempty"`
	// TopInterests 排名靠前的兴趣。如果您在请求中的dimensions里输入了top_interests，则会返回此字段。
	// 在下面列出的all和selected对象中，有几个字段传达了具体信息。
	// label：兴趣类别标签。它的数据类型是string。
	// rate：兴趣类别标签所占百分比。它的数据类型是number。
	// ad_interest_category：一级兴趣类别。它的数据类型是string。
	TopInterests *InfoItem[InterestCategoryLabelRate] `json:"top_interests,omitempty"`
	// BottomInterests 排名靠后的兴趣。如果您在请求中的dimensions里输入了bottom_interests，则会返回此字段。
	// 在下面列出的all和selected对象中，有几个字段传达了具体信息。
	// label：兴趣类别标签。它的数据类型是string。
	// rate：兴趣类别标签所占百分比。它的数据类型是number。
	// ad_interest_category：一级兴趣类别。它的数据类型是string。
	BottomInterests *InfoItem[InterestCategoryLabelRate] `json:"bottom_interests,omitempty"`
	// AdInterestCategories 兴趣类别。如果您在请求中的dimensions里输入了ad_interest_categories，则会返回此字段。
	// 在下面列出的all和selected对象中，有几个字段传达了具体信息。
	// label：兴趣类别标签。它的数据类型是string。
	// rate：兴趣类别标签所占百分比。它的数据类型是number。
	// type：兴趣类别层级。它的数据类型是string。例如：“first”。
	AdInterestCategories *InfoItem[InterestCategoryTypeLabelRate] `json:"ad_interest_categories,omitempty"`
	// TopHashtags 热门话题。如果您在请求中的dimensions里输入了top_hashtags，则会返回此字段
	TopHashtags *TopHashtags `json:"top_hashtags,omitempty"`
	// Engagement 互动数据指标。如果您在请求中的dimensions里输入了engagement，则会返回此字段。
	// 在下面列出的all和selected对象中，有几个字段传达了具体信息。
	// daily_sessions：用户每天平均打开应用的次数。它的数据类型是number。
	// daily_session_duration_in_minutes：用户每天平均使用应用的分钟数。它的数据类型是number。
	// daily_video_views：用户每天平均观看视频的次数。它的数据类型是number。
	// daily_comments： 用户每天平均评论数量。它的数据类型是number。
	// daily_likes：用户每天平均点赞数量。它的数据类型是number。
	// daily_shares：用户每天平均分享数量。它的数据类型是number。
	// total_video_uploads：选定/所有用户每天平均上传视频的次数。它的数据类型是number。
	Engagement *InfoItem[Engagement] `json:"engagement,omitempty"`
}

type InfoItem[T LabelRate | RegionLabelRate | InterestCategoryLabelRate | InterestCategoryTypeLabelRate | Engagement] struct {
	// All 所有类别数据
	All []T `json:"all,omitempty"`
	// Selected 您筛选的类别数据
	Selected []T `json:"selected,omitempty"`
}

type LabelRate struct {
	Label string  `json:"label,omitempty"`
	Rate  float64 `json:"rate,omitempty"`
}

type RegionLabelRate struct {
	LabelRate
	RegionID    string `json:"region_id,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type InterestCategoryLabelRate struct {
	LabelRate
	AdInterestCategory string `json:"ad_interest_category,omitempty"`
}

type InterestCategoryTypeLabelRate struct {
	LabelRate
	Type string `json:"type,omitempty"`
}

// TopHashtags 热门话题
type TopHashtags struct {
	// HashtagID 话题ID
	HashtagID string `json:"hashtag_id,omitempty"`
	// HashtagName 话题名称
	HashtagName string `json:"hashtag_name,omitempty"`
	// Count 该话题被引用的次数
	Count int64 `json:"count,omitempty"`
}

// Engagement 互动数据指标
type Engagement struct {
	// DailySessions：用户每天平均打开应用的次数。它的数据类型是number。
	DailySessions int64 `json:"daily_sessions,omitempty"`
	// DailySessionDurationInMinutes：用户每天平均使用应用的分钟数。它的数据类型是number。
	DailySessionDurationInMinutes int `json:"daily_session_duration_in_minutes,omitempty"`
	// DailyVideoViews：用户每天平均观看视频的次数。它的数据类型是number。
	DailyVideoViews int64 `json:"daily_video_views,omitempty"`
	// DailyComments： 用户每天平均评论数量。它的数据类型是number。
	DailyComments int64 `json:"daily_comments,omitempty"`
	// DailyLikes：用户每天平均点赞数量。它的数据类型是number。
	DailyLikes int64 `json:"daily_likes,omitempty"`
	// DailyShares：用户每天平均分享数量。它的数据类型是number。
	DailyShares int64 `json:"daily_shares,omitempty"`
	// TotalVideoUploads：选定/所有用户每天平均上传视频的次数。它的数据类型是number。
	TotalVideoUploads int64 `json:"total_video_uploads,omitempty"`
}
