package video

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdInfoRequest 获取视频信息 API Request
type AdInfoRequest struct {
	// AdvertiserID 广告账号
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// VideoIDs 视频 ID 列表。
	// 最大数量：60。
	// 注意：若视频不可用于创建广告，视频 ID 将被忽略。
	VideoIDs []string `json:"video_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r AdInfoRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("video_ids", string(util.JSONMarshal(r.VideoIDs)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AdInfoResponse 获取视频信息 API Response
type AdInfoResponse struct {
	model.BaseResponse
	Data struct {
		// List 视频数据
		List []Video `json:"list,omitempty"`
	} `json:"data"`
}
