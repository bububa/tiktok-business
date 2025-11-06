package creator

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AuthorizedRequest 授权后获取账户洞察 API Request
type AuthorizedRequest struct {
	// CreatorID TikTok 创作者的公开ID
	CreatorID string `json:"creator_id,omitempty"`
	// Fields 请求的字段。
	// 若未设置，将返回默认字段（profile_image 和 display_name）。
	// 若想返回指定的字段，需要有对应的字段权限。您可根据返回结构中每个字段对应的 “所需权限” 列，查看所需的权限。
	// 若想确认您已具备对应的权限范围，您可查看/tt_user/token_info/get/接口返回的权限范围。若您还未获得所需的权限，则需要遵循获取创作者授权（新）中列出的步骤，使 TikTok 账户所有者授予您的 App 对应的权限。
	// 枚举值：
	// profile_image
	// display_name
	// followers_count
	// handle_name
	// bio
	// is_on_vacation
	// vacation_start_date
	// vacation_end_date
	// creator_rate
	// location
	// topics
	// creator_badges
	// audience_countries
	// audience_genders
	// audience_ages
	// audience_locales
	// audience_device
	// audience_usage
	// average_views
	// average_views_benchmark
	// six_seconds_views_rate
	// six_seconds_views_rate_benchmark
	// engagement_rate
	// engagement_rate_benchmark
	// engagements_benchmark
	// completion_rate
	// completion_rate_benchmark
	// followers_growth_rate
	// followers_growth_rate_benchmark
	// average_likes
	// average_shares
	// average_comments
	// sponsored_video
	Fields []string `json:"fields,omitempty"`
}

// Encode implements GetRequest
func (r *AuthorizedRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("creator_id", r.CreatorID)
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AuthorizedResponse 授权后获取账户洞察 API Response
type AuthorizedResponse struct {
	model.BaseResponse
	Data *Creator `json:"data,omitempty"`
}
