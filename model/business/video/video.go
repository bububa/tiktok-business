package video

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/business"
)

// Video 视频
type Video struct {
	// ItemID 视频的唯一标识符。
	// 示例：6995316737801538821
	// 更新时间：实时
	ItemID string `json:"item_id,omitempty"`
	// MediaType The type of the post.
	// Enum values:
	// VIDEO: a video post.
	// PHOTO: a photo or text post.
	MediaType string `json:"media_type,omitempty"`
	// IsAd Whether the post is being used in an ad.
	// A post can be promoted via Spark Ads Push or Spark Ads Pull. Learn more about Spark Ads.
	// Supported values:
	// true: The post is used in an ad.
	// false: The post is not used in an ad.
	IsAd bool `json:"is_ad,omitempty"`
	// ThumbnailURL 视频内容缩略图的临时 URL。
	// 过期日期和时间以 Epoch/Unix 时间格式包含在x-expires查询参数中，以秒为单位。
	// 示例：https://p16-sign-va.tiktokcdn.com/tos-maliva-avt-0068/75dec21d63500917fb6ec8bc59415156~c5_300x300.jpeg?x-expires=1614099600&x-signature=PmK%2BWs3LzSzRL2tYs%2FZx7EjG3Gk%3D
	// 更新时间：实时
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// ShareURL 此 TikTok 视频的可分享 URL。请注意，网站在移动设备和桌面设备的呈现有所不同。
	// 示例：https://www.tiktok.com/@feather_in_the_w1nd/video/6994189750940863749?utm_campaign=tt4d_open_api&utm_source=6945731504177676290
	// 更新时间：实时
	ShareURL string `json:"share_url,omitempty"`
	// EmbedURL 此 TikTok 视频的可嵌入链接。此 URL 可用于在外部网站或应用程序中嵌入 TikTok 视频。
	// 如果帖子的隐私设置为"好友可见"或"仅自己可见"，则无法通过此链接观看视频。
	// 示例：https://www.tiktok.com/player/v1/1234567891234567891?music_info=1&description=1&autoplay=1&loop=1&utm_campaign=tt4d_open_api&utm_source=1234567891234567891
	// 更新时间：实时
	EmbedURL string `json:"embed_url,omitempty"`
	// Caption 视频描述。
	// 示例：Nuki: me and me casa 🏠 #cat #kitten #britishshorthair
	// 更新时间：实时
	Caption string `json:"caption,omitempty"`
	// VideoDuration 视频时长，以秒为单位，保留三位小数。
	// 示例：20.001
	// 更新时间：T + 24-48 小时（UTC 时间）
	VideoDuration float64 `json:"video_duration,omitempty"`
	// Likes 点赞数。
	// 视频获得的总点赞数。
	// 示例：243
	// 更新时间：T + 24-48 小时（UTC时间）
	Likes int64 `json:"likes,omitempty"`
	// Comments 评论数。
	// 视频获得的总评论数。
	// 示例：10
	// 更新时间：T + 24-48 小时（UTC 时间）
	Comments int64 `json:"comments,omitempty"`
	// Shares 分享数。
	// 视频被分享的总次数。
	// 本字段返回的分享次数为 TikTok 手机端应用中 For You （推荐）页显示的分享次数，包含外部分享次数和通过 "Send to friends"（分享给好友）按钮内部分享的次数。
	// 例如，若视频外部分享次数为30次，内部分享次数为20次，则本字段返回的值为50。
	// 示例：10
	// 更新时间：T + 24-48 小时（UTC 时间）
	Shares int64 `json:"shares,omitempty"`
	// Favorites 收藏数。
	// 视频被收藏的总次数。
	// 示例：10
	// 更新时间：T + 24-48 小时（UTC 时间）
	Favorites int64 `json:"favorities,omitempty"`
	// CreateTime 视频发布的日期和时间，采用 Unix/epoch 格式和 UTC 时区。
	// 示例：1628461703
	// 更新时间：实时
	CreateTime model.Int64 `json:"create_time,omitempty"`
	// Reach 触达率。
	// 表示至少观看过一次已发布内容的用户。
	// 示例：136
	// 更新时间：T + 24-48 小时（UTC 时间）
	Reach int64 `json:"reach,omitempty"`
	// VideoViews 观看次数。
	// 视频被观看的总次数。
	// 注意:
	// 此指标当视频回放时长大于0，且回放为一次展示中的首次回放时也计数。
	// 若用户刷过广告后又回刷广告，将视作两次展示。第二次展示也会计入播放次数。
	// 示例：603
	// 更新时间：T + 24-48 小时（UTC 时间）
	VideoViews int64 `json:"video_views,omitempty"`
	// TotalTimeWatched 总观看时长。
	// 表示用户观看视频的总时长，以秒为单位。
	// 示例：587.024
	// 更新时间：T + 24-48 小时（UTC 时间）
	TotalTimeWatched float64 `json:"total_time_watched,omitempty"`
	// AverageTimeWatched 平均观看时长。
	// 表示用户观看视频的平均时长，以秒为单位。
	// 示例：3.862
	// 更新时间：T + 24-48 小时（UTC 时间）
	AverageTimeWatched float64 `json:"average_time_watched,omitempty"`
	// FullVideoWatchedRate 完播率。
	// 表示看完此视频的用户的百分比。
	// 示例：0.0395
	// 更新时间：T + 24-48 小时（UTC 时间）
	FullVideoWatchedRate float64 `json:"full_video_watched_rate,omitempty"`
	// NewFollowers 新粉丝数。
	// 关注您账号的新粉丝总数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	NewFollowers int64 `json:"new_followers,omitempty"`
	// ProfileViews 个人主页访问量。
	// 用户通过你的视频访问时查看个人主页的总次数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	ProfileViews int64 `json:"profile_views,omitempty"`
	// WebsiteClicks 网站点击数。
	// 用户通过你的视频访问你个人资料页时点击网站链接的总次数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	WebsiteClicks int64 `json:"website_clicks,omitempty"`
	// PhoneNumberClicks 手机号码点击数。
	// 用户通过你的视频访问你个人资料页时点击手机号码链接的总次数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	PhoneNumberClicks int64 `json:"phone_number_clicks,omitempty"`
	// LeadSubmissions 潜在客户提交次数。
	// 从你的消费者收集的潜在客户总数（如报价、新闻资讯订阅等）。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	LeadSubmissions int64 `json:"lead_submissions,omitempty"`
	// AppDownloadClicks 应用下载链接点击数。
	// 用户通过你的视频访问你个人资料页时点击应用下载链接的总次数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	AppDownloadClicks int64 `json:"app_download_clicks,omitempty"`
	// EmailClicks 邮箱点击数。
	// 用户观看此视频后点击主页邮箱按钮的总次数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	EmailClicks int64 `json:"email_clicks,omitempty"`
	// AddressClicks 地址点击数。
	// 用户观看此视频后点击主页地址按钮的总次数。
	// 示例：123
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为已完成企业注册的企业号返回本指标数据。
	AddressClicks int64 `json:"address_clicks,omitempty"`
	// VideoViewRetention 观众留存率。
	// 此指标说明了在一段时间后仍在观看视频的观众数量。
	// 更新时间：T + 24-48 小时（UTC 时间）
	VideoViewRetention []VideoViewRetention `json:"video_view_retention,omitempty"`
	// ImpressionSources 流量源。此指标可帮助您了解观看视频的用户的流量来源细分，按占比从高到低排列。
	// 示例：[{impression_source: "For You", percentage: 0.7554}, {impression_source: "Follow", percentage: 0.2445}, ...]
	// 更新时间：T + 24-48 小时（UTC 时间）
	ImpressionSources []ImpressionSource `json:"impression_sources,omitempty"`
	// AudienceGenders 观众性别。
	// 按性别列出粉丝的分布情况。该人口统计数据基于许多因素，包括用户在其个人主页中提供的信息。
	// 示例： [{gender: "Female", percentage: 0.7554470336505679}, {gender: "Male", percentage: 0.24455296634943213}, ...]
	// 更新时间：T+ 24-48 小时（UTC 时间）
	AudienceGenders []business.AudienceGender `json:"audience_genders,omitempty"`
	// AudienceCountries 热门国家/地区。
	// 按粉丝的位置（国家/地区）列出粉丝的分布情况。该人口统计数据基于许多因素，包括用户在其个人主页中提供的信息。
	// 示例： [{country: "US", percentage: 0.7554470336505679}, {country: "UK", percentage: 0.24455296634943213}, ...]
	// 更新时间：T+ 24-48 小时（UTC 时间）
	AudienceCountries []business.AudienceCountry `json:"audience_countries,omitempty"`
	// AudienceCities 热门城市。
	// 粉丝所在城市分布。该人口统计数据基于许多因素，包括用户在其个人主页中提供的信息。
	// 更新时间：T+ 24-48 小时（UTC 时间）
	// 注意：仅为粉丝数至少为 100 的 TikTok 账号返回本指标数据。
	AudienceCities []business.AudienceCity `json:"audience_cities,omitempty"`
	// AudienceTypes 观众类型分布。
	// 观新观众和回看观众，以及粉丝和非粉丝在观众中的占比。
	// 示例：[{"percentage": 0,"type": "FOLLOWER_PERCENT"},{"percentage": 0.98,"type": "NEW_VIEWER"},{"percentage": 1,"type": "NON_FOLLOWER_PERCENT"},{"percentage": 0.02,"type": "RETURN_VIEWER"}]
	// 更新时间：T + 24-48 小时（UTC 时间）
	AudienceTypes []AudienceType `json:"audience_types,omitempty"`
	// EngagementLikes 互动点赞。
	// 在视频的某个时间点点赞视频的观众的分布。
	// 更新时间：T + 24-48 小时（UTC 时间）
	EngagementLikes []EngagementLike `json:"engagement_likes,omitempty"`
}

// VideoViewRetention 观众留存率。
type VideoViewRetention struct {
	// Second 视频时长中的某一秒
	Second string `json:"second,omitempty"`
	// Percentage 在当前秒仍在观看此视频的观众占比
	Percentage float64 `json:"percentage,omitempty"`
}

// ImpressionSource 流量源
type ImpressionSource struct {
	// ImpressionSource 流量来源。
	// 枚举值：
	// For You：【推荐】页。
	// Follow：【已关注】页。
	// Sound：音乐页面。
	// Personal Profile：个人资料页（个人主页）
	// Search：搜索页面。
	// Hashtag：其他。
	// Direct Message：私信。
	// 注意：是否返回 Direct Message 的值取决于您能否在 TikTok Studio 中看到对应数据。若 TikTok Studio 中无对应数据，您将无法获取到私信的相关数据。
	ImpressionSource string `json:"impression_source,omitempty"`
	// Percentage 某一流量来源的百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// AudienceType 观众类型分布。
type AudienceType struct {
	// Type 观众类型。
	// 枚举值：
	// NEW_VIEWER：新观众。首次查看您的作品或距离上次查看您的作品已超过 1 年并在最近返回观看的观众。
	// RETURN_VIEWER：回看观众。过去1年内看过你的作品，并在最近返回观看的观众。
	// FOLLOWER_PERCENT：粉丝。TikTok 上目前关注您的观众。
	// NON_FOLLOWER_PERCENT：非粉丝。TikTok 上目前没有关注您的观众。
	Type string `json:"type,omitempty"`
	// Percentage 此观众类型关联的观众百分比
	Percentage float64 `json:"percentage,omitempty"`
}

// EngagementLike 互动点赞。
type EngagementLike struct {
	// Second 视频时长中的某一秒
	Second string `json:"second,omitempty"`
	// Percentage 在当前秒点赞视频的观众占比
	Percentage float64 `json:"percentage,omitempty"`
}
