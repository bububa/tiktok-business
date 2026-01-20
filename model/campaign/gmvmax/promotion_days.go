package gmvmax

// PromotionDays 促销日设置详情。
type PromotionDays struct {
	// IsEnabled 是否开启促销日设置，在高意向购物日自动增加预算并优化推广系列以获得更高的总收入。
	// 支持的值: true, false。
	// 默认值: false。
	// 若指定了promotion_days，本字段自动设置为true。
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// AutoScheduleEnabled 仅当 is_enabled 为 true 时生效。
	// 是否使用 TikTok Shop 促销日排期。
	// 支持的值: true, false。
	// 若 is_enabled 为 true，本字段自动设置为true。
	AutoScheduleEnabled *bool `json:"auto_schedule_enabled,omitempty"`
	// CustomScheduleList 自定义促销日期排期详情
	CustomScheduleList []CustomSchedule `json:"custom_schedule_list,omitempty"`
	// ROASBidMultiplier 目标投资回报率指数。
	// 目标投资回报率（ROI）的调整将在促销日生效，可为该推广系列带来更高的 GMV。
	// 枚举值:
	// 90：-10%。在促销日期间将目标投资回报率降低 10%。
	// 80：-20%。在促销日期间将目标投资回报率降低 20%。
	// 70：-30%。在促销日期间将目标投资回报率降低 30%。
	ROASBidMultiplier int `json:"roas_bid_multipiler,omitempty"`
	// AdjustedROASBid 促销日目标投资回报率。
	AdjustedROASBid float64 `json:"adjusted_roas_bid,omitempty"`
	// BudgetIncreasePercentage 促销日的预算增加百分比。
	// 例如，50 表示触发自动增加预算时，预算增加 50%。
	BudgetIncreasePercentage int `json:"budget_increase_percentage,omitempty"`
	// IncreaseLimit 促销日的预算增加次数上限。
	// 例如，10 表示触发自动增加预算时，最多允许每日增加 10 次。
	IncreaseLimit int `json:"increase_limit,omitempty"`
	// CurrentBudget 所指定的推广系列日预算（budget）
	CurrentBudget float64 `json:"current_budget,omitempty"`
	// NextIncrease 下次增加额，即与指定的推广系列日预算（budget）相比预算的增加额。
	// 计算公式：next_increase = current_budget x budget_increase_percentage。
	// 例如，若budget为 300 美元，budget_increase_percentage为 50，则next_increase为 150 美元，表示日预算增加 150 美元。
	NextIncrease float64 `json:"next_increase,omitempty"`
	// RemainedTimes 剩余的自动增加预算次数
	RemainedTimes int `json:"remained_times,omitempty"`
	// MaximumBudget 预算上限。
	// 计算公式：maximum_budget = current_budget + current_budget x budget_increase_percentage x remained_times。
	// 例如，若current_budget为 200 美元，budget_increase_percentage为 50，remained_times为 10（即预算最多还可增加 10 次），则maximum_budget为 1200 美元。
	MaximumBudget float64 `json:"maximum_budget,omitempty"`
	// EstimatedGrossRevenueIncrease 由促销日期间投资回报率的变动带来的预估总收入提升百分百。
	// 示例: 24%。
	EstimatedGrossRevenueIncrease string `json:"estimated_gross_revenue_increase,omitempty"`
}

// CustomSchedule 自定义促销日期排期详情
type CustomSchedule struct {
	// StartTime 自定义促销日期排期的开始日期，格式为 YYYY-MM-DD（广告账号时区）
	StartTime string `json:"start_time,omitempty"`
	// EndTime 自定义促销日期排期的结束日期，格式为 YYYY-MM-DD（广告账号时区）
	EndTime string `json:"end_time,omitempty"`
}
