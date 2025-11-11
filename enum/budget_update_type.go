package enum

// BudgetUpdateType 预算执行的更改类型。
type BudgetUpdateType string

const (
	// BudgetUpdateType_UPDATE：更新预算模式和预算金额。
	BudgetUpdateType_UPDATE BudgetUpdateType = "UPDATE"
	// 若将本字段设置为 UPDATE，需传入advertiser_budgets 对象中的 advertiser_id 和 budget，budget_mode 为可选字段。
	// BudgetUpdateType_RESET：将广告账号的预算消耗重置为 0。
	BudgetUpdateType_RESET BudgetUpdateType = "RESET"
	// 仅支持对预算模式为自定义（budget_mode 为 CUSTOM_BUDGET）的广告账号执行重置预算消耗操作。
	// 若将本字段设置为 RESET，仅需传入advertiser_budgets 对象中的 advertiser_id。若传入了 budget_mode 和 budget 字段，将忽略budget_mode 和 budget。
	// BudgetUpdateType_INCREMENTAL_UPDATE：增加或减少当前广告账户预算。
	BudgetUpdateType_INCREMENTAL_UPDATE BudgetUpdateType = "INCREMENTAL_UPDATE"
	// BudgetUpdateType_ONE_CLICK_SET：将广告账户预算一键设置为最小值。此时无需传入对象数组 advertiser_budgets 中的 budget 字段。您可通过返回参数 one_click_set_amount 获取更改后的预算。
	BudgetUpdateType_ONE_CLICK_SET BudgetUpdateType = "ONE_CLICK_SET"
)
