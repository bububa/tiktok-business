package product

import "github.com/bububa/tiktok-business/enum"

// FlightProduct 航班商品
type FlightProduct struct {
	// FlightID 航班ID
	FlightID string `json:"flight_id,omitempty"`
	// OriginAirport 该航班的出发地机场。格式为三位字母的国际航空运输协会机场代码（IATA代码）。
	// 示例：AAR。
	// 注意: 同一个航班商品库中，不支持指定相同 origin_airport 和 destination_airport 的航班。
	OriginAirport string `json:"origin_airport,omitempty" csv:"origin_airport"`
	// DestinationAirport 该航班的目的地机场。格式为三位字母的国际航空运输协会机场代码。
	// 示例：ABD。
	// 注意: 同一个航班商品库中，不支持指定相同 origin_airport 和 destination_airport 的航班。
	DestinationAirport string `json:"destination_airport,omitempty" cvs:"destination_airport"`
	// OriginCity 出发地城市。
	// 长度限制：150 字符。
	OriginCity string `json:"origin_city,omitempty" csv:"origin_city"`
	// DestinationCity 目的地城市。
	// 长度限制：150 字符。
	DestinationCity string `json:"destination_city,omitempty" csv:"destination_city"`
	// Description 该航班的简短描述。
	// 长度限制：10,000 字符。
	Description string `json:"description,omitempty" csv:"description"`
	// CabinClass 该航班的客舱等级。
	// 枚举值：
	// FIRST_CLASS: 头等舱。享受奢华设施与私人空间。
	// FIRST_CLASS_SUITE: 头等舱套间。提供私密套间及顶级豪华体验。
	// BUSINESS_CLASS: 商务舱。为商务和休闲旅客提供舒适与高品质设施。
	// BUSINESS_CLASS_SUITE: 商务舱套间。宽敞舒适的私人空间。
	// COMFORT_CLASS: 舒适舱。介于高档经济舱与经济舱之间的升级服务。
	// ECONOMY_CLASS: 经济舱。标准旅行配置与基础设施，价格实惠。
	// BASIC_ECONOMY: 基本经济舱。提供部分基础设施，经济实惠。
	// STANDARD_ECONOMY: 标准经济舱。提供基本服务和合理定价。
	// PREMIUM_ECONOMY: 高级经济舱。比标准经济舱有更多腿部空间和更好的机上服务。
	// LIE_FLAT_SEAT: 平躺座椅。提供长途航班的舒适平躺体验。
	// CHARTER_CLASS: 包机舱。私人飞机租赁，享有独特的旅行体验。
	// ELITE_CLASS: 会员舱。融合头等与商务舱的顶级服务。
	// QUIET_CLASS: 静音舱。噪音减弱区，提供宁静的飞行体验。
	CabinClass enum.CabinClass `json:"cabin_class,omitempty" csv:"cabin_class"`
	// AirlineCompany 运营该航班的航空公司的名称。
	// 枚举值：
	// AEROFLOT: 俄罗斯航空。
	// AIR_FRANCE_KLM_GROUP: 法航-荷航集团。
	// ALASKA_AIRLINES: 阿拉斯加航空。
	// ANA: 全日空航空。
	// AMERICAN_AIRLINES: 美国航空。
	// BRITISH_AIRWAYS: 英国航空。
	// CATHAY_PACIFIC: 国泰航空。
	// CHINA_SOUTHERN_AIRLINES: 中国南方航空。
	// CHINA_EASTERN_AIRLINES: 中国东方航空。
	// DELTA_AIR_LINES: 达美航空。
	// EMIRATES: 阿联酋航空。
	// JAPAN_AIRLINES: 日本航空。
	// JETBLUE_AIRWAYS: 捷蓝航空。
	// KLM_ROYAL_DUTCH_AIRLINES: 荷兰皇家航空。
	// LUFTHANSA_GROUP: 汉莎航空集团。
	// QANTAS_AIRWAYS: 澳洲航空。
	// QATAR_AIRWAYS: 卡塔尔航空。
	// SINGAPORE_AIRLINES: 新加坡航空。
	// SOUTHWEST_AIRLINES: 西南航空。
	// TURKISH_AIRLINES: 土耳其航空。
	// UNITED_AIRLINES: 美国联合航空。
	// VIRGIN_ATLANTIC: 维珍航空。
	AirlineCompany enum.AirlineCompany `json:"airline_company,omitempty" csv:"airline_company"`
	// ImageURL 航班的主图片的 URL。
	// 图片规格必须大于等于500x500像素，否则图片将被过滤，商品无法通过审核。详细信息请参考商品图片要求。
	// 所有图片应为 JPG 或 PNG 格式。
	// 示例： https://www.tiktok.com/flight_image_001.jpg。
	ImageURL string `json:"image_url,omitempty" csv:"image_link"`
	// AdditionalImageURLs 航班的附加图片的 URL。
	// 最大数量：20。
	// 图片规格必须大于等于500x500像素，否则图片将被过滤，商品无法通过审核。详细信息请参考商品图片要求。
	// 所有图片应为 JPG 或 PNG 格式。
	// 示例：["https://www.tiktok.com/flight_image_002.jpg","https://www.tiktok.com/flight_image_003.jpg"] 。
	AdditionalImageURLs CSVStringList `json:"additional_image_urls,omitempty" csv:"additional_image_link"`
	// VideoURL 广告所使用视频的 URL。
	// 建议宽高比：9:16（竖版）；
	// 建议用于 TikTok 的分辨率：≥720*1280；
	// 建议用于 TikTok 的比特率：≥516kbps；
	// 声音：已启用，包含字幕。
	// 示例： https://www.tiktok.com/flight.mp4 。
	VideoURL string `json:"video_url,omitempty" csv:"video_link"`
	// Priority 航班的优先级。
	// 取值范围：0-5。
	// 默认值：0。
	// 值越高，代表优先级越高。
	Priority int `json:"priority" csv:"priority"`
	// PriceInfo 价格信息
	PriceInfo *PriceInfo `json:"price_info,omitempty" csv:""`
	// LanddingPage 落地页信息
	LandingPage *LandingPage `json:"landing_page,omitempty" csv:""`
	// ExtraInfo 其他信息
	ExtraInfo *ExtraInfo `json:"extra_info,omitempty" csv:""`
}

func (p FlightProduct) CatalogType() enum.CatalogType {
	return enum.CatalogType_FLIGHT
}
