package video

import "github.com/bububa/tiktok-business/model"

// Video 商品库视频
type Video struct {
	// CatalogVideoID 商品库视频的 ID。
	// 示例：1234567891234567891_1234567891234567891_12345678912345_12345678-a12b-34cd-ef56-7gh89i1j2k3l 。
	CatalogVideoID string `json:"catalog_video_id,omitempty" csv:"-"`
	// VideoName 视频名称
	VideoName string `json:"video_name,omitempty" csv:"video_name"`
	// VideoLink 视频文件的 URL
	VideoLink string `json:"video_link,omitempty" csv:"video_link"`
	// SkuIDList 与该视频绑定的电商商品库商品的 SKU ID 列表
	SkuIDList model.CSVStringList `json:"sku_id_list,omitempty" csv:"sku_id_list"`
	// Category 视频的商品类别
	Category string `json:"category,omitempty" csv:"category"`
	// Brand 视频的品牌名称
	Brand string `json:"brand,omitempty" csv:"brand"`
	// Creator 视频的创作者
	Creator string `json:"creator,omitempty" csv:"creator"`
	// VideoType 视频的类型
	VideoType string `json:"video_type,omitempty" csv:"video_type"`
	// Description 视频的简短描述
	Description string `json:"description,omitempty" csv:"description"`
	// LandingPageURL 可浏览及购买与该视频关联的商品的落地页 URL
	LandingPageURL string `json:"landing_page_url,omitempty" csv:"landing_page_url"`
	// CustomLabel0 视频的其他信息
	CustomLabel0 string `json:"custom_label_0,omitempty" csv:"custom_label_0"`
	// CustomLabel1 视频的其他信息
	CustomLabel1 string `json:"custom_label_1,omitempty" csv:"custom_label_1"`
	// CustomLabel2 视频的其他信息
	CustomLabel2 string `json:"custom_label_2,omitempty" csv:"custom_label_2"`
	// CustomLabel3 视频的其他信息
	CustomLabel3 string `json:"custom_label_3,omitempty" csv:"custom_label_3"`
	// CustomLabel4 视频的其他信息
	CustomLabel4 string `json:"custom_label_4,omitempty" csv:"custom_label_4"`
}
