package video

import "github.com/bububa/tiktok-business/util"

// AdUpdateRequest 更新视频名称 API Request
type AdUpdateRequest struct {
	// AdvertiserID 广告账号
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// FileName 视频名称。长度限制： 1-100 个字符。如果账户下存在相同名称的素材，那么将会在名称末尾自动增加毫秒级时间戳。如果最终长度超过 100 个字符，将会自动截断。
	FileName string `json:"file_name,omitempty"`
	// VideoID 视频 ID
	VideoID string `json:"video_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *AdUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
