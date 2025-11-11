package product

import "github.com/bububa/tiktok-business/enum"

// IProduct 商品interface
type IProduct interface {
	CatalogType() enum.CatalogType
}

type Product struct {
	// Audit 商品审核信息
	Audit *Audit `json:"audit,omitempty"`
	// AdCreationEligible 本字段仅为电商商品库商品返回。
	// 该电商商品库商品是否可用于创建广告。
	// 枚举值：AVAILABLE，NOT_AVAILABLE。
	// 电商商品库商品仅在满足以下条件时ad_creation_eligible 为AVAILABLE：
	// audit_status 为 approved。
	// active_status 为 ACTIVATED。
	// availability 为 IN_STOCK，AVAILABLE_FOR_ORDER或PREORDER。
	AdCreationEligible enum.Availability `json:"ad_creation_eligible,omitempty"`
	// ActiveStatus 商品的上架状态。
	// 枚举值：
	// ACTIVATED：已上架。
	// DEACTIVATED：已下架。
	// 本字段代表广告主所控制的商品的上架或下架状态。商品上架并不一定意味着该商品可以用于创建广告。
	ActiveStatus enum.ProductActiveStatus `json:"active_status,omitempty"`
	// SkuID SKU ID，需要保证唯一性
	SkuID string `json:"sku_id,omitempty"`
	// ProductID 商品ID
	ProductID string `json:"product_id,omitempty"`
	// Title 商品标题
	Title string `json:"title,omitempty"`
	// Description 商品描述
	Description string `json:"description,omitempty"`
	// Availability 库存情况。
	// 对于电商商品库，枚举值为：
	// IN_STOCK：有库存。
	// AVAILABLE_FOR_ORDER：可订购。
	// PREORDER：可预订。
	// OUT_OF_STOCK：缺货。
	// DISCONTINUED：停售。
	// 对于酒店商品库，枚举值为：
	// IN_STOCK：可预定。
	// OUT_OF_STOCK：不可预定。
	// 对于娱乐商品库，枚举值为：
	// AVAILABLE
	// SUBSCRIBERS_ONLY
	// ON_DEMAND
	// NOT_AVAILABLE
	Availability enum.ProductAvailability `json:"availability,omitempty"`
	// ImageURL 商品主图的URL
	ImageURL string `json:"image_url,omitempty"`
	// VideoURL 广告视频链接
	VideoURL string `json:"video_url,omitempty"`
	// Brand 品牌信息
	Brand string `json:"brand,omitempty"`
	// AdditionalImageURLs 商品的附加图片的 URL 列表。
	AdditionalImageURLs []string `json:"additional_image_urls,omitempty"`
	// ItemGroupID 商品 SPU ID
	ItemGroupID string `json:"item_group_id,omitempty"`
	// GoogleProductCategory 品类信息
	GoogleProductCategory string `json:"google_product_category,omitempty"`
	// GlobalTradeItemNumber 全球贸易项目编号
	GlobalTradeItemNumber string `json:"global_trade_item_number,omitempty"`
	// ManufacturerPartNumber 制造商零件号
	ManufacturerPartNumber string `json:"manufacturer_part_number,omitempty"`
	// HotelID 酒店的唯一 ID
	HotelID string `json:"hotel_id,omitempty"`
	// Name 酒店的名称
	Name string `json:"name,omitempty"`
	// HotelCategory 酒店类别。
	// 枚举值：
	// HOTEL：普通酒店。
	// APARTMENT：公寓式酒店。
	// BOAT：船屋。
	// CAMPSITE ：露营地。
	// CAPSULE_HOTEL：胶囊酒店。
	// GUESTHOUSE：家庭旅馆。
	// HOMESTAY： 寄宿旅馆。
	// HOSTEL：青年旅社。
	// LUXURY_TENT：豪华帐篷酒店。
	// RESORT：度假酒店。
	// VILLA：别墅酒店。
	HotelCategory enum.HotelCategory `json:"hotel_category,omitempty"`
	// HotelRetailerID 酒店零售商的唯一 ID
	HotelRetailerID string `json:"hotel_retailer_id,omitempty"`
	// Address 地址信息
	Address *Address `json:"address,omitempty"`
	// Neighborhood 酒店或目的地所在社区。
	// 示例：Bangkok Old Town。
	Neighborhood string `json:"neighborhood,omitempty"`
	// Latitude 酒店或目的地的纬度。
	// 示例：37.484100
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude 酒店或目的地的经度。
	// 示例： -122.148252
	Longitude float64 `json:"longitude,omitempty"`
	// SeriesID 短剧的自定义唯一 ID。
	SeriesID string `json:"series_id,omitempty"`
	// SeriesName 短剧名称
	SeriesName string `json:"series_name,omitempty"`
	// Recharge 短剧收费详情
	Recharge *MiniSeriesRecharge `json:"recharge,omitempty"`
	// Phone 酒店电话号码
	Phone string `json:"phone,omitempty"`
	// MarginLevel 酒店按酒店房间的基础价格（price）上另外加收作为押金的比例。
	// 例如，10代表该酒店按酒店房间的基础价格的10%加收押金。
	MarginLevel int `json:"margin_level,omitempty"`
	// LoyaltyProgram 酒店提供的忠诚度计划
	LoyaltyProgram string `json:"loyalty_program,omitempty"`
	// CancellationPolicy 酒店的取消政策
	CancellationPolicy string `json:"cancellation_policy,omitempty"`
	// GuestRatings 酒店的顾客评分信息
	GuestRatings []GuestRating `json:"guest_ratings,omitempty"`
	// StarRating 酒店或媒体标题项的星级
	StarRating float64 `json:"star_rating,omitempty"`
	// HotelRoomID 酒店房间的唯一ID
	HotelRoomID string `json:"hotel_room_id,omitempty"`
	// RoomType 酒店房间类型
	RoomType string `json:"room_type,omitempty"`
	// MealPolicy 酒店房间对应的餐饮套餐
	MealPolicy string `json:"meal_policy,omitempty"`
	// Amenities 酒店的生活设施。
	// 示例：swimming pool, fitness center
	Amenities string `json:"amenities,omitempty"`
	// Priority 酒店或航班的优先级。
	// 取值范围：0-5。
	// 值越高，代表优先级越高。
	Priority int `json:"priority,omitempty"`
	// FlightID 航班的唯一 ID。
	// 此 ID 由系统自动生成
	FlightID string `json:"flight_id,omitempty"`
	// OriginAirport 该航班的出发地机场。格式为三位字母的国际航空运输协会机场代码（IATA代码）。
	// 示例：AAR。
	OriginAirport string `json:"origin_airport,omitempty"`
	// DestinationAirport 该航班的目的地机场。格式为三位字母的国际航空运输协会机场代码。
	// 示例： ABD。
	DestinationAirport string `json:"destination_airport,omitempty"`
	// OriginCity 出发地城市
	OriginCity string `json:"origin_city,omitempty"`
	// DestinationCity 目的地城市
	DestinationCity string `json:"destination_city,omitempty"`
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
	CabinClass enum.CabinClass `json:"cabin_class,omitempty"`
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
	AirlineCompany enum.AirlineCompany `json:"airline_company,omitempty"`
	// DestinationID 目的地的唯一 ID
	DestinationID string `json:"destination_id,omitempty"`
	// DestinationName 目的地的名称
	DestinationName string `json:"destination_name,omitempty"`
	// PostalCode 目的地的邮政编码
	PostalCode string `json:"postal_code,omitempty"`
	// Types 目的地的类型
	Types []string `json:"types,omitempty"`
	// MediaTitleID 媒体标题项的唯一ID。
	// 媒体标题项指可供发行或消费的媒体产品，例如电影、电视节目、音乐和体育比赛。
	MediaTitleID string `json:"media_title_id,omitempty"`
	// Timeline 媒体标题项的上线状态。
	// 枚举值：
	// COMING_SOON ：该媒体标题项即将上线，暂不支持购买或浏览。
	// ONLINE：该媒体标题项已上线，支持购买或浏览。
	// EXPIRE_SOON：该媒体标题项即将下线，将不支持支持购买或浏览。
	Timeline enum.EntertainmentProductTimeline `json:"timeline,omitempty"`
	// Category 媒体标题项的种类。
	// 枚举值：
	// MOVIE：电影。
	// MUSIC：音乐。
	// TV_SHOW：电视节目。
	// TV_SERIES：电视连续剧。
	// SPORTS_GAME：体育比赛。
	// LIVE_EVENT：直播活动。
	Category enum.EntertainmentProductCategory `json:"category,omitempty"`
	// Genres 媒体标题项的类型。
	// 若想获取本字段的允许值，请查看下文的“ genres可选值”小节。
	Genres []string `json:"genres,omitempty"`
	// QID 该媒体标题项在维基数据（Wikidata）上的唯一标识符 QID（或Q号码），由字母“Q”及其后的一个或多个数字组成
	QID string `json:"qid,omitempty"`
	// Profession 商品详情
	Profession *ProductDetail `json:"profession,omitempty"`
	// PriceInfo 价格信息
	PriceInfo *PriceInfo `json:"price_info,omitempty"`
	// LandingPage 落地页信息
	LandingPage *LandingPage `json:"landing_page,omitempty"`
	// ExtraInfo 其他信息
	ExtraInfo *ExtraInfo `json:"extra_info,omitempty"`
	// ImageStatus 商品图片的转存状态。枚举值:
	// PROCESSING: 图片正在被转存。
	// SUCCESS：图片转存成功。
	// FAIL：图片转存失败。
	// FILTERED：图片转存成功，但是因为尺寸不合尺寸要求（需要不小于500 x 500 像素)，无法使用。如果图片不满足要求，产品的审核状态会一直是processing。对于这种图片，你需要使用上传或更新产品接口上传一个满足条件的新图片。
	// NOT_SUPPORTED：图片格式不支持。
	// NO_FOUND：无法通过地址找到资源。请确保地址正确，文件可在网上公开访问。
	ImageStatus enum.ProductImageStatus `json:"image_status,omitempty"`
}

// ProductDetail 商品详情
type ProductDetail struct {
	EcomProductDetail
	MiniSeriesProductDetail
}

// Audit 商品审核信息
type Audit struct {
	// AuditStatus 审核状态。枚举值: approved（通过）, rejected（未通过）, processing（处理中）。
	AuditStatus enum.ProductAuditStatus `json:"audit_status,omitempty"`
	// RejectInfo 拒审原因列表。只有audit_status状态为rejected时返回
	RejectInfo []RejectInfo `json:"reject_info,omitempty"`
}

// RejectInfo 拒审原因
type RejectInfo struct {
	// Reason 拒审原因
	Reason string `json:"reason,omitempty"`
	// Suggestion 修改建议
	Suggestion string `json:"suggestion,omitempty"`
	// AffectedPlacement 受影响版位列表
	AffectedPlacement []enum.Placement `json:"affected_placement,omitempty"`
	// AffectedCountry 受影响国家及地区列表
	AffectedCountry []string `json:"affected_country,omitempty"`
}

var _ IProduct = (*Product)(nil)

// CatalogType implements IProduct interface
func (p Product) CatalogType() enum.CatalogType {
	if p.SkuID != "" {
		return enum.CatalogType_ECOM
	} else if p.HotelID != "" {
		return enum.CatalogType_HOTEL
	} else if p.FlightID != "" || p.OriginAirport != "" {
		return enum.CatalogType_FLIGHT
	} else if p.DestinationID != "" {
		return enum.CatalogType_DESTINATION
	} else if p.MediaTitleID != "" {
		return enum.CatalogType_ENTERTAINMENT
	} else if p.SeriesID != "" {
		return enum.CatalogType_MINI_SERIES
	}
	return ""
}
