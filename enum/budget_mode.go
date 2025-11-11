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

// AdvertiserBudgetMode 广告账户预算模式
type AdvertiserBudgetMode string

const (
	// AdvertiserBudgetMode_UNLIMITED：不限。广告账户无预算限制。
	AdvertiserBudgetMode_UNLIMITED AdvertiserBudgetMode = "UNLIMITED"
	// AdvertiserBudgetMode_MONTHLY_BUDGET：月预算。广告账号在月度预算内消耗信用额度。
	AdvertiserBudgetMode_MONTHLY_BUDGET AdvertiserBudgetMode = "MONTHLY_BUDGET"
	// AdvertiserBudgetMode_DAILY_BUDGET：日预算。广告账号在日预算内消耗信用额度。
	AdvertiserBudgetMode_DAILY_BUDGET AdvertiserBudgetMode = "DAILY_BUDGET"
	// AdvertiserBudgetMode_CUSTOM_BUDGET：自定义预算。广告账号在单次自定义预算内消耗信用额度。
	AdvertiserBudgetMode_CUSTOM_BUDGET AdvertiserBudgetMode = "CUSTOM_BUDGET"
)
