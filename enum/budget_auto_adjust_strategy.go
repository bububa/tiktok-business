package enum

// BudgetAutoAdjustStrategy 日预算策略。
type BudgetAutoAdjustStrategy string

const (
	// BudgetAutoAdjustStrategy_AUTO_BUDGET_INCREASE：启用自动增加预算。当广告成效良好并满足目标 CPA、第 0 天目标 ROAS 和预算要求时，允许预算自动增加。
	// budget_auto_adjust_strategy 为 AUTO_BUDGET_INCREASE 时，指定的 budget 为初始日预算。当预算使用率达到 90% 或以上时，允许日预算自动增加 20%，每天最多可增加 10 次。日预算每天都将重置为初始日预算。
	BudgetAutoAdjustStrategy_AUTO_BUDGET_INCREASE BudgetAutoAdjustStrategy = "AUTO_BUDGET_INCREASE"
)
