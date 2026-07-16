package anchor

import (
	"io"
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

type CallToAction struct {
	LandingPageURL      string `json:"landing_page_url,omitempty"`
	IOSDownloadLink     string `json:"ios_download_link,omitempty"`
	IOSDeepLink         string `json:"ios_deep_link,omitempty"`
	AndroidDownloadLink string `json:"android_download_link,omitempty"`
	AndroidDeepLink     string `json:"android_deep_link,omitempty"`
}

type CouponLink struct {
	Discount   model.Float64 `json:"discount,omitempty"`
	CouponURL  string        `json:"coupon_url,omitempty"`
	CouponCode string        `json:"coupon_code,omitempty"`
}

type Anchor struct {
	AnchorID             string                `json:"anchor_id,omitempty"`
	Status               enum.TTOAnchorStatus  `json:"status,omitempty"`
	AnchorType           enum.TTOAnchorType    `json:"anchor_type,omitempty"`
	AnchorSubType        enum.TTOAnchorSubType `json:"anchor_sub_type,omitempty"`
	CategoryLabelID      model.Int64           `json:"category_label_id,omitempty"`
	CountryCode          string                `json:"country_code,omitempty"`
	CallToAction         *CallToAction         `json:"call_to_action,omitempty"`
	AnchorTitle          string                `json:"anchor_title,omitempty"`
	AnchorName           string                `json:"anchor_name,omitempty"`
	AnchorSubTitle       string                `json:"anchor_sub_title,omitempty"`
	ThumbnailURL         string                `json:"thumbnail_url,omitempty"`
	CouponLink           *CouponLink           `json:"coupon_link,omitempty"`
	PixelID              string                `json:"pixel_id,omitempty"`
	CreatorPreviewURL    string                `json:"creator_preview_url,omitempty"`
	AdvertiserPreviewURL string                `json:"advertiser_preview_url,omitempty"`
}

type CreateRequest struct {
	AccountID       string                `json:"tto_tcm_account_id,omitempty"`
	AnchorType      enum.TTOAnchorType    `json:"anchor_type,omitempty"`
	AnchorSubType   enum.TTOAnchorSubType `json:"anchor_sub_type,omitempty"`
	CategoryLabelID model.Int64           `json:"category_label_id,omitempty"`
	CountryCode     string                `json:"country_code,omitempty"`
	CallToAction    *CallToAction         `json:"call_to_action,omitempty"`
	AnchorTitle     string                `json:"anchor_title,omitempty"`
	AnchorName      string                `json:"anchor_name,omitempty"`
	UploadType      enum.UploadType       `json:"upload_type,omitempty"`
	ThumbnailURL    string                `json:"thumbnail_url,omitempty"`
}

func (r *CreateRequest) Encode() []byte { return util.JSONMarshal(r) }

type CreateByFileRequest struct {
	AccountID         string                `json:"tto_tcm_account_id,omitempty"`
	AnchorType        enum.TTOAnchorType    `json:"anchor_type,omitempty"`
	AnchorSubType     enum.TTOAnchorSubType `json:"anchor_sub_type,omitempty"`
	CategoryLabelID   model.Int64           `json:"category_label_id,omitempty"`
	CountryCode       string                `json:"country_code,omitempty"`
	CallToAction      *CallToAction         `json:"call_to_action,omitempty"`
	AnchorTitle       string                `json:"anchor_title,omitempty"`
	AnchorName        string                `json:"anchor_name,omitempty"`
	ThumbnailFile     io.Reader             `json:"-"`
	ThumbnailFileName string                `json:"-"`
}

func (r *CreateByFileRequest) Encode() []model.UploadField {
	fields := []model.UploadField{
		{Key: "tto_tcm_account_id", Value: r.AccountID}, {Key: "anchor_type", Value: string(r.AnchorType)},
		{Key: "anchor_sub_type", Value: string(r.AnchorSubType)}, {Key: "category_label_id", Value: r.CategoryLabelID.String()},
		{Key: "country_code", Value: r.CountryCode}, {Key: "call_to_action", Value: string(util.JSONMarshal(r.CallToAction))},
		{Key: "anchor_title", Value: r.AnchorTitle}, {Key: "anchor_name", Value: r.AnchorName}, {Key: "upload_type", Value: string(enum.UPLOAD_BY_FILE)},
	}
	if r.ThumbnailFile != nil {
		fields = append(fields, model.UploadField{Key: "thumbnail_file", Value: r.ThumbnailFileName, Reader: r.ThumbnailFile})
	}
	return fields
}

type CreateResponse struct {
	model.BaseResponse
	Data *Anchor `json:"data,omitempty"`
}

type GetRequest struct {
	AccountID string   `json:"tto_tcm_account_id,omitempty"`
	AnchorIDs []string `json:"anchor_ids,omitempty"`
	Page      int      `json:"page,omitempty"`
	PageSize  int      `json:"page_size,omitempty"`
}

func (r *GetRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("tto_tcm_account_id", r.AccountID)
	if len(r.AnchorIDs) > 0 {
		v.Set("anchor_ids", string(util.JSONMarshal(r.AnchorIDs)))
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

type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	Anchors  []Anchor        `json:"anchors,omitempty"`
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type DeleteRequest struct {
	AccountID string `json:"tto_tcm_account_id,omitempty"`
	AnchorID  string `json:"anchor_id,omitempty"`
}

func (r *DeleteRequest) Encode() []byte { return util.JSONMarshal(r) }

type DeleteResponse struct {
	model.BaseResponse
	Data struct{} `json:"data"`
}
