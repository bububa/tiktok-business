package product

import "github.com/bububa/tiktok-business/enum"

// DestinationProduct 目的地商品
type DestinationProduct struct {
	// DestinationID 目的地的唯一 ID。
	// 长度限制：100 字符。
	DestinationID string `json:"destination_id,omitempty" csv:"destination_id"`
	// DestinationName 目的地的名称。
	// 长度限制：150 字符。
	DestinationName string `json:"destination_name,omitempty" csv:"destination_name"`
	// Description 目的地的简短描述。
	// 长度限制：10,000 字符。
	// 示例：A visit to one of the most famous sights in London.
	Description string `json:"description,omitempty" csv:"description"`
	// ImageURL 目的地的主图片的 URL。
	// 图片规格必须大于等于500x500像素，否则图片将被过滤，商品无法通过审核。详细信息请参考商品图片要求。
	// 所有图片应为 JPG 或 PNG 格式。
	// 示例： https://www.tiktok.com/hotel_room_image_001.jpg。
	ImageURL string `json:"image_url,omitempty" csv:"image_link"`
	// AdditionalImageURLs 目的地的附加图片的 URL。
	// 最大数量：10。
	// 图片规格必须大于等于500x500像素，否则图片将被过滤，商品无法通过审核。详细信息请参考商品图片要求。
	// 所有图片应为 JPG 或 PNG 格式。
	// 示例：["https://www.tiktok.com/destination_image_002.jpg","https://www.tiktok.com/destination_image_003.jpg"] 。
	AdditionalImageURLs CSVStringList `json:"additional_image_urls,omitempty" csv:"additional_image_link"`
	// VideoURL 广告所使用视频的 URL。
	// 建议宽高比：9:16（竖版）；
	// 建议用于 TikTok 的分辨率：≥720*1280；
	// 建议用于 TikTok 的比特率：≥516kbps；
	// 声音：已启用，包含字幕。
	// 示例： https://www.tiktok.com/destination.mp4 。
	VideoURL string `json:"video_url,omitempty" csv:"video_link"`
	// Address 地址信息
	Address *Address `json:"address,omitempty" csv:"address"`
	// Neighborhood 目的地所在社区。
	// 长度限制：100 字符。
	// 示例：Bangkok Old Town。
	Neighborhood string `json:"neighborhood,omitempty" csv:"neighborhood"`
	// PostalCode 目的地的邮政编码。
	PostalCode string `json:"postal_code,omitempty" csv:"post_code"`
	// Latitude 目的地的纬度。
	// 取值范围： [-90, 90]。
	// 示例：37.484100
	Latitude float64 `json:"latitude,omitempty" csv:"latitude"`
	// Longitude目的地的经度。
	// 取值范围：[-180, 180]。
	// 示例： -122.148252
	Longitude float64 `json:"longitude,omitempty" csv:"longitude"`
	// Types 目的地的类型。
	// 您可为目的地指定一个或多个类型。例如，对于目的地纽约市，您可以将本字段设置为 ["city","sightseeing"]。
	// 最大数量：20。
	Types CSVStringList `json:"types,omitempty" csv:"types"`
	// PriceInfo 价格信息
	PriceInfo *PriceInfo `json:"price_info,omitempty" csv:""`
	// LandingPage 落地页信息
	LandingPage *LandingPage `json:"landing_page,omitempty" csv:""`
	// ExtraInfo 其他信息
	ExtraInfo *ExtraInfo `json:"extra_info,omitempty" csv:""`
}

func (p DestinationProduct) CatalogType() enum.CatalogType {
	return enum.CatalogType_DESTINATION
}
