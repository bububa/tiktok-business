package identity

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/ttvideo"
	"github.com/bububa/tiktok-business/util"
)

// VideoInfoRequest 获取 TikTok 帖子信息 API Request
type VideoInfoRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// IdentityID 认证身份ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。枚举值：CUSTOMIZED_USER, AUTH_CODE, TT_USER, BC_AUTH_TT。如果不传入，则返回所有结果。获取更多相关信息，请访问认证身份
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。当identity_type为BC_AUTH_TT时必填
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// ItemType 想要获取的 TikTok 帖子类型。
	// 枚举值：
	// VIDEO：视频帖子。
	// CAROUSEL：照片帖子。
	// 默认值：VIDEO
	ItemType enum.TikTokItemType `json:"item_type,omitempty"`
	// ItemID item_id 和 item_ids 二者必传其一。
	// 若想获取单个 TikTok 帖子的信息，传入 item_id。
	// 若想获取多个 TikTok 帖子的信息，传入 item_ids。
	// 与指定的认证身份 ID （identity_id）绑定的 TikTok 帖子 ID。
	// 若想获取与某一认证身份绑定的 TikTok 帖子，可使用 /identity/video/get/，通过返回中的 item_id 字段获取 TikTok 帖子 ID
	ItemID string `json:"item_id,omitempty"`
	// ItemIDs item_id 和 item_ids 二者必传其一。
	// 若想获取单个 TikTok 帖子的信息，传入 item_id。
	// 若想获取多个 TikTok 帖子的信息，传入 item_ids。
	// 与指定的认证身份 ID （identity_id）绑定的 TikTok 帖子 ID 列表。
	// 最大数量：20。
	// 若想获取与某一认证身份绑定的 TikTok 帖子，可使用/identity/video/get/，通过返回中的 item_id 字段获取 TikTok 帖子 ID
	ItemIDs []string `json:"item_ids,omitempty"`
}

// Encode implements GetRequest
func (r *VideoInfoRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("identity_id", r.IdentityID)
	values.Set("identity_type", string(r.IdentityType))
	if r.IdentityAuthorizedBcID != "" {
		values.Set("identity_authorized_bc_id", r.IdentityAuthorizedBcID)
	}
	if r.ItemType != "" {
		values.Set("item_type", string(r.ItemType))
	}
	if r.ItemID != "" {
		values.Set("item_id", r.ItemID)
	}
	if len(r.ItemIDs) > 0 {
		values.Set("item_ids", string(util.JSONMarshal(r.ItemIDs)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// VideoInfoResponse 获取 TikTok 帖子信息 API Response
type VideoInfoResponse struct {
	model.BaseResponse
	Data *VideoInfoResult `json:"data,omitempty"`
}

type VideoInfoResult struct {
	// VideoDetail 请求中传入 item_id 字段时返回。
	// TikTok 帖子信息。
	VideoDetail *ttvideo.ItemInfo `json:"video_detail,omitempty"`
	// VideoDetails 请求中传入 item_ids 字段时返回。
	// TikTok 帖子信息
	VideoDetails []ttvideo.ItemInfo `json:"video_details,omitempty"`
}
