package gmvmax

// AutoBudget 非促销日的自动增加预算设置详情
type AutoBudget struct {
	// AutoBudgetEnabled 是否在非促销日启用自动增加预算。
	// 启用自动增加预算即在您的推广系列达到 ROI 目标的 90% 以上且预算使用率不低于 80% 时，允许系统自动增加日预算，以优化推广系列，提升销量。日预算每天都将重置为原始金额。
	// 支持的值：true，false。
	AutoBudgetEnabled *bool `json:"auto_budget_enabled,omitempty"`
	// BudgetIncreasePercentage 非促销日的预算增加百分比。
	// 例如，50 表示触发自动增加预算时，预算增加 50%
	BudgetIncreasePercentage int `json:"budget_increase_percentage,omitempty"`
	// IncreaseLimit 非促销日的预算增加次数上限。
	// 例如，10 表示触发自动增加预算时，最多允许每日增加 10 次。
	IncreaseLimit int `json:"increase_limit,omitempty"`
	// CurrentBudget 所指定的推广系列日预算（budget）
	CurrentBudget float64 `json:"current_budget,omitempty"`
	// NextIncrease 下次增加额，即与指定的推广系列日预算（budget）相比预算的增加额。
	// 计算公式：next_increase = current_budget x budget_increase_percentage。
	// 例如，若budget为 300 美元，budget_increase_percentage为50，则next_increase为 150 美元，表示日预算增加 150 美元。
	NextIncrease float64 `json:"next_increase,omitempty"`
	// RemainedTimes 剩余的自动增加预算次数。
	RemainedTimes int `json:"remained_times,omitempty"`
	// MaximumBudget 预算上限。
	// 计算公式：maximum_budget = current_budget + current_budget x budget_increase_percentage x remained_times。
	// 例如，若current_budget为 200 美元，budget_increase_percentage为 50，remained_times为 10（即预算最多还可增加 10 次），则maximum_budget为 1200 美元。
	MaximumBudget float64 `json:"maximum_budget,omitempty"`
}
