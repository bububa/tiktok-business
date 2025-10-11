package image

import "github.com/bububa/tiktok-business/model"

// Image 图片素材
type Image struct {
	// ImageID 图片 ID，用于广告投放中创建广告
	ImageID string `json:"image_id,omitempty"`
	// MaterialID 素材 ID
	MaterialID string `json:"material_id,omitempty"`
	// IsCarouselUsable 该图片的来源是否适用于轮播广告。
	// 注意：
	// 由于图片数据写入存在时延，可能不返回本字段。在此情况下，建议您等待 10 秒后，将本接口返回的 image_id 传入 /file/image/ad/info/ 或 /file/image/ad/search/，查询图片的此项信息。
	// 若要在轮播广告中使用图片，需确保本字段的值为true，且图片满足创建轮播广告-步骤-创建一个广告中列出的规格要求。
	IsCarouselUsable bool `json:"is_carousel_usable,omitempty"`
	// Displayable 图片能否在平台中展示
	Displayable bool `json:"displayable,omitempty"`
	// Width 图片宽度
	Width int `json:"width,omitempty"`
	// Height 图片高度
	Height int `json:"height,omitempty"`
	// Format 图片格式
	Format string `json:"format,omitempty"`
	// URL 图片 URL，1小时有效期，过期后需重新获取
	URL string `json:"url,omitempty"`
	// ImageURL 图片 URL，1小时有效期，过期后需重新获取
	ImageURL string `json:"image_url,omitempty"`
	// Signature 图片MD5
	Signature string `json:"signature,omitempty"`
	// Size 图片大小，单位byte
	Size int64 `json:"size,omitempty"`
	// FileName 图片名称
	FileName string `json:"file_name,omitempty"`
	// CreateTime 创建时间。UTC 时间，格式：2020-06-10T07:39:14Z。
	// 注意：由于图片数据写入存在时延，可能不返回本字段。在此情况下，建议您等待 10 秒后，将本接口返回的 image_id 传入 /file/image/ad/info/ 或 /file/image/ad/search/，查询图片的此项信息。
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ModifyTime 修改时间。UTC 时间，格式：2020-06-10T07:39:14Z 。
	// 注意：由于图片数据写入存在时延，可能不返回本字段。在此情况下，建议您等待 10 秒后，将本接口返回的 image_id 传入 /file/image/ad/info/ 或 /file/image/ad/search/，查询图片的此项信息。
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
}
