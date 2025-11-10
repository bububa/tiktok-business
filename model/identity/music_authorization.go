package identity

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// MusicAuthorizationRequest 获取音乐版权信息 API Request
type MusicAuthorizationRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// VideoID TikTok帖子ID
	VideoID string `json:"video_id,omitempty"`
	// IdentityID 认证身份ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。枚举值：CUSTOMIZED_USER, AUTH_CODE, TT_USER, BC_AUTH_TT。如果不传入，则返回所有结果。获取更多相关信息，请访问认证身份
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。当identity_type为BC_AUTH_TT时必填
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// Locations 定向地区列表。必须使用与您创建广告时相同的地区列表
	Locations []string `json:"locations,omitempty"`
	// StartTime 广告投放开始时间 (UTC), 格式为 "YYYY-MM-DD HH:MM:SS"
	StartTime string `json:"start_time,omitempty"`
	// EndTime 广告投放结束时间(UTC), 格式为 "YYYY-MM-DD HH:MM:SS"
	EndTime string `json:"end_time,omitempty"`
}

// Encode implements GetRequest
func (r *MusicAuthorizationRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("video_id", r.VideoID)
	values.Set("identity_id", r.IdentityID)
	values.Set("identity_type", string(r.IdentityType))
	if r.IdentityAuthorizedBcID != "" {
		values.Set("identity_authorized_bc_id", r.IdentityAuthorizedBcID)
	}
	values.Set("locations", string(util.JSONMarshal(r.Locations)))
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// MusicAuthorizationResponse 获取音乐版权信息 API Response
type MusicAuthorizationResponse struct {
	model.BaseResponse
	Data struct {
		// MusicAuthorization TikTok视频的音乐版权信息
		MusicAuthorization []MusicAuthorization `json:"music_authorization,omitempty"`
	} `json:"data"`
}

// MusicAuthorization TikTok视频的音乐版权信息
type MusicAuthorization struct {
	// VideoID TikTok视频ID
	VideoID string `json:"video_id,omitempty"`
	// MusicStatus TikTok帖子中所有音乐的聚合授权状态。枚举值：WITHOUT_SONG_ID,  AUTHORIZATION_MISSING, WITH_FULL_AUTHORIZATION
	MusicStatus enum.MusicAuthorizationStatus `json:"music_status,omitempty"`
	// AuthorizationInfos 音乐的授权信息
	AuthorizationInfos []MusicAuthorizationInfo `json:"authorization_infos,omitempty"`
}

// MusicAuthorizationInfo 音乐的授权信息
type MusicAuthorizationInfo struct {
	// MusicID 歌曲ID
	MusicID string `json:"music_id,omitempty"`
	// Author 歌手
	Author string `json:"author,omitempty"`
	// Title 歌曲名称
	Title string `json:"title,omitempty"`
	// Labels 歌曲标签
	Labels []string `json:"labels,omitempty"`
	// Lyricist 音乐作词者
	Lyricist string `json:"lyricist,omitempty"`
	// Composer 音乐作曲者
	Composer string `json:"composer,omitempty"`
	// Publisher 音乐发行商
	Publisher string `json:"publisher,omitempty"`
	// AuthorizationType 授权类型。枚举值：NOT_AUTHORIZED, AUTHORIZED
	AuthoriztionType enum.AuthorizationType `json:"authorization_type,omitempty"`
}
