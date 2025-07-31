package enum

// SpecialIndustry 特殊广告分类。
type SpecialIndustry string

const (
	// SpecialIndustry_HOUSING：房地产，房屋保险，抵押贷款或相关的广告。
	SpecialIndustry_HOUSING SpecialIndustry = "HOUSING"
	// SpecialIndustry_EMPLOYMENT：工作机会，实习机会，职业认证项目或相关的广告。
	SpecialIndustry_EMPLOYMENT SpecialIndustry = "EMPLOYMENT"
	// SpecialIndustry_CREDIT：信用卡申请，汽车贷款，长期融资或相关的广告。
	SpecialIndustry_CREDIT SpecialIndustry = "CREDIT"
)
