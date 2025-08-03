package enum

// DynamicDestination 动态目标页面策略。
type DynamicDestination string

const (
	// DynamicDestination_DLP：动态落地页。基于用户的主页、行为和意向，向每位用户展示不同的目标页面，包括自定义网站或即时体验商品页面等，从而最大限度地提高广告转化量。
	DynamicDestination_DLP DynamicDestination = "DLP"
	// DynamicDestination_UNSET：未设置。
	DynamicDestination_UNSET DynamicDestination = "UNSET"
)
