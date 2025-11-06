package business

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取 TikTok 账号主页数据 API Request
type GetRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// StartDate 查询起始日期（闭区间）。格式如下：2021-06-01。日期基于UTC 时区。
	// 若不进行设置，start_date 将为[当前日期- 7 天]。
	// 支持的最长回溯期为 60 天。
	StartDate string `json:"start_date,omitempty"`
	// EndDate 查询结束日期（闭区间）。格式如下：2021-06-01。日期基于UTC 时区。
	// 若不进行设置，end_date将为 [当前日期 - 1 天]。
	EndDate string `json:"end_date,omitempty"`
	// Fields 请求字段。若不进行设置，仅返回默认字段。
	// 默认值：["display_name", "profile_image"]。
	// 若想返回指定的字段，需要有对应的字段权限。您可根据返回结构中每个字段对应的“所需权限”列，查看所需的权限。
	// 若想确认您已具备对应的权限范围，您可查看 /tt_user/token_info/get/接口返回的权限范围。若您还未获得所需的权限，则需要遵循授权中列出的步骤，使 TikTok 账户所有者授予您的 App 对应的权限。
	// 示例：["display_name", "profile_image"]。
	// 支持字段：
	// is_business_account：该 TikTok 账户是否为 TikTok 企业号。欲获取本指标的数据，需先从 TikTok 账户所有者处获取权限范围"user.account.type"。
	// 注意：欲确认是否已经获取了"user.account.type"的权限范围，您可以查看/tt_user/token_info/get/返回的scope。若您还未获取该权限，则需遵循授权中列出的步骤，使 TikTok 账户所有者重新对您的App进行授权，授予您该权限。
	//
	// profile_image： TikTok 账号主页头像的 URL。
	// username：TikTok 账户的用户名。
	// profile_deep_link：TikTok 账号主页链接。
	// display_name ： TikTok 账号的显示名称（昵称）。
	// bio_description：TikTok 账号的个人简介。
	// is_verified：TikTok 账号是否带有认证徽章（蓝 V 认证），从而表明账户所有者的真实性已经过核验。
	// following_count：TikTok 账号关注的账号数。
	// followers_count： TikTok 账号粉丝数。
	// 注意：若您在 fields 的值中包含 followers_count，followers_count将同时作为总指标和每日指标返回。
	//
	// total_likes：总点赞数。用户对已发布视频的总点赞次数。
	// videos_count：TikTok 账号发布的公开视频总数。
	// video_views： 每日观看次数。您的视频的每日观看次数。
	// 注意：
	//
	// 此指标当视频回放时长大于 0，且回放为一次曝光中的首次回放时也计数。
	// 若用户刷过广告后又回刷广告，将视作两次曝光。第二次曝光也会计入播放次数。
	// unique_video_views：每日触达的观众。每日至少观看过一次已发布内容的用户。
	// profile_views：每日个人主页访问量。每日用户查看您的个人主页的次数。
	// likes：每日点赞数。每日用户对已发布视频的点赞次数。
	// comments：每日评论数。每日用户对已发布视频的评论次数。
	// shares：每日分享次数。每日用户对已发布视频的分享次数。
	// phone_number_clicks：每日手机号码点击数。在选定日期范围内收集的每日手机号码链接的总点击数。
	// lead_submissions：每日潜在客户提交次数。在选定日期范围内收集的从您的消费者收集的每日潜在客户总数（如报价、新闻资讯订阅等）。
	// app_download_clicks：每日应用下载链接点击数。在选定日期范围内收集的每日应用下载链接的总点击数。
	// bio_link_clicks：每日个人简介链接点击数。在选定日期范围内收集的每日个人简介链接的总点击数。
	// email_clicks：每日邮箱点击数。在选定日期范围内收集的每日主页邮箱按钮的总点击数。
	// address_clicks：每日地址点击数。在选定日期范围内收集的每日主页地址按钮的总点击数。
	// daily_total_followers：每日净增长。表示在一定时段内每日关注您账号的用户数。
	// daily_new_followers：每日新粉丝数。表示在一定时段内每日获得的新粉丝总数。
	// daily_lost_followers：每日掉粉数。表示在一定时段内每日失去的粉丝总数。
	// audience_activity：每小时粉丝活动。表示您的粉丝一天中在 TikTok 上活跃的时段。
	// 仅为粉丝数至少为 100 的账号返回本指标数据。
	// engaged_audience：每日互动的观众。表示每日与您发布的至少一项内容有过互动（点赞、评论、分享）的用户。
	// audience_ages：粉丝年龄。按年龄列出粉丝的分布情况。该人口统计数据基于许多因素，包括用户在其个人主页中提供的信息。
	// 仅为粉丝数至少为 100 的账号返回本指标数据。
	// audience_genders： 粉丝性别。按性别列出粉丝的分布情况。该人口统计数据基于许多因素，包括用户在其个人主页中提供的信息。
	// 仅为粉丝数至少为 100 的账号返回本指标数据。
	// audience_countries：粉丝热门国家/地区。按粉丝的位置（国家/地区）列出粉丝的分布情况。该人口统计数据基于许多因素，包括用户在其个人主页中提供的信息。
	// 仅为粉丝数至少为 100 的账号返回本指标数据。
	// audience_cities：粉丝热门城市。粉丝所在城市分布。该人口统计数据基于许多因素，包括用户在其个人主页中提供的信息。
	// 仅为粉丝数至少为 100 的账号返回本指标数据。
	Fields []string `json:"fields,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	if r.StartDate != "" {
		values.Set("start_date", r.StartDate)
	}
	if r.EndDate != "" {
		values.Set("end_date", r.EndDate)
	}
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 获取 TikTok 账号主页数据 API Response
type GetResponse struct {
	model.BaseResponse
	Data *Insight `json:"data,omitempty"`
}
