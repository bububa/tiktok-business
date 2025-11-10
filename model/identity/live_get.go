package identity

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// LiveGetRequest 获取认证身份下直播视频 API Request
type LiveGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// IdentityID 认证身份ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。枚举值：CUSTOMIZED_USER, AUTH_CODE, TT_USER, BC_AUTH_TT。如果不传入，则返回所有结果。获取更多相关信息，请访问认证身份
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。当identity_type为BC_AUTH_TT时必填
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// Cursor 分页游标
	Cursor string `json:"cursor,omitempty"`
}

// Encode implements GetRequest
func (r *LiveGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("identity_id", r.IdentityID)
	values.Set("identity_type", string(r.IdentityType))
	if r.IdentityAuthorizedBcID != "" {
		values.Set("identity_authorized_bc_id", r.IdentityAuthorizedBcID)
	}
	if r.Cursor != "" {
		values.Set("cursor", r.Cursor)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// LiveGetResponse 获取认证身份下直播视频 API Response
type LiveGetResponse struct {
	model.BaseResponse
	Data *LiveGetResult `json:"data,omitempty"`
}

type LiveGetResult struct {
	// Cursor 时间戳游标。 当前请求返回的最后一项的时间值。您需要在后续请求中使用此游标值来获取下一个直播视频
	Cursor string `json:"cursor,omitempty"`
	// HasMore 是否有更多数据
	HasMore bool `json:"has_more,omitempty"`
	// LiveList 直播视频列表
	LiveList []Live `json:"live_list,omitempty"`
}

// Live 直播视频
type Live struct {
	// LiveID 视频列表ID
	LiveID string `json:"live_id,omitempty"`
	// FinishTimestamp 直播视频结束时间戳。 示例：2023-12-07 08:13:44
	FinishTimestamp model.DateTime `json:"finish_timestamp,omitzero"`
}
