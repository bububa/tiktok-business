package enum

// CabinClass 该航班的客舱等级。
type CabinClass string

const (
	// CabinClass_FIRST_CLASS: 头等舱。享受奢华设施与私人空间。
	CabinClass_FIRST_CLASS CabinClass = "FIRST_CLASS"
	// CabinClass_FIRST_CLASS_SUITE: 头等舱套间。提供私密套间及顶级豪华体验。
	CabinClass_FIRST_CLASS_SUITE CabinClass = "FIRST_CLASS_SUITE"
	// CabinClass_BUSINESS_CLASS: 商务舱。为商务和休闲旅客提供舒适与高品质设施。
	CabinClass_BUSINESS_CLASS CabinClass = "BUSINESS_CLASS"
	// CabinClass_BUSINESS_CLASS_SUITE: 商务舱套间。宽敞舒适的私人空间。
	CabinClass_BUSINESS_CLASS_SUITE CabinClass = "BUSINESS_CLASS_SUITE"
	// CabinClass_COMFORT_CLASS: 舒适舱。介于高档经济舱与经济舱之间的升级服务。
	CabinClass_COMFORT_CLASS CabinClass = "COMFORT_CLASS"
	// CabinClass_ECONOMY_CLASS: 经济舱。标准旅行配置与基础设施，价格实惠。
	CabinClass_ECONOMY_CLASS CabinClass = "ECONOMY_CLASS"
	// CabinClass_BASIC_ECONOMY: 基本经济舱。提供部分基础设施，经济实惠。
	CabinClass_BASIC_ECONOMY CabinClass = "BASIC_ECONOMY"
	// CabinClass_STANDARD_ECONOMY: 标准经济舱。提供基本服务和合理定价。
	CabinClass_STANDARD_ECONOMY CabinClass = "STANDARD_ECONOMY"
	// CabinClass_PREMIUM_ECONOMY: 高级经济舱。比标准经济舱有更多腿部空间和更好的机上服务。
	CabinClass_PREMIUM_ECONOMY CabinClass = "PREMIUM_ECONOMY"
	// CabinClass_LIE_FLAT_SEAT: 平躺座椅。提供长途航班的舒适平躺体验。
	CabinClass_LIE_FLAT_SEAT CabinClass = "LIE_FLAT_SEAT"
	// CabinClass_CHARTER_CLASS: 包机舱。私人飞机租赁，享有独特的旅行体验。
	CabinClass_CHARTER_CLASS CabinClass = "CHARTER_CLASS"
	// CabinClass_ELITE_CLASS: 会员舱。融合头等与商务舱的顶级服务。
	CabinClass_ELITE_CLASS CabinClass = "ELITE_CLASS"
	// CabinClass_QUIET_CLASS: 静音舱。噪音减弱区，提供宁静的飞行体验。
	CabinClass_QUIET_CLASS CabinClass = "QUIET_CLASS"
)
