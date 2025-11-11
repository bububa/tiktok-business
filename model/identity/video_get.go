package identity

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/ttvideo"
	"github.com/bububa/tiktok-business/util"
)

// VideoGetRequest 获取认证身份下帖子 API Request
type VideoGetRequest struct {
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
	// Keyword 仅当identity_type设置为AUTH_CODE时可用。否则会报错。
	// 可以作为搜索 TikTok 帖子的关键词，可以为文本或 TikTok 帖子ID。
	// 如果您想根据文本搜索TikTok 帖子，提供最大长度为 500 个字符的文本字符串。这种搜索类型支持多语言搜索和模糊匹配。
	// 示例："summer"。
	// 如果你想根据item_id搜索 TikTok 帖子，提供最小长度为 19 个字符的数字字符串。这种搜索类型支持精确匹配。
	// 示例："1234567891234567891"
	Keyword string `json:"keyword,omitempty"`
	// Cursor 分页游标
	Cursor string `json:"cursor,omitempty"`
	// Count 请求的 TikTok 帖子数量。
	// 取值范围： 1 - 20。
	// 若本字段的值大于 20，将忽略本字段，返回与认证身份绑定的最多 20 个 TikTok 帖子
	Count int `json:"count,omitempty"`
}

// Encode implements GetRequest
func (r *VideoGetRequest) Encode() string {
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
	if r.Keyword != "" {
		values.Set("keyword", r.Keyword)
	}
	if r.Cursor != "" {
		values.Set("cursor", r.Cursor)
	}
	if r.Count > 0 {
		values.Set("count", strconv.Itoa(r.Count))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// VideoGetResponse 获取认证身份下帖子 API Response
type VideoGetResponse struct {
	model.BaseResponse
	Data *VideoGetResult `json:"data,omitempty"`
}

type VideoGetResult struct {
	// Cursor 时间戳游标或分页游标。
	// 若 identity_type 为 TT_USER 或 BC_AUTH_TT, cursor为当前请求返回的最后一个帖子的时间值。若想获取后续视频，需在下一次请求中使用此游标值。
	// 若 identity_type 为 AUTH_CODE, cursor 为分页游标。若返回中的has_more为 true, 则需将返回的 cursor 传递到后续请求，以获取下一页结果。
	Cursor string `json:"cursor,omitempty"`
	// HasMore 是否有更多数据
	HasMore bool `json:"has_more,omitempty"`
	// VideoList 帖子列表
	VideoList []ttvideo.ItemInfo `json:"video_list,omitempty"`
}
