package enum

// ProductAgeGroup 商品所适用的年龄组
type ProductAgeGroup string

const (
	// ProductAgeGroup_NEW_BORN：新生儿（0-3 个月）。
	ProductAgeGroup_NEW_BORN ProductAgeGroup = "NEW_BORN"
	// ProductAgeGroup_INFANT：婴儿（3-12 个月）。
	ProductAgeGroup_INFANT ProductAgeGroup = "INFANT"
	// ProductAgeGroup_TODDLER：幼儿（1-5 岁）。
	ProductAgeGroup_TODDLER ProductAgeGroup = "TODDLER"
	// ProductAgeGroup_KIDS：儿童（5-13 岁）。
	ProductAgeGroup_KIDS ProductAgeGroup = "KIDS"
	// ProductAgeGroup_ADULT：成人（13 岁或以上）。
	ProductAgeGroup_ADULT ProductAgeGroup = "ADULT"
)
