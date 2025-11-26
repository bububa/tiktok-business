package audienceinsight

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// OverlapRequest 获取受众重叠详情 API Request
type OverlapRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// BenchmarkCustomAudienceID 选择作为基准受众的自定义受众ID。
	// 注意：这里输入的受众ID必须能通过您的广告主ID访问，否则将会报错。
	BenchmarkCustomAudienceID string `json:"benchmark_custom_audience_id,omitempty"`
	// ComparisonCustomAudienceIDs 选择作为对比受众的自定义受众ID。
	// 最大数量：4
	// 注意：这里输入的受众ID必须能通过您的广告主ID访问，否则将会报错
	ComparisonCustomAudienceIDs []string `json:"comparison_custom_audience_ids,omitempty"`
}

// Encode implements GetRequest
func (r *OverlapRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("benchmark_custom_audience_id", r.BenchmarkCustomAudienceID)
	if len(r.ComparisonCustomAudienceIDs) > 0 {
		values.Set("comparison_custom_audience_ids", string(util.JSONMarshal(r.ComparisonCustomAudienceIDs)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// OverlapResponse 获取受众重叠详情 API Response
type OverlapResponse struct {
	model.BaseResponse
	Data *Overlap `json:"data,omitempty"`
}

type Overlap struct {
	// BenchmarkAudience 基准受众数据
	BenchmarkAudience *BenchmarkAudience `json:"benchmark_audience,omitempty"`
	// ComparisonAudiences 对比受众数据。
	// 如果您在请求中输入了comparison_custom_audience_ids，将会返回该字段。
	ComparisonAudiences []ComparisonAudience `json:"comparison_audiences,omitempty"`
}

// BenchmarkAudience 基准受众数据
type BenchmarkAudience struct {
	// AudienceID 受众ID
	AudienceID string `json:"audience_id,omitempty"`
	// AudienceName 受众名称
	AudienceName string `json:"audience_name,omitempty"`
	// AudienceSize 受众规模
	AudienceSize int64 `json:"audience_size,omitempty"`
	// TargetableUsersCountRange 在自定义受众中，近 30 天内处于活跃状态且可被定向的用户的数量范围
	TargetableUsersCountRange string `json:"targetable_users_count_range,omitempty"`
}

// ComparisonAudience 对比受众数据
type ComparisonAudience struct {
	BenchmarkAudience
	// BenchmarkOverlapRate 基准受众中可用于对比受众的用户的基础百分比
	BenchmarkOverlapRate model.Float64 `json:"benchmark_overlap_rate,omitempty"`
	// BenchmarkOverlapRateRange 基准受众中可用于对比受众的用户的比例范围
	BenchmarkOverlapRateRange string `json:"benchmark_overlap_rate_range,omitempty"`
	// BenchmarkOverlapCountRange 基准受众中可用于对比受众的用户的数量范围
	BenchmarkOverlapCountRange string `json:"benchmark_overlap_count_range,omitempty"`
}
