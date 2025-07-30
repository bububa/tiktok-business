package enum

// BudgetMode 预算类型。
type BudgetMode string

const (
	// BUDGET_MODE_INFINITE：不限预算。
	BUDGET_MODE_INFINITE BudgetMode = "BUDGET_MODE_INFINITE"
	// BUDGET_MODE_TOTAL：总预算。
	BUDGET_MODE_TOTAL BudgetMode = "BUDGET_MODE_TOTAL"
	// BUDGET_MODE_DAY：日预算。
	BUDGET_MODE_DAY BudgetMode = "BUDGET_MODE_DAY"
	// BUDGET_MODE_DYNAMIC_DAILY_BUDGET：动态日预算，即一周七日的平均预算。单日实际消耗不超过设定的七日平均预算的125%。周消耗不会超过七日平均预算*7。
	BUDGET_MODE_DYNAMIC_DAILY_BUDGET BudgetMode = "BUDGET_MODE_DYNAMIC_DAILY_BUDGET"
)
