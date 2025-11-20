package asset

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ShareRequest 分享创意资产 API Request
type ShareRequest struct {
	// AdvertiserID 执行操作的广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AssetType 您想要分享的创意资产类型。
	// 枚举值：VIDEO, IMAGE, MUSIC
	// 默认值：VIDEO
	AssetType enum.CreativeAssetType `json:"asset_type,omitempty"`
	// MaterialIDs 素材ID。 最大数量：20。
	// 分享操作成功后，分享的创意素材的名字将被复制到被分享的广告账户下。
	// 注意: 目前，仅当被分享的广告账户下有某素材与被分享素材同名时，分享操作才会失败
	MaterialIDs []string `json:"material_ids,omitempty"`
	// SharedAdvertiserIDs 您想要分享给的广告主ID。
	// 最大数量：10
	SharedAdvertiserIDs []string `json:"shared_advertiser_ids,omitempty"`
}

// Encode implements PostRequest
func (r *ShareRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ShareResponse 分享创意资产 API Response
type ShareResponse struct {
	model.BaseResponse
	Data struct {
		// FailedInfos 失败分享的信息。
		// 注意: 目前，仅当被分享的广告账户下有某素材与被分享素材同名时，分享才会失败
		FailedInfos map[string][]string `json:"failed_infos,omitempty"`
	} `json:"data"`
}
