package product

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// HotelProduct 酒店商品
type HotelProduct struct {
	// HotelID 酒店的唯一 ID。
	// 长度限制：100 字符。
	HotelID string `json:"hotel_id,omitempty" csv:"hotel_id"`
	// Name 酒店的名称。
	// 长度限制：150 字符。
	Name string `json:"name,omitempty" csv:"name"`
	// Description 酒店房间的简短描述。
	// 长度限制：10,000 字符。
	// 示例：Free Wifi, Free Breakfast, Free Cancellation
	Description string `json:"description,omitempty" csv:"description"`
	// Availability 酒店的可用性。
	// 枚举值：
	// IN_STOCK：可预定。
	// OUT_OF_STOCK：不可预定。
	// 示例：IN_STOCK。
	Availabliity enum.ProductAvailability `json:"availability,omitempty" csv:"availability"`
	// Brand 酒店品牌名称。
	// 长度限制：150 字符。
	Brand string `json:"brand,omitempty" csv:"brand"`
	// ImageURL 酒店房间的主图片的 URL。
	// 图片规格必须大于等于500x500像素，否则图片将被过滤，商品无法通过审核。详细信息请参考商品图片要求。
	// 所有图片应为 JPG 或 PNG 格式。
	// 示例： https://www.tiktok.com/hotel_room_image_001.jpg。
	ImageURL string `json:"image_url,omitempty" csv:"image_link"`
	// AdditionalImageURLs 酒店房间的附加图片的 URL。
	// 最大数量：20。
	// 图片规格必须大于等于500x500像素，否则图片将被过滤，商品无法通过审核。详细信息请参考商品图片要求。
	// 所有图片应为 JPG 或 PNG 格式。
	// 示例：["https://www.tiktok.com/hotel_room_image_002.jpg","https://www.tiktok.com/hotel_room_image_003.jpg"] 。
	AdditionalImageURLs model.CSVStringList `json:"additional_image_urls,omitempty" csv:"additional_image_link"`
	// VideoURL 广告所使用视频的 URL。
	// 建议宽高比：9:16（竖版）；
	// 建议用于 TikTok 的分辨率：≥720*1280；
	// 建议用于 TikTok 的比特率：≥516kbps；
	// 声音：已启用，包含字幕。
	// 示例： https://www.tiktok.com/hotel.mp4 。
	VideoURL string `json:"video_url,omitempty" csv:"video_link"`
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
	// 最大数量：3
	HotelCategory []enum.HotelCategory `json:"hotel_category,omitempty" csv:"category"`
	// HotelRetailerID 酒店零售商的唯一 ID。
	// 长度限制：100 字符。
	HotelRetailerID string `json:"hotel_retailer_id,omitempty" csv:"hotel_retailer_id"`
	// Address 地址信息
	Address *Address `json:"address,omitempty" csv:"address"`
	// Neighborhood 酒店所在社区。
	// 长度限制：100 字符。
	// 示例：Bangkok Old Town。
	Neighborhood string `json:"neighborhood,omitempty" csv:"neighborhood"`
	// PostalCode 酒店的邮政编码。
	// 长度限制：150 个字符。
	PostalCode string `json:"postal_code,omitempty" csv:"postal_code"`
	// Latitude 酒店所在地的纬度。
	// 取值范围： [-90, 90]。
	// 示例：37.484100。
	Latitude float64 `json:"latitude,omitempty" csv:"latitude"`
	// Longitude 酒店所在地的经度。
	// 取值范围：[-180, 180]。
	// 示例： -122.148252。
	Longitude float64 `json:"longitude,omitempty" csv:"longitude"`
	// MarginLevel 酒店按酒店房间的基础价格（price）上另外加收作为押金的比例。
	// 取值范围：[1, 10]。
	// 例如，10 代表该酒店按酒店房间的基础价格的10%加收押金。
	MarginLevel int `json:"margin_level,omitempty" csv:"margin_level"`
	// LoyaltyProgram 酒店提供的忠诚度计划。
	// 长度限制：150 字符。
	LoyaltyProgram string `json:"loyalty_program,omitempty" csv:"loyalty_program"`
	// GuestRatings 酒店的顾客评分信息。
	// 最大数量：5。
	GuestRatings []GuestRating `json:"guest_ratings,omitempty" csv:"guest_rating"`
	// StarRating 酒店的星级
	StarRating int `json:"star_rating,omitempty" csv:"star_rating"`
	// RoomType 酒店房间类型。
	// 长度限制：150 字符。
	// 最大数量：3
	RoomType []string `json:"room_type,omitempty" csv:"room_type"`
	// Priority 酒店的优先级。
	// 取值范围：0-5。
	// 默认值：0。
	// 值越高，代表优先级越高。
	Priority int `json:"priority,omitempty" csv:"priority"`
	// PriceInfo 价格信息
	PriceInfo *PriceInfo `json:"price_info,omitempty" csv:",inline"`
	// LandingPage 落地页信息
	LandingPage *LandingPage `json:"landing_page,omitempty" csv:",inline"`
	// ExtraInfo 其他信息
	ExtraInfo *ExtraInfo `json:"extra_info,omitempty" csv:",inline"`
}

func (p HotelProduct) CatalogType() enum.CatalogType {
	return enum.CatalogType_HOTEL
}

// Address 地址信息
type Address struct {
	// Address 酒店的具体地址。
	// 长度限制：150 字符。
	Address string `json:"address,omitempty" csv:"address"`
	// SecondaryAddress 酒店的次要街道地址。
	// 长度限制：150 个字符。
	SecondaryAddress string `json:"secondary_address,omitempty" csv:"secondary_address"`
	// TertiaryAddress 酒店的三级街道地址。
	// 长度限制：150 个字符。
	TertiaryAddress string `json:"tertiary_address,omitempty" csv:"tertiary_address"`
	// City 酒店所在城市。
	// 长度限制：150 字符。
	City string `json:"city,omitempty" csv:"city"`
	// Region 酒店所在州/省。
	// 长度限制：150 字符。
	Region string `json:"region,omitempty" csv:"region"`
	// Country 酒店所在国家或地区。
	// 长度限制：150 字符。
	Country string `json:"country,omitempty" csv:"country"`
}

// GuestRating 酒店的顾客评分信息。
type GuestRating struct {
	// RatingSystem 酒店的评分系统。
	// 长度限制：150 字符。
	// 示例：Expedia，Tripadvisor。
	RatingSystem string `json:"rating_system,omitempty" csv:"rating_system"`
	// MaxScore 该评分系统的最高评分。
	// 取值范围： [0,100]。
	MaxScore int `json:"max_score" csv:"max_score"`
	// Score 该酒店的实际评分。
	// 取值范围： ≥0。
	Score int `json:"score" csv:"score"`
	// NumberOfReviewers 评价过该酒店的人数。
	// 取值范围： ≥0。
	NumberOfReviewers int `json:"number_of_reviewers" csv:"number_of_reviewers"`
}
