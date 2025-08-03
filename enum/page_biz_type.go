package enum

// PageBizType 落地页业务类型。
type PageBizType string

const (
	// PageBizType_CUSTOM（定制）
	PageBizType_CUSTOM PageBizType = "CUSTOM"
	// PageBizType_APP_PROFILE（应用信息）
	PageBizType_APP_PROFILE PageBizType = "APP_PROFILE"
	// PageBizType_INSTANT_FORM（即时表单）
	PageBizType_INSTANT_FORM PageBizType = "INSTANT_FORM"
)
