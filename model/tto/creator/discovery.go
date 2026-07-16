package creator

import (
	"net/url"
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

type LabelGetRequest struct {
	AccountID string            `json:"tto_tcm_account_id,omitempty"`
	LabelType enum.TTOLabelType `json:"label_type,omitempty"`
}

func (r *LabelGetRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("tto_tcm_account_id", r.AccountID)
	v.Set("label_type", string(r.LabelType))
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type LabelGetResponse struct {
	model.BaseResponse
	Data *LabelGetResult `json:"data,omitempty"`
}

type LabelGetResult struct {
	IndustryLabels []Label `json:"industry_labels,omitempty"`
	ContentLabels  []Label `json:"content_labels,omitempty"`
}

type RankGetRequest struct {
	AccountID          string                  `json:"tto_tcm_account_id,omitempty"`
	RankingType        enum.TTORankingType     `json:"ranking_type,omitempty"`
	TimePeriod         enum.TTORankingPeriod   `json:"time_period,omitempty"`
	TimePeriodLookback enum.TTORankingLookback `json:"time_period_lookback,omitempty"`
	LabelID            string                  `json:"label_id,omitempty"`
	CountryCode        string                  `json:"country_code,omitempty"`
	Page               int                     `json:"page,omitempty"`
	PageSize           int                     `json:"page_size,omitempty"`
}

func (r *RankGetRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("tto_tcm_account_id", r.AccountID)
	v.Set("ranking_type", string(r.RankingType))
	v.Set("time_period", string(r.TimePeriod))
	v.Set("time_period_lookback", string(r.TimePeriodLookback))
	v.Set("label_id", r.LabelID)
	v.Set("country_code", r.CountryCode)
	if r.Page > 0 {
		v.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		v.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type RankedCreator struct {
	Creator
	RankingChange                string        `json:"ranking_change,omitempty"`
	Country                      string        `json:"country,omitempty"`
	Score                        model.Float64 `json:"score,omitempty"`
	NumberOfFollowers            model.Int64   `json:"number_of_followers,omitempty"`
	BrandedContentVideoViews     model.Int64   `json:"branded_content_video_views,omitempty"`
	BrandedContentEngagementRate model.Float64 `json:"branded_content_engagement_rate,omitempty"`
	TotalVideoViews              model.Int64   `json:"total_video_views,omitempty"`
	StartingPrice                model.Int64   `json:"starting_price,omitempty"`
	StartingPriceCurrency        string        `json:"starting_price_currency,omitempty"`
	FollowerGrowth               model.Int64   `json:"follower_growth,omitempty"`
}
type RankGetResponse struct {
	model.BaseResponse
	Data *RankGetResult `json:"data,omitempty"`
}

type RankGetResult struct {
	Creators      []RankedCreator `json:"creators,omitempty"`
	LastUpdatedAt model.DateTime  `json:"last_updated_at,omitzero"`
	PageInfo      *model.PageInfo `json:"page_info,omitempty"`
}

type DiscoverRequest struct {
	AccountID            string                   `json:"tto_tcm_account_id,omitempty"`
	CountryCodes         []string                 `json:"country_codes,omitempty"`
	StateProvinces       []string                 `json:"state_provinces,omitempty"`
	ContentLabelIDs      []string                 `json:"content_label_ids,omitempty"`
	IndustryLabelIDs     []string                 `json:"industry_label_ids,omitempty"`
	MinFollowers         model.Int64              `json:"min_followers,omitempty"`
	MaxFollowers         model.Int64              `json:"max_followers,omitempty"`
	Languages            []string                 `json:"languages,omitempty"`
	MinCreatorPrice      model.Float64            `json:"min_creator_price,omitempty"`
	MaxCreatorPrice      model.Float64            `json:"max_creator_price,omitempty"`
	CreatorPriceCurrency string                   `json:"creator_price_currency,omitempty"`
	MinAvgViews          model.Int64              `json:"min_avg_views,omitempty"`
	MaxAvgViews          model.Int64              `json:"max_avg_views,omitempty"`
	MinMedianViews       model.Int64              `json:"min_median_views,omitempty"`
	MaxMedianViews       model.Int64              `json:"max_median_views,omitempty"`
	MinEngagementRate    model.Float64            `json:"min_engagement_rate,omitempty"`
	MaxEngagementRate    model.Float64            `json:"max_engagement_rate,omitempty"`
	FollowerCountryCodes []string                 `json:"follower_country_codes,omitempty"`
	FollowerGenderRatio  string                   `json:"follower_gender_ratio,omitempty"`
	FollowerAge          enum.TTOAudienceAge      `json:"follower_age,omitempty"`
	KeywordSearch        string                   `json:"keyword_search,omitempty"`
	SortField            enum.TTOCreatorSortField `json:"sort_field,omitempty"`
	SortOrder            enum.TTOSortOrder        `json:"sort_order,omitempty"`
	Page                 int                      `json:"page,omitempty"`
	PageSize             int                      `json:"page_size,omitempty"`
}

func setJSON(v url.Values, key string, value any) {
	b := util.JSONMarshal(value)
	if string(b) != "null" && string(b) != "[]" {
		v.Set(key, string(b))
	}
}
func (r *DiscoverRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("tto_tcm_account_id", r.AccountID)
	setJSON(v, "country_codes", r.CountryCodes)
	setJSON(v, "state_provinces", r.StateProvinces)
	setJSON(v, "content_label_ids", r.ContentLabelIDs)
	setJSON(v, "industry_label_ids", r.IndustryLabelIDs)
	if r.MinFollowers != 0 {
		v.Set("min_followers", r.MinFollowers.String())
	}
	if r.MaxFollowers != 0 {
		v.Set("max_followers", r.MaxFollowers.String())
	}
	setJSON(v, "languages", r.Languages)
	if r.MinCreatorPrice != 0 {
		v.Set("min_creator_price", r.MinCreatorPrice.String(-1))
	}
	if r.MaxCreatorPrice != 0 {
		v.Set("max_creator_price", r.MaxCreatorPrice.String(-1))
	}
	if r.CreatorPriceCurrency != "" {
		v.Set("creator_price_currency", r.CreatorPriceCurrency)
	}
	if r.MinAvgViews != 0 {
		v.Set("min_avg_views", r.MinAvgViews.String())
	}
	if r.MaxAvgViews != 0 {
		v.Set("max_avg_views", r.MaxAvgViews.String())
	}
	if r.MinMedianViews != 0 {
		v.Set("min_median_views", r.MinMedianViews.String())
	}
	if r.MaxMedianViews != 0 {
		v.Set("max_median_views", r.MaxMedianViews.String())
	}
	if r.MinEngagementRate != 0 {
		v.Set("min_engagement_rate", r.MinEngagementRate.String(-1))
	}
	if r.MaxEngagementRate != 0 {
		v.Set("max_engagement_rate", r.MaxEngagementRate.String(-1))
	}
	setJSON(v, "follower_country_codes", r.FollowerCountryCodes)
	if r.FollowerGenderRatio != "" {
		v.Set("follower_gender_ratio", r.FollowerGenderRatio)
	}
	if r.FollowerAge != "" {
		v.Set("follower_age", string(r.FollowerAge))
	}
	if r.KeywordSearch != "" {
		v.Set("keyword_search", r.KeywordSearch)
	}
	if r.SortField != "" {
		v.Set("sort_field", string(r.SortField))
	}
	if r.SortOrder != "" {
		v.Set("sort_order", string(r.SortOrder))
	}
	if r.Page > 0 {
		v.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		v.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type DiscoverResponse struct {
	model.BaseResponse
	Data *DiscoverResult `json:"data,omitempty"`
}

type DiscoverResult struct {
	Creators []Creator       `json:"creators,omitempty"`
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
