package pixel

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建Pixel API Request
type CreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// PixelName Pixel 名称。
	// 名称不能超过40字符（不区分语言），不能包含emoji。
	PixelName string `json:"pixel_name,omitempty"`
	// PixelCategory Pixel 追踪事件类型。
	// 注意：无论您是否传入本字段，都不再影响 Pixel 可以追踪的 Pixel 事件。建议新建 Pixel 时不要传入本字段。本字段将在下一个 API 版本中废弃。
	PixelCategory enum.PixelCategory `json:"pixel_category,omitempty"`
}

// Encode implements PostRequest interface
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建Pixel API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Pixel `json:"data,omitempty"`
}
