package product

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// EntertainmentProduct 娱乐商品
type EntertainmentProduct struct {
	// MediaTitleID 媒体标题项的唯一 ID。
	// 媒体标题项指可供发行或消费的媒体产品，例如电影、电视节目、音乐和体育比赛。
	// 长度限制：100 字符。
	// 示例：TIKTOKLIVE001。
	MediaTitleID string `json:"media_title_id,omitempty" csv:"media_title_id"`
	// Title 媒体标题项的名称。
	// 请只使用有效的 Unicode 字符，避免使用无效字符，如控制符、功能符或表情符号。
	// 长度限制：150 字符。
	// 示例：TikTok LIVE Series。
	Title string `json:"title,omitempty" csv:"title"`
	// Description 媒体标题项的简短描述。
	// 长度限制：10,000 字符。
	Description string `json:"description,omitempty" csv:"description"`
	// ImageURL 媒体标题项的主图片的 URL。
	// 图片规格必须大于等于500 x 500 像素，否则图片将被过滤，商品无法通过审核。详细信息请参考商品图片要求。
	// 所有图片应为 JPG 或 PNG 格式。
	// 示例：https://www.tiktok.com/entertainment_image_001.jpg。
	ImageURL string `json:"image_url,omitempty" csv:"image_link"`
	// AdditionalImageURLs 媒体标题项的附加图片的 URL。
	// 最大数量：20。
	// 图片规格必须大于等于500x500像素，否则图片将被过滤，商品无法通过审核。详细信息请参考商品图片要求。
	// 所有图片应为 JPG 或 PNG 格式。
	// 示例：["https://www.tiktok.com/entertainment_image_002.jpg","https://www.tiktok.com/entertainment_image_003.jpg"] 。
	AdditionalImageURLs model.CSVStringList `json:"additional_image_urls" csv:"additional_image_link"`
	// Genres 媒体标题项的类型。
	// 若想获取本字段的允许值，请查看下文的“ genres可选值”小节。
	// 最大数量：2。
	// 示例： TEEN。
	Genres model.CSVStringList `json:"genres,omitempty" csv:"genres"`
	// QID 该媒体标题项在维基数据（Wikidata）上的唯一标识符 QID（或Q号码），由字母“Q”及其后的一个或多个数字组成。
	// 长度限制：150 字符。
	// 示例： Q48938223。
	QID string `json:"qid,omitempty" csv:"qid"`
	// PriceInfo 价格信息
	PriceInfo *PriceInfo `json:"price_info,omitempty" csv:""`
	// LandingPage 落地页信息
	LandingPage *LandingPage `json:"landing_page,omitempty" csv:""`
	// ExtraInfo 其他信息
	ExtraInfo *ExtraInfo `json:"extra_info,omitempty" csv:""`
}

func (p EntertainmentProduct) CatalogType() enum.CatalogType {
	return enum.CatalogType_ENTERTAINMENT
}
