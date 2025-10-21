package product

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// EcomProduct 电商商品
type EcomProduct struct {
	// SkuID 商品的唯一 ID，例如 SKU。
	// 如果有相同 ID 的多个实例，这些实例将被忽略。如果要更新原始数据流中的同一商品，则需要使用相同的 ID 更换该商品。
	// 请只使用有效的 Unicode 字符，避免使用无效字符，如控制符或功能符。
	// 长度限制：100 字符，不能包含表情符号。
	// 示例：tiktok_item_1234。
	// 注意：不支持重复的 SKU ID。
	SkuID string `json:"sku_id,omitempty" csv:"sku_id"`
	// Title 商品标题。
	// 请只使用有效的 Unicode 字符，避免使用无效字符，如控制符、功能符或表情符号。
	// 请勿包含“免运费、高品质、批发价”等促销文案。
	// 长度限制：500 字符，不能包含表情符号。
	// 示例：T-Shirt (Unisex)。
	Title string `json:"title,omitempty" csv:"title"`
	// Description 关于商品的简短描述。
	// 长度限制：10,000 字符。
	// 示例：A vibrant crewneck for all shapes and sizes. Made from 100% cotton.
	Description string `json:"description,omitempty"`
	// Availability 商店中该商品的当前库存情况。
	// 枚举值：
	// IN_STOCK：有库存。
	// AVAILABLE_FOR_ORDER：可订购。
	// PREORDER：可预订。
	// OUT_OF_STOCK：缺货。
	// DISCONTINUED：停售。
	// 示例：IN_STOCK。
	Availability enum.ProductAvailability `json:"availability,omitempty" csv:"availability"`
	// Brand 商品品牌名称。
	// 长度限制：150 字符。
	// 示例：TikTok。
	Brand string `json:"brand,omitempty" csv:"brand"`
	// ImageURL 商品主图的 URL。
	// 图片规格必须大于等于500x500像素，否则图片将被过滤，商品无法通过审核。详细信息请参考商品图片要求。
	// 所有图片应为 JPG 或 PNG 格式。
	// 示例：https://www.tiktok.com/t_shirt_image_001.jpg。
	ImageURL string `json:"image_url,omitempty" csv:"image_link"`
	// AdditionalImageURLs 商品的附加图片的 URL 列表。
	// 最大数量：10。
	// 图片规格必须大于等于500x500像素，否则图片将被过滤，商品无法通过审核。详细信息请参考商品图片要求。
	// 所有图片应为 JPG 或 PNG 格式。
	// 示例：["https://www.tiktok.com/t_shirt_image_002.jpg","https://www.tiktok.com/t_shirt_image_003.jpg"] 。
	AdditionalImageURLs model.CSVStringList `json:"additional_image_urls,omitempty" csv:"additional_image_link"`
	// VideoURL 广告所使用视频的 URL。
	// 建议宽高比：9:16（竖版）；
	// 建议用于 TikTok 的分辨率：≥720*1280；
	// 建议用于 TikTok 的比特率：≥516kbps；
	// 声音：已启用，包含字幕。
	// 示例： https://www.tiktok.com/tiktok_t_shirt 。
	VideoURL string `json:"video_url,omitempty" csv:"video_link"`
	// ItemGroupID 商品 SPU ID。
	// 属于特定商品变化版本的商品。为所有变化版本的商品需提供相同的 item_group_id。例如，红色马球衫是马球衫的变化版本。
	// 示例：TikTok_1234_sample。
	ItemGroupID string `json:"item_group_id,omitempty" csv:"item_group_id"`
	// GoogleProductCategory Google 商品类目中的预设值。
	// 仅支持 3 个层级。
	// 请参考 Google 网站： https://support.google.com/merchants/answer/6324436?hl=en
	// 示例：Apparel & Accessories> Clothing> Shirts。
	GoogleProductCategory string `json:"google_product_category,omitempty" csv:"google_product_category"`
	// GlobalTradeItemNumber 商品的全球贸易项目代码。
	// 示例：3234567890126。
	GlobalTradeItemNumber string `json:"global_trade_item_number,omitempty" csv:"global_trade_item_number"`
	// ManufacturerPartNumber 商品的制造商部件号。
	// 示例：12433。
	ManufacturerPartNumber string `json:"manufacturer_part_number,omitempty" csv:"manufacturer_part_number"`
	// ProductDetail 商品的其他信息
	ProductDetail *EcomProductDetail `json:"product_detail,omitempty" csv:""`
	// PriceInfo 价格信息
	PriceInfo *PriceInfo `json:"price_info,omitempty" csv:""`
	// LandingPage 落地页信息
	LandingPage *LandingPage `json:"landing_page,omitempty" csv:""`
	// ExtraInfo 其他信息
	ExtraInfo *ExtraInfo `json:"extra_info,omitempty" csv:""`
}

func (p EcomProduct) CatalogType() enum.CatalogType {
	return enum.CatalogType_ECOM
}

// EcomProductDetail 商品的其他信息
type EcomProductDetail struct {
	// Condition 商店中商品的当前状况。
	//
	// 枚举值：
	// NEW：全新。
	// REFURBISHED：翻新。
	// USED：二手。
	//
	// 示例： NEW。
	Condition enum.ProductCondition `json:"condition,omitempty" csv:"condition"`
	// AgeGroup 商品所适用的年龄组。为了确保向适当年龄的目标人群提供商品，请准确填写此字段。
	// 枚举值：
	// NEW_BORN：新生儿（0-3 个月）。
	// INFANT：婴儿（3-12 个月）。
	// TODDLER：幼儿（1-5 岁）。
	// KIDS：儿童（5-13 岁）。
	// ADULT：成人（13 岁或以上）。
	// 示例： ADULT。
	AgeGroup enum.ProductAgeGroup `json:"age_group,omitempty" csv:"age_group"`
	// Color 商品的颜色。
	// 长度限制：100 字符。
	// 示例：blue。
	Color string `json:"color,omitempty" csv:"color"`
	// Gender 商品的适用性别。为了确保向适当性别的目标人群提供商品，请准确填写 此字段。
	// 枚举值：MALE（男性），FEMALE（女性），UNISEX（中性）。
	// 示例：UNISEX。
	Gender enum.Gender `json:"gender,omitempty" csv:"gender"`
	// Material 商品的材质，如棉、牛仔布或皮革。
	// 长度限制：200 字符。
	// 示例：cotton。
	Material string `json:"material,omitempty" csv:"material"`
	// Pattern 商品上的图案或印刷图形。
	// 长度限制：100 字符。
	// 示例：stripes1。
	Pattern string `json:"pattern,omitempty" csv:"pattern"`
	// ProductCategory 商品所属的类别。
	// 仅指明前 3 个层级。
	// 示例：Apparel & Accessories> Clothing> Shirts。
	ProductCategory string `json:"product_category,omitempty" csv:"product_type"`
	// Shipping 商品的运输类型。
	// 建议格式为：COUNTRY:STATE:SHIPPING_TYPE: PRICE（国家/地区:州/省:发货类型: 价格）。
	// 使用“;”分隔不同的区域。
	// 示例：US:CA:Ground:9.99 USD。
	Shipping string `json:"shipping,omitempty" csv:"shipping"`
	// ShippingWeight 商品重量。
	// 支持的单位：lb，oz，g，kg。
	// 示例：1 oz，1 kg。
	ShippingWeight string `json:"shipping_weight,omitempty" csv:"shipping_weight"`
	// Size 商品尺寸。
	// 示例： Small，Large，Extra Large。
	Size string `json:"size,omitempty" csv:"size"`
	// Tax 商品的税费
	Tax string `json:"tax,omitempty" csv:"tax"`
}

// PriceInfo 价格信息
type PriceInfo struct {
	// Price 商品的基础价格。
	// 示例：30。
	Price float64 `json:"price,omitempty" csv:"price"`
	// Currency 货币单位。
	// 若指定本字段，则需填入商品库所定向地域对应的币种。
	// 若未指定本字段，则默认设置为商品库所定向地域对应的币种。
	// 示例：USD。
	Currency string `json:"currency,omitempty" csv:"currency"`
	// SalePrice 商品的促销价格（如果正在折价销售）。
	// 示例：20。
	SalePrice float64 `json:"sale_price,omitempty" csv:"sale_price"`
	// SalePriceEffectiveDate 促销开始和结束的日期及时间，格式为：YYYY-MM-DDT23:59+00:00/YYYY-MM-DDT23:59+00:00。
	// 开始日期应在结束日期之前。
	// 如果商品的促销价没有限定期限，请不要传入本字段。
	// 示例：2023-12-11T23:59+00:00/2023-12-15T23:59+00:00。
	SalePriceEffectiveDate string `json:"sale_price_effective_date,omitempty" csv:"sale_price_effective_date"`
	// TotalPrice 仅对酒店商品库中的酒店返回本字段。
	// 总价格。
	TotalPrice float64 `json:"total_price,omitempty" csv:"total_price"`
}

// LandingPage 落地页信息
type LandingPage struct {
	// LandingPageURL 可浏览及购买该商品的落地页 URL。
	// 示例： https://www.tiktok.com/tiktok_t_shirt
	LandingPageURL string `json:"landing_page_url,omitempty" csv:"link"`
	// IosURL 作为 URL 的 iOS 应用自定义架构。对于 iOS，如果 iPhone 和 iPad 应用详情与普通 iOS 应用不同，请同时提供 iPhone 和 iPad 应用详情。
	// 示例：iOS://electronic
	IosURL string `json:"ios_url,omitempty" csv:"ios_url"`
	// IosAppStoreID iOS 应用的 App Store ID。
	// 示例：1234。
	IosAppStoreID string `json:"ios_app_store_id,omitempty" csv:"ios_app_store_id"`
	// IosAppName 要显示的 iOS 应用名称。
	// 示例：Electronic iOS。
	IosAppName string `json:"ios_app_name,omitempty" csv:"ios_app_name"`
	// IPhoneAppStoreID iPhone 应用的 App Store ID。
	// 示例：1234。
	IPhoneAppStoreID string `json:"iphone_app_store_id,omitempty" csv:"iphone_app_store_id"`
	// IPhoneAppName 要显示的 iPhone 应用名称。
	// 示例：Electronic iOS。
	IPhoneAppName string `json:"iphone_app_name,omitempty" csv:"iphone_app_name"`
	// IPadAppStoreID iPad 应用的 App Store ID。
	// 示例：1234。
	IPadAppStoreID string `json:"ipad_app_store_id,omitempty" csv:"ipad_app_store_id"`
	// IPadAppName 要显示的 iPad 应用名称。
	// 示例：Electronic iOS。
	IPadAppName string `json:"ipad_app_name,omitempty" csv:"ipad_app_name"`
	// AndroidURL 作为 URL 的 Android 应用自定义架构。
	// 示例：android://electronic。
	AndroidURL string `json:"android_url,omitempty" csv:"android_url"`
	// AndroidPackage Android 包的名称。
	// 示例：com.electronic。
	AndroidPackage string `json:"android_package,omitempty" csv:"android_package"`
	// AndroidAppName 要显示的 Android 应用名称。
	// 示例：Electronic Android。
	AndroidAppName string `json:"android_app_name,omitempty" csv:"android_app_name"`
}

// ExtraInfo 其他信息
type ExtraInfo struct {
	// CustomLabel0 您要添加的商品的其他相关信息。
	// 长度限制：100 字符。
	CustomLabel0 string `json:"custom_label_0,omitempty" csv:"custom_label_0"`
	// CustomLabel1 您要添加的商品的其他相关信息。
	// 长度限制：100 字符。
	CustomLabel1 string `json:"custom_label_1,omitempty" csv:"custom_label_1"`
	// CustomLabel2 您要添加的商品的其他相关信息。
	// 长度限制：100 字符。
	CustomLabel2 string `json:"custom_label_2,omitempty" csv:"custom_label_2"`
	// CustomLabel3 您要添加的商品的其他相关信息。
	// 长度限制：100 字符。
	CustomLabel3 string `json:"custom_label_3,omitempty" csv:"custom_label_3"`
	// CustomLabel4 您要添加的商品的其他相关信息。
	// 长度限制：100 字符。
	CustomLabel4 string `json:"custom_label_4,omitempty" csv:"custom_label_4"`
	// InternalLabel0 分配给该航班的内部标签。您可以使用这些内部标签在创建商品系列时筛选航班。这些标签仅对您可见。
	// 长度限制：100字符。
	// 如果已在创建商品系列时使用自定义标签（custom_label_0到custom_label_4）来筛选航班，建议改用内部标签。内部标签与自定义标签不同，可以根据需要添加或修改，而无需每次提交航班进行审核，避免广告投放的中断。
	InternalLabel0 string `json:"internal_label_0,omitempty" csv:"internal_label_0"`
	// InternalLabel1 分配给该航班的内部标签。您可以使用这些内部标签在创建商品系列时筛选航班。这些标签仅对您可见。
	// 长度限制：100字符。
	// 如果已在创建商品系列时使用自定义标签（custom_label_0到custom_label_4）来筛选航班，建议改用内部标签。内部标签与自定义标签不同，可以根据需要添加或修改，而无需每次提交航班进行审核，避免广告投放的中断。
	InternalLabel1 string `json:"internal_label_1,omitempty" csv:"internal_label_1"`
}
