package creator

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

type StatusGetRequest struct {
	AccountID   string   `json:"tto_tcm_account_id,omitempty"`
	HandleNames []string `json:"handle_names,omitempty"`
}

func (r *StatusGetRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("tto_tcm_account_id", r.AccountID)
	if len(r.HandleNames) > 0 {
		v.Set("handle_names", string(util.JSONMarshal(r.HandleNames)))
	}
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type OnboardingStatus struct {
	HandleName string                `json:"handle_name,omitempty"`
	Status     enum.TCMCreatorStatus `json:"status,omitempty"`
}

type StatusGetResponse struct {
	model.BaseResponse
	Data struct {
		OnboardingStatus []OnboardingStatus `json:"onboarding_status,omitempty"`
	} `json:"data"`
}

type PublicGetRequest struct {
	AccountID  string `json:"tto_tcm_account_id,omitempty"`
	HandleName string `json:"handle_name,omitempty"`
}

func (r *PublicGetRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("tto_tcm_account_id", r.AccountID)
	v.Set("handle_name", r.HandleName)
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type CreatorResponse struct {
	model.BaseResponse
	Data *Creator `json:"data,omitempty"`
}

type PublicVideoListRequest struct {
	AccountID  string      `json:"tto_tcm_account_id,omitempty"`
	HandleName string      `json:"handle_name,omitempty"`
	VideoIDs   []string    `json:"video_ids,omitempty"`
	Cursor     model.Int64 `json:"cursor,omitempty"`
	Limit      int         `json:"limit,omitempty"`
}

func (r *PublicVideoListRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("tto_tcm_account_id", r.AccountID)
	v.Set("handle_name", r.HandleName)
	if len(r.VideoIDs) > 0 {
		v.Set("video_ids", string(util.JSONMarshal(r.VideoIDs)))
	}
	if r.Cursor != 0 {
		v.Set("cursor", r.Cursor.String())
	}
	if r.Limit > 0 {
		v.Set("limit", strconv.Itoa(r.Limit))
	}
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type PublicVideoListResponse struct {
	model.BaseResponse
	Data *PublicVideoListResult `json:"data,omitempty"`
}

type PublicVideoListResult struct {
	Posts    []Post `json:"posts,omitempty"`
	PageInfo struct {
		HasMore bool        `json:"has_more,omitempty"`
		Cursor  model.Int64 `json:"cursor,omitempty"`
	} `json:"page_info,omitempty"`
}

type AuthorizedGetRequest struct {
	CreatorID string   `json:"creator_id,omitempty"`
	Fields    []string `json:"fields,omitempty"`
}

func (r *AuthorizedGetRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("creator_id", r.CreatorID)
	if len(r.Fields) > 0 {
		v.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type AuthorizedVideoListRequest struct {
	CreatorID string      `json:"creator_id,omitempty"`
	VideoIDs  []string    `json:"video_ids,omitempty"`
	Limit     int         `json:"limit,omitempty"`
	Cursor    model.Int64 `json:"cursor,omitempty"`
}

func (r *AuthorizedVideoListRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("creator_id", r.CreatorID)
	if len(r.VideoIDs) > 0 {
		v.Set("video_ids", string(util.JSONMarshal(r.VideoIDs)))
	}
	if r.Limit > 0 {
		v.Set("limit", strconv.Itoa(r.Limit))
	}
	if r.Cursor != 0 {
		v.Set("cursor", r.Cursor.String())
	}
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type AuthorizedVideoListResponse struct {
	model.BaseResponse
	Data *AuthorizedVideoListResult `json:"data,omitempty"`
}

type AuthorizedVideoListResult struct {
	Posts   []Post      `json:"posts,omitempty"`
	HasMore bool        `json:"has_more,omitempty"`
	Cursor  model.Int64 `json:"cursor,omitempty"`
}
