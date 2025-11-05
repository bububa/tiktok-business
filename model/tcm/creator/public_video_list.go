package creator

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// PublicVideoListRequest 获取公开媒体洞察 API Request
type PublicVideoListRequest struct {
	// TCMAccountID TikTok Creator Marketplace 账户 ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// HandleName 您想要搜索的 TikTok 创作者的用户名
	HandleName string `json:"handle_name,omitempty"`
	// Limit 每个请求中的视频数量。可以是 1 到 10 之间的一个数字。默认值为 10。
	Limit int `json:"limit,omitempty"`
	// VideoIDs 某个视频的ID，用于获取信息。
	// 您可以从本接口的返回中获取视频ID，或从/tcm/report/get/或/tcm/report/get/v2/的返回中获取。
	VideoIDs []string `json:"video_ids,omitempty"`
	// Cursor 返回某时间戳值（以毫秒为单位）之前的视频。默认值是 0，将返回最近的视频
	Cursor int64 `json:"cursor,omitempty"`
}

// Encode implements GetRequest
func (r *PublicVideoListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("tcm_account_id", r.TCMAccountID)
	values.Set("handle_name", r.HandleName)
	if r.Limit > 0 {
		values.Set("limit", strconv.Itoa(r.Limit))
	}
	if len(r.VideoIDs) > 0 {
		values.Set("video_ids", string(util.JSONMarshal(r.VideoIDs)))
	}
	if r.Cursor > 0 {
		values.Set("cursor", strconv.FormatInt(r.Cursor, 10))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// PublicVideoListResponse 获取公开媒体洞察 API Response
type PublicVideoListResponse struct {
	model.BaseResponse
	Data *PublicVideoListResult `json:"data,omitempty"`
}

type PublicVideoListResult struct {
	// PageInfo 返回数据的信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Posts 帖子列表
	Posts []Post `json:"posts,omitempty"`
}
