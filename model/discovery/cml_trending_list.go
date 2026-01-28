package discovery

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CMLTrendingListRequest Get popular tracks from the Commercial Music Library API Request
type CMLTrendingListRequest struct {
	// BusinessID Application specific unique identifier for the TikTok account.
	// Pass the value of the open_id field returned in the response of /tt_user/oauth2/token/ to this field.
	BusinessID string `json:"business_id,omitempty"`
	// Genre The specific genre to filter the results by.
	// Default value: ALL.
	// For enum values, see List of values for genre.
	Genre string `json:"genre,omitempty"`
	// CountryCode The code of the location to filter the results by.
	// Default value: US.
	// See Location code for the supported values.
	CountryCode string `json:"country_code,omitempty"`
	// DateRange The timeframe in which the tracks have been popular.
	// Enum values:
	// 1DAY: last one day.
	// 7DAY: last seven days.
	// 30DAY: last 30 days.
	// 90DAY: last 90 days.
	// Default value: 7DAY.
	DateRange string `json:"date_range,omitempty"`
}

// Encode implements GetRequest
func (r *CMLTrendingListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	if r.Genre != "" {
		values.Set("genre", r.Genre)
	}
	if r.CountryCode != "" {
		values.Set("country_code", r.CountryCode)
	}
	if r.DateRange != "" {
		values.Set("date_range", r.DateRange)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// CMLTrendingListResponse Get popular tracks from the Commercial Music Library API Response
type CMLTrendingListResponse struct {
	model.BaseResponse
	Data struct {
		// List The list of the trending tracks.
		List []CommercialMusicLibrary `json:"list,omitempty"`
	} `json:"data"`
}

type CommercialMusicLibrary struct {
	// CommercialMusicID The ID of the track.
	CommercialMusicID string `json:"commercial_music_id,omitempty"`
	// CommercialMusicName The name of the track
	CommercialMusicName string `json:"commercial_music_name,omitempty"`
	// Duration The duration of the track in seconds.
	// Note: For a track that is only a part (clips) of the full piece of music and not the entire piece, the duration might be recorded as 0.
	Duration int `json:"duration,omitempty"`
	// ThumbnailURL The URL of the thumbnail image associated with the track.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// Artist The name of the artist who performed the track.
	Artist string `json:"artist,omitempty"`
	// PreviewURL The URL for a preview of the track.
	// Validity period: six hours.
	PreviewURL string `json:"preview_url,omitempty"`
	// Genres A list of genres associated with the track.
	Genres []string `json:"genres,omitempty"`
	// RankPositon Current rank position of the track in the selected region (country_code).
	RankPosition model.Int `json:"rank_position,omitempty"`
	// TrendingHistory The historical ranking data for the track, spanning up to the past 30 days.
	TrendingHistory []TrendingHistory `json:"trending_history,omitempty"`
	// FullDurationSongClip Information related to the track, including the full duration of the track
	FullDurationSongClip *SongClip `json:"full_duration_song_clip,omitempty"`
	// TrendingSongClip Information related to the most popular track in the last 30 days.
	TrendingSongClip *SongClip `json:"trending_song_clip,omitempty"`
	// TopVideoList The list of the top 20 trending videos for the track
	TopVideoList []Video `json:"top_video_list,omitempty"`
}

// SongClip Information related to the track
type SongClip struct {
	// PreviewURL The track preview URL.
	// Validity period: 6 hours.
	// Note: This is the same as the preview_url parameter outside full_duration_song_clip.
	PreviewURL string `json:"preview_url,omitempty"`
	// Duration The track duration in seconds.
	// Note: This value is the same as the existing duration value..
	Duration int `json:"duration,omitempty"`
	// SongClipID The track ID.
	// Note: This value can be passed to the music_sound_id in /business/video/publish/.
	SongClipID string `json:"song_clip_id,omitempty"`
}
