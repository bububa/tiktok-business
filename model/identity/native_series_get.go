package identity

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// NativeSeriesGetRequest 获取广告账号内可用 TikTok Series API Request。
type NativeSeriesGetRequest struct {
	// AdvertiserID 广告主 ID。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// IdentityID 认证身份 ID。
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。支持 TT_USER、BC_AUTH_TT。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 当 IdentityType 为 BC_AUTH_TT 时必填。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// Cursor 分页游标。默认值：0。
	Cursor int64 `json:"cursor,omitempty"`
	// Count 每页数量。取值范围：1-100，默认值：10。
	Count int `json:"count,omitempty"`
}

// Encode implements GetRequest.
func (r *NativeSeriesGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("identity_id", r.IdentityID)
	values.Set("identity_type", string(r.IdentityType))
	if r.IdentityAuthorizedBcID != "" {
		values.Set("identity_authorized_bc_id", r.IdentityAuthorizedBcID)
	}
	if r.Cursor > 0 {
		values.Set("cursor", strconv.FormatInt(r.Cursor, 10))
	}
	if r.Count > 0 {
		values.Set("count", strconv.Itoa(r.Count))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// NativeSeriesGetResponse 获取广告账号内可用 TikTok Series API Response。
type NativeSeriesGetResponse struct {
	model.BaseResponse
	Data *NativeSeriesGetResult `json:"data,omitempty"`
}

// NativeSeriesGetResult 广告账号内可用 TikTok Series 列表。
type NativeSeriesGetResult struct {
	// List TikTok Series 列表。
	List []NativeSeries `json:"list,omitempty"`
	// Cursor 下一页游标。
	Cursor int64 `json:"cursor,omitempty"`
	// HasMore 是否还有更多数据。
	HasMore bool `json:"has_more,omitempty"`
}

// NativeSeries TikTok Series 信息。
type NativeSeries struct {
	// NativeSeriesID TikTok Series ID。
	NativeSeriesID string `json:"native_series_id,omitempty"`
	// NativeSeriesName TikTok Series 名称。
	NativeSeriesName string `json:"native_series_name,omitempty"`
	// NativeSeriesCoverURL TikTok Series 封面 URL。
	NativeSeriesCoverURL string `json:"native_series_cover_url,omitempty"`
	// NativeSeriesTotalEpisode TikTok Series 总集数。
	NativeSeriesTotalEpisode int `json:"native_series_total_episode,omitempty"`
	// NativeSeriesTotalDuration TikTok Series 总时长，单位为秒。
	NativeSeriesTotalDuration int `json:"native_series_total_duration,omitempty"`
}
