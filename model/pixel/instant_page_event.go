package pixel

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InstantPageEventRequest 获取即时体验页面转化事件 API Request
type InstantPageEventRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// ObjectiveType 推广目标。枚举值请查看枚举值-推广目标
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// OptimizationGoal 优化目标
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// IsRetargeting 广告是否用于重定向
	IsRetargeting bool `json:"is_retargeting,omitempty"`
}

// Encode implements GetRequest interface
func (r *InstantPageEventRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("objective_type", string(r.ObjectiveType))
	values.Set("optimization_goal", string(r.OptimizationGoal))
	if r.IsRetargeting {
		values.Set("is_retargeting", "true")
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// InstantPageEventResponse 获取即时体验页面转化事件 API Response
type InstantPageEventResponse struct {
	model.BaseResponse
	Data struct {
		List []InstantPageEvent `json:"list,omitempty"`
	} `json:"data"`
}

// InstantPageEvent 即时体验页面转化事件
type InstantPageEvent struct {
	// BusinessType 业务类型。枚举值：TIKTOK_INSTANT_PAGE
	BusinessType string `json:"business_type,omitempty"`
	// InstantPageEvents 即时体验页面转化事件
	InstantPageEvents struct {
		// ObjectiveTypes 推广目标列表
		ObjectiveTypes []InstantPageEventObjectiveType `json:"objective_types,omitempty"`
	} `json:"instant_page_events"`
}

// InstantPageEventObjectiveType 推广目标
type InstantPageEventObjectiveType struct {
	// ObjectiveType 推广目标。枚举值请查看枚举值-推广目标。
	// 注意：若您选择推广目标WEB_CONVERSIONS作为筛选条件，目前会返回以CONVERSIONS为推广目标的转化事件数据。
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// OptimizationGoals 优化目标列表
	OptimizationGoals []InstantPageEventOptimizationGoal `json:"optimization_goals,omitempty"`
}

// InstantPageEventOptimizationGoal 优化目标
type InstantPageEventOptimizationGoal struct {
	// OptimizationGoal 优化目标
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// OptimizationEvents 转化事件列表。目前只支持BUTTON，该值代表点击按钮事件。
	OptimizationEvents []enum.OptimizationEvent `json:"optimization_events,omitempty"`
}
