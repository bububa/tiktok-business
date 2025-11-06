package enum

// Gender 受众性别
type Gender string

const (
	// FEMALE
	FEMALE Gender = "FEMALE"
	// MALE
	MALE Gender = "MALE"
	// UNISEX 中性
	UNISEX Gender = "UNISEX"
	// UnknownGender
	UnknownGender Gender = "NONE"
	// GenderFemale
	GenderFemale Gender = "Female"
	// GenderMale
	GenderMale Gender = "Male"
	// GenderOther
	GenderOther Gender = "Other"
)

// GENDER  性别
type GENDER string

const (
	// GENDER_FEMALE
	GENDER_FEMALE GENDER = "GENDER_FEMALE"
	// GENDER_MALE
	GENDER_MALE GENDER = "GENDER_MALE"
	// GENDER_UNLIMITED。
	GENDER_UNLIMITED GENDER = "GENDER_UNLIMITED"
)
