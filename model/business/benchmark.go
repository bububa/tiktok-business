package business

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// BenchmarkRequest 获取企业类别的基准数据 API Request
type BenchmarkRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// BusinessCategory 企业类别。
	// 枚举值参见business_category 可选值。
	// 示例：ART_AND_CRAFTS。
	BusinessCategory enum.BusinessCategory `json:"business_category,omitempty"`
}

// Encode implements GetRequest
func (r *BenchmarkRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	values.Set("business_category", string(r.BusinessCategory))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// BenchmarkResponse 获取企业类别的基准数据 API Response
type BenchmarkResponse struct {
	model.BaseResponse
	Data *Benchmark `json:"data,omitempty"`
}

// Benchmark 企业类别的基准数据
type Benchmark struct {
	// BusinessCategory 企业类别。
	// 枚举值参见business_category 可选值。
	// 示例：ART_AND_CRAFTS。
	BusinessCategory enum.BusinessCategory `json:"business_category,omitempty"`
	// AverageLikes 所有属于此企业类别的企业号中自有视频的平均点赞数。
	// 示例：22618.7106
	AverageLikes float64 `json:"average_likes,omitempty"`
	// AverageComments 所有属于此企业类别的企业号中自有视频的平均评论数。
	// 示例：625.6003
	AverageComments float64 `json:"average_comments,omitempty"`
	// AverageShares 所有属于此企业类别的企业号中自有视频的平均分享数。
	AverageShares float64 `json:"average_shares,omitempty"`
	// AverageVideoCount 所有属于此企业类别的企业号的平均视频数量
	AverageVideoCount float64 `json:"average_video_count,omitempty"`
	// AverageFollowerCount 所有属于此企业类别的企业号的平均粉丝数
	AverageFollowerCount float64 `json:"average_follower_count,omitempty"`
	// AverageFollowerGrowth 所有属于此企业类别的企业号过去 30 天的平均涨粉数
	AverageFollowerGrowth float64 `json:"average_follower_growth,omitempty"`
	// AverageEngagementRate 所有属于此企业类别的企业号的平均互动率
	AverageEngagementRate float64 `json:"average_engagement_rate,omitempty"`
	// AverageVideoViews 所有属于此企业类别的企业号的平均视频观看次数
	AverageVideoViews float64 `json:"average_video_views,omitempty"`
}
