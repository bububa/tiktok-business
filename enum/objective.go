package enum

// Objective 推广类型（应用或落地页），
type Objective string

const (
	// Objective_APP 应用
	Objective_APP Objective = "APP"
	// Objective_LANDING_PAGE 落地页
	Objective_LANDING_PAGE Objective = "LANDING_PAGE"
)
