package creator

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AuthorizedVideoListRequest 授权后获取媒体洞察 API Request
type AuthorizedVideoListRequest struct {
	// CreatorID TikTok 创作者 ID
	CreatorID string `json:"creator_id,omitempty"`
	// Limit 每个请求中的视频数量。可以是 1 到 10 之间的一个数字。默认值为 10。
	Limit int `json:"limit,omitempty"`
	// VideoIDs 某个视频的ID，用于获取信息。
	// 您可以从本接口的返回中获取视频ID，或从/tcm/report/get/或/tcm/report/get/v2/的返回中获取。
	VideoIDs []string `json:"video_ids,omitempty"`
	// Cursor 返回某时间戳值（以毫秒为单位）之前的视频。默认值是 0，将返回最近的视频
	Cursor int64 `json:"cursor,omitempty"`
}

// Encode implements GetRequest
func (r *AuthorizedVideoListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("creator_id", r.CreatorID)
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

// AuthorizedVideoListResponse 授权后获取媒体洞察 API Response
type AuthorizedVideoListResponse struct {
	model.BaseResponse
	Data *AuthorizedVideoListResult `json:"data,omitempty"`
}

type AuthorizedVideoListResult struct {
	// PageInfo 返回数据的信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Posts 帖子列表
	Posts []Post `json:"posts,omitempty"`
}
