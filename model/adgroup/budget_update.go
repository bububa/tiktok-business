package adgroup

import "github.com/bububa/tiktok-business/util"

// BudgetUpdateRequest 更新广告组预算 API Request
type BudgetUpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Budget
	Budget []Budget `json:"budget,omitempty"`
	// ScheduleBudget
	ScheduleBudget []ScheduleBudget `json:"schedule_budget,omitempty"`
}

// 想要为广告组更新的总预算的有关信息
type Budget struct {
	// AdgroupID 传入对象数组 budget 时，本字段必填。
	// 广告组 ID。
	AdgroupID string `json:"adgroup_id,omitempty"`
	// Budget 传入对象数组 budget 时，本字段必填。
	// 想要为广告组（adgroup_id）更新的总预算
	Budget float64 `json:"budget,omitempty"`
}

// ScheduleBudget 想要为广告组日预算或动态日预算设置的预定预算更改（定时预算）的有关信息
type ScheduleBudget struct {
	// AdgroupID 传入对象数组 schedule_budget 时，本字段必填。
	// 广告组 ID。
	AdgroupID string `json:"adgroup_id,omitempty"`
	// ScheduleBudget 传入对象数组 scheduled_budget 时，本字段必填。
	// 想要为广告组（adgroup_id）更新的日预算或动态日预算。
	ScheduleBudget float64 `json:"schedule_budget,omitempty"`
}

// Encode implements PostRequest interface
func (r BudgetUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
