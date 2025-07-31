package enum

// HouseholdIncome 您想定向的收入群体。
type HouseholdIncome string

const (
	// HouseholdIncome_TOP5(家庭收入群体的前5%)
	HouseholdIncome_TOP5 HouseholdIncome = "TOP5"
	// HouseholdIncome_TOP10(家庭收入群体的前10%)
	HouseholdIncome_TOP10 HouseholdIncome = "TOP10"
	// HouseholdIncome_TOP10_25(家庭收入群体的前10-25%)
	HouseholdIncome_TOP10_25 HouseholdIncome = "TOP10_25"
	// HouseholdIncome_TOP25_50(家庭收入群体的前25-50%)
	HouseholdIncome_TOP25_50 HouseholdIncome = "TOP25_50"
)
