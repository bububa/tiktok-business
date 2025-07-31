package enum

// PromotionTargetType 推广目标线索收集的专属推广对象类型。
type PromotionTargetType string

const (
	// PromotionTargetType_INSTANT_PAGE：即时表单（线索表单）。创建可快速加载的应用内 TikTok 即时表单来收集更多线索。
	PromotionTargetType_INSTANT_PAGE PromotionTargetType = "INSTANT_PAGE"
	// PromotionTargetType_EXTERNAL_WEBSITE：第三方网站表单。使用包含第三方网站表单的落地页，或使用会重定向至包含第三方网站表单网站的TikTok即时体验页面来收集更多线索。
	PromotionTargetType_EXTERNAL_WEBSITE PromotionTargetType = "EXTERNAL_WEBSITE"
	// PromotionTargetType_UNSET。
	PromotionTargetType_UNSET PromotionTargetType = "UNSET"
)
