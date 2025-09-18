package video

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SuggestCoverRequest 获取视频推荐封面 API Request
type SuggestCoverRequest struct {
	// AdvertiserID 广告账号
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// PosterNumber 要获取的封面数目。取值范围：1-10，默认为10。若候选封面总数少于本字段设置的数量，仅返回全部候选封面。
	PosterNumber int `json:"poster_number,omitempty"`
}

// Encode implements GetRequest interface
func (r *SuggestCoverRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("video_id", r.VideoID)
	if r.PosterNumber > 0 {
		values.Set("poster_number", strconv.Itoa(r.PosterNumber))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// SuggestCoverResponse 获取视频推荐封面 API Response
type SuggestCoverResponse struct {
	model.BaseResponse
	Data struct {
		// List 封面列表
		List []SuggestCover `json:"list,omitempty"`
	} `json:"data"`
}

type SuggestCover struct {
	// ID 图片ID
	ID string `json:"id,omitempty"`
	// Width 图片宽度
	Width int `json:"width,omitempty"`
	// Height 图片高度
	Height int `json:"height,omitempty"`
	// URL 图片预览地址，有效期1小时，过期需重新获取
	URL string `json:"url,omitempty"`
}
