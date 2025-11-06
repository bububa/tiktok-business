package ttvideo

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InfoRequest 获取 Spark Ads 帖子信息 API Request
type InfoRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AuthCode Spark Ads帖子授权码
	AuthCode string `json:"auth_code,omitempty"`
}

// Encode implements GetRequest
func (r *InfoRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("auth_code", r.AuthCode)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// InfoResponse 获取 Spark Ads 帖子信息 API Response
type InfoResponse struct {
	model.BaseResponse
	Data *Item `json:"data,omitempty"`
}

// Item Spark Ads 帖子信息
type Item struct {
	// AuthInfo 授权信息
	AuthInfo *AuthInfo `json:"auth_info,omitempty"`
	// ItemInfo Spark Ads 帖子信息
	ItemInfo *ItemInfo `json:"item_info,omitempty"`
	// UserInfo TikTok 账户信息
	UserInfo *UserInfo `json:"user_info,omitempty"`
	// VideoInfo 视频帖子信息。
	// 注意：item_type 为 CAROUSEL 时，video_info 对象中的字段值为 null。
	VideoInfo *VideoInfo `json:"video_info,omitempty"`
}

// AuthInfo 授权信息
type AuthInfo struct {
	// AuthEndTime 授权码有效期结束的时间(UTC+0)。时间格式：2017-01-01 00:00:00
	AuthEndTime model.DateTime `json:"auth_end_time,omitzero"`
	// AuthStartTime 授权码有效期开始的时间(UTC+0)。时间格式：2017-01-01 00:00:00
	AuthStartTime model.DateTime `json:"auth_start_time,omitzero"`
}

// ItemInfo Spark Ads 帖子信息
type ItemInfo struct {
	// AuthCode Spark Ads 帖子的授权码
	AuthCode string `json:"auth_code,omitempty"`
	// ItemID Spark Ads 帖子的ID
	ItemID string `json:"item_id,omitempty"`
	// ItemType Spark Ads 帖子类型。
	// 枚举值：
	// VIDEO：视频帖子。
	// CAROUSEL：照片帖子
	ItemType string `json:"item_type,omitempty"`
	// Text Spark Ads 帖子描述
	Text string `json:"text,omitempty"`
	// Status Spark Ads 帖子状态。
	// 枚举值参见枚举值 - TikTok 帖子状态
	Status enum.TTItemStatus `json:"status,omitempty"`
	// CarouselInfo 照片帖子信息。
	// 注意：item_type 为 VIDEO 时，本字段值为 null
	CarouselInfo *CarouselInfo `json:"carousel_info,omitempty"`
	// StitchOriginalItemID 如果本帖子是另一帖子的拼接，本字段代表源帖子的视频ID
	StitchOriginalItemID string `json:"stitch_original_item_id,omitempty"`
	// DuetOriginalItemID 如果本帖子是另一帖子的合拍，本字段代表该源帖子的视频ID
	DuetOriginalItemID string `json:"duet_original_item_id,omitempty"`
	// IsMultiDuetStitch 本拼接帖是否是另一拼接帖的拼接，或本合拍帖是否是另一合拍帖的合拍
	IsMultiDuetStitch bool `json:"is_multi_duet_stitch,omitempty"`
	// MentionedItemIDs 如果本帖子提及了其他视频，本字段代表该源帖子的视频ID。
	// 注意：帖子只能提及一个视频
	MentionedItemIDs []string `json:"mentioned_item_ids,omitempty"`
	// AnchorList Spark Ads 帖子中的锚点列表
	AnchorList []Anchor `json:"anchor_list,omitempty"`
}

// CarouselInfo 照片帖子信息。
type CarouselInfo struct {
	// ImageInfo 照片帖子中使用的图片的信息。
	// 图片返回的顺序与在照片帖子中的顺序一致。
	ImageInfo []ImageInfo `json:"image_info,omitempty"`
	// MusicInfo照片帖子中使用的音乐的信息。
	// 注意：若照片帖子是从 TikTok 应用中直接发布且无音乐的帖子，则本字段的值将为 null 。
	MusicInfo []MusicInfo `json:"music_info,omitempty"`
}

// ImageInfo 照片帖子中使用的图片的信息。
type ImageInfo struct {
	// ImageURL 图片 URL。
	// 有效期：90 天。
	ImageURL string `json:"image_url,omitempty"`
	// ImageHeight 图片高度，单位为像素
	ImageHeight int `json:"image_height,omitempty"`
	// ImageWidth 图片宽度，单位为像素
	ImageWidth int `json:"image_width,omitempty"`
}

// MusicInfo 照片帖子中使用的音乐的信息。
type MusicInfo struct {
	// MusicURL 音乐 URL。
	// 有效期：90 天。
	MusicURL string `json:"music_url,omitempty"`
	// MusicDuration 音乐时长，单位为秒
	MusicDuration int `json:"music_duration,omitempty"`
}

// Anchor Spark Ads 帖子中的锚点
type Anchor struct {
	// ID 锚点 ID
	ID string `json:"id,omitempty"`
	// Title 锚点标题
	Title string `json:"title,omitempty"`
	// Status 锚点状态。枚举值: CHECK_ING (锚点正在审核中), CHECK_FAILED (锚点审核未通过), CHECK_SUCCESS (锚点审核已通过)
	Status string `json:"status,omitempty"`
	// ProductRegions 广告可投放区域
	ProductRegions []string `json:"product_regions,omitempty"`
	// URL 视频 URL
	URL string `json:"url,omitempty"`
	// SpuID 仅当帖子中包含商品锚点链接时返回本字段。
	// 该帖子绑定的商品的 SPU ID
	SpuID string `json:"spu_id,omitempty"`
	// SpuName 仅当帖子中包含商品锚点链接时返回本字段。
	// 该帖子绑定的商品的名称。
	SpuName string `json:"spu_name,omitempty"`
	// StoreID 仅当帖子中包含商品锚点链接时返回本字段。
	// 商品所属的 TikTok Shop 的 ID。
	StoreID string `json:"store_id,omitempty"`
}

// UserInfo TikTok 账户信息
type UserInfo struct {
	// TikTokName TikTok 账户昵称
	TikTokName string `json:"tiktok_name,omitempty"`
	// IdentityID 认证身份ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
}

// VideoInfo 视频帖子信息。
type VideoInfo struct {
	// BidRate 视频比特率，单位为位/秒
	BidRate int64 `json:"bit_rate,omitempty"`
	// Duration 视频时长，单位秒
	Duration float64 `json:"duration,omitempty"`
	// Size 视频大小，单位byte
	Size int64 `json:"size,omitempty"`
	// Width 视频宽度
	Width int `json:"width,omitempty"`
	// Height 视频高度
	Height int `json:"height,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// FileName 视频名称
	FileName string `json:"file_name,omitempty"`
	// PosterURL 视频封面 URL。有效期1小时，过期需重新获取
	PosterURL string `json:"poster_url,omitempty"`
	// PreviewURL 视频预览链接, 有效期1小时，过期需重新获取
	PreviewURL string `json:"preview_url,omitempty"`
	// Signature 视频文件MD5
	Signature string `json:"signature,omitempty"`
}
