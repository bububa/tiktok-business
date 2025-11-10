package enum

// UserFinanceRole 用户财务角色
type UserFinanceRole string

const (
	// UserFinanceRole_MANAGER(财务经理)
	UserFinanceRole_MANAGER UserFinanceRole = "MANAGER"
	// UserFinanceRole_ANALYST(财务分析师)
	UserFinanceRole_ANALYST UserFinanceRole = "ANALYST"
)
