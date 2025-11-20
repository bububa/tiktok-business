package asset

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DeleteRequest 删除创意资产 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AssetType 您想要分享的创意资产类型。
	// 枚举值：VIDEO, IMAGE, MUSIC
	// 默认值：VIDEO
	AssetType enum.CreativeAssetType `json:"asset_type,omitempty"`
	// VideoIDs 视频ID列表。最多可包含50个ID
	VideoIDs []string `json:"video_ids,omitempty"`
	// ImageIDs 图片ID列表。最多可包含50个图片ID
	ImageIDs []string `json:"image_ids,omitempty"`
}

// Encode implements PostRequest
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeleteResponse 删除创意资产 API Response
type DeleteResponse struct {
	model.BaseResponse
	Data *DeleteResult `json:"data,omitempty"`
}

type DeleteResult struct {
	// FailedVideoIDs 删除失败的视频ID列表
	FailedVideoIDs []string `json:"failed_video_ids,omitempty"`
	// FailedImageIDs 删除失败的图片ID列表
	FailedImageIDs []string `json:"failed_image_ids,omitempty"`
}
