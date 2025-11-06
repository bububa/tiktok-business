package video

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取 TikTok 账号视频数据 API Request
type ListRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// 请求字段。若不进行设置，仅返回默认字段。
	// 默认值：["item_id"]。
	// 若您想请求多个数据字段，您需在fields的值中包含"item_id" 。否则，可能出现报错。
	// 示例：["item_id", "create_time"]。
	// 注意：所有视频互动数据指标是视频投放期间的所有数据汇总。与主页级数据指标不同，此汇总不受 60 天回溯期的限制。
	//
	// 若想返回指定的字段，需要有对应的字段权限。您可根据返回结构中每个字段对应的“所需权限”列，查看所需的权限。
	// 若想确认您已具备对应的权限范围，您可查看 /tt_user/token_info/get/接口返回的权限范围。若您还未获得所需的权限，则需要遵循授权中列出的步骤，使 TikTok 账户所有者授予您的 App 对应的权限。
	//
	// 支持字段：
	//
	// item_id：视频的唯一标识符。
	// thumbnail_url：视频预览缩略图 URL。
	// share_url：视频的可分享 URL 。
	// embed_url：视频的可嵌入 URL。
	// caption：视频的字幕/描述。
	// video_duration：视频时长（以秒为单位）。
	// likes：点赞数。视频获得的总点赞数。
	// comments：评论数。视频获得的总评论数。
	// shares：分享数。视频被分享的总次数。
	// favorites：收藏数。视频被收藏的总次数。
	// create_time：视频发布的日期和时间（Unix/epoch 格式）。
	// reach：触达率。表示至少观看过一次已发布内容的用户。
	// video_views：观看次数。视频被观看的总次数。
	// 注意:
	//
	// 此指标当视频回放时长大于0，且回放为一次曝光中的首次回放时也计数。
	// 若用户刷过广告后又回刷广告，将视作两次曝光。第二次曝光也会计入播放次数。
	// total_time_watched：总观看时长。表示用户观看视频的总时长。
	// average_time_watched：平均观看时长。表示用户观看视频的平均时长。
	// full_video_watched_rate：完播率。表示看完此视频的用户的百分比。
	// new_followers：新粉丝数。关注您账号的新粉丝总数。
	// profile_views：个人主页访问量。用户通过你的视频访问时查看个人主页的总次数。
	// website_clicks：网站点击数。用户通过你的视频访问你个人资料页时点击网站链接的总次数。
	// phone_number_clicks：手机号码点击数。用户通过你的视频访问你个人资料页时点击手机号码链接的总次数。
	// lead_submissions：潜在客户提交次数。从你的消费者收集的潜在客户总数（如报价、新闻资讯订阅等）。
	// app_download_clicks：应用下载链接点击数。用户通过你的视频访问你个人资料页时点击应用下载链接的总次数。
	// email_clicks：邮箱点击数。用户观看此视频后点击主页邮箱按钮的总次数。
	// address_clicks：地址点击数。用户观看此视频后点击主页地址按钮的总次数。
	// video_view_retention：观众留存率。此指标说明了在一段时间后仍在观看视频的观众数量。
	// impression_sources：流量来源。此指标可帮助您了解观看视频的用户的流量来源细分。
	// audience_genders：性别。观众的性别分布。该数据基于许多因素，包括观众提供的信息。
	// audience_countries：热门国家/地区。观众所在国家或地区分布。该数据基于许多因素，包括用户提供的信息。
	// audience_cities：热门城市。观众所在城市分布。该数据基于许多因素，包括用户提供的信息。
	// audience_types：观众类型分布。新观众和回看观众，以及粉丝和非粉丝在观众中的占比。
	// engagement_likes：互动点赞。在视频的某个时间点点赞视频的观众的分布。
	Fields []string `json:"fields,omitempty"`
	// Filters 对结果集应用的筛选器
	Filters *ListFilter `json:"filters,omitempty"`
	// Cursor 分页光标。如果返回参数has_more为true，请将返回中此参数中的光标值传递到后续请求，以获取下一页结果。
	// 默认值：当前时间，采用 Epoch/Unix 时间戳，以毫秒为单位。
	// 注意：光标值采用 UTC Unix 时间戳，以毫秒为单位。您可以传递定制时间戳，获取用户在所提供的时间戳之前发布的视频。
	Cursor int64 `json:"cursor,omitempty"`
	// MaxCount 每页返回的最大视频数。
	// 默认值：10。
	// 最大值：20。
	// 注意：受信任与安全政策影响，接口可能会返回少于 max_count 的视频数，即使 返回参数has_more 为 true。
	MaxCount int `json:"max_count,omitempty"`
}

// ListFilter 对结果集应用的筛选器
type ListFilter struct {
	// VideoIDs 用于筛选洞察信息的视频 ID 列表。
	// 示例：["6995316737801538821", "8015388216995316737"]。
	VideoIDs []string `json:"video_ids,omitempty"`
	// AdPostOnly 仅当传入 video_ids 时生效。
	// 是否从指定的视频列表（video_ids）中筛选已用作 Spark Ads 的视频。
	// 支持的值：true，false。
	// 默认值：false。
	AdPostOnly bool `json:"ad_post_only,omitempty"`
}

// Encode implements GetRequest
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if r.Filters != nil {
		values.Set("filters", string(util.JSONMarshal(r.Filters)))
	}
	if r.Cursor > 0 {
		values.Set("cursor", strconv.FormatInt(r.Cursor, 10))
	}
	if r.MaxCount > 0 {
		values.Set("max_count", strconv.Itoa(r.MaxCount))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ListResponse 获取 TikTok 账号视频数据 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// Videos 视频列表，包括展示在账户主页的原生视频，以及"只作为广告展示“的帖子（即专门用作公开可访问的广告、但不展示在账户主页的视频帖子）
	Videos []Video `json:"videos,omitempty"`
	// Cursor 用于获取下一页结果的光标，可传递到后续接口请求的 cursor 查询参数。
	// 更新时间：实时
	Cursor int64 `json:"cursor,omitempty"`
	// HasMore 是否有下一页数据。
	// 更新时间：实时
	HasMore bool `json:"has_more,omitempty"`
}
