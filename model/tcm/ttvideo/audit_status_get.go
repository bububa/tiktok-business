package ttvideo

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AuditStatusGetRequest 获取视频审核状态 API Request
type AuditStatusGetRequest struct {
	// TCMAccountID TikTok Creator Marketplace 账户 ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// V2OrderID 与您想要了解审核状态的视频相关联的订单的 ID。
	// 您可通过/tcm/order/get/v2/获取订单的 ID。
	V2OrderID string `json:"v2_order_id,omitempty"`
	// Lang 审核状态信息的语言。
	// 枚举值：ja，ko，en，de，zh，th-TH，ms-MY，id-ID，it，fr，ru-RU，ar。
	Lang string `json:"lang,omitempty"`
}

// Encode implements GetRequest
func (r *AuditStatusGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("tcm_account_id", r.TCMAccountID)
	values.Set("v2_order_id", r.V2OrderID)
	if r.Lang != "" {
		values.Set("lang", r.Lang)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AuditStatusGetResponse 获取视频审核状态 API Response
type AuditStatusGetResponse struct {
	model.BaseResponse
	Data struct {
		// AuditStatusList 视频状态列表
		AuditStatusList []AuditStatus `json:"audit_status_list,omitempty"`
	} `json:"data"`
}

// AuditStatus 视频状态
type AuditStatus struct {
	// VideoID 视频 ID
	VideoID string `json:"video_id,omitempty"`
	// VideoStatus 视频状态。枚举值：INIT，UPLOAD，RELEASE，DELETE
	VideoStatus enum.TTVideoStatus `json:"video_status,omitempty"`
	// BindStatus 绑定状态。枚举值：BIND，UNBIND
	BindStatus enum.BindStatus `json:"bind_status,omitempty"`
	// ReviewStatus 审核状态。枚举值：NO_REVIEW，CONTENT_REVIEWING，AD_REVIEWING，CONTENT_REVIEW_FAIL，AD_REVIEW_FAIL，SUCCEED，ALL_FAIL。
	ReviewStatus enum.TTVideoReviewStatus `json:"review_status,omitempty"`
	// ReviewCreateTime 创建审核的 UTC+0 日期和时间
	ReviewCreateTime model.DateTime `json:"review_create_time,omitzero"`
	// UploadTime 视频上传的 UTC+0 日期和时间
	UploadTime model.DateTime `json:"upload_time,omitzero"`
	// CoverImageURL 视频封面图片的 URI
	CoverImageURL string `json:"cover_image_url,omitempty"`
	// IsAdvertiserReject 是否被广告主拒绝
	IsAdvertiserReject bool `json:"is_advertiser_reject,omitempty"`
	// MusicReviewStatus 音乐审核状态。枚举值：INIT,REVIEWING,SUCCEED,FAIL
	MusicReviewStatus enum.ReviewStatus `json:"music_review_status,omitempty"`
	// MusicReviewTime 音乐审核的 UTC+0 日期和时间
	MusicReviewTime model.DateTime `json:"music_review_time,omitzero"`
	// RejectReason 视频遭拒的原因描述
	RejectReason string `json:"reject_reason,omitempty"`
	// OperationTime 视频获批或遭拒的 UTC+0 日期和时间
	OperationTime model.DateTime `json:"operation_time,omitzero"`
	// DeleteTime 视频被删除的 UTC+0 日期和时间
	DeleteTime model.DateTime `json:"delete_time,omitzero"`
	// PublishTime 视频发布的 UTC+0 日期和时间
	PublishTime model.DateTime `json:"publish_time,omitzero"`
	// ContentJudgeTime 作出内容审核决定的 UTC+0 日期和时间
	ContentJudgeTime model.DateTime `json:"content_judge_time,omitzero"`
	// AdJudgeTime 作出广告审核决定的 UTC+0 日期和时间
	AdJudgeTime model.DateTime `json:"ad_judge_time,omitzero"`
	// AdRejectReasons 广告遭拒原因列表
	AdRejectReasons []string `json:"ad_reject_reasons,omitempty"`
	// AdJudgeStatus 广告审核状态。枚举值：INIT、REVIEWING，SUCCEED，FAIL
	AdJudgeStatus enum.ReviewStatus `json:"ad_judge_status,omitempty"`
	// ContentJudgeStatus 内容审核状态。枚举值：INIT、REVIEWING，SUCCEED，FAIL
	ContentJudgeStatus enum.ReviewStatus `json:"content_judge_status,omitempty"`
	// AdvertiserOperate 广告主对创作者视频的决定。枚举值：DEFAULT，REJECT，APPROVE
	AdvertiserOperate string `json:"advertiser_operate,omitempty"`
	// EmbedURL 用于访问视频的url。 注意：仅当视频的review_status是SUCCEED时，此字段才不会为空值。
	EmbedURL string `json:"embed_url,omitempty"`
}
