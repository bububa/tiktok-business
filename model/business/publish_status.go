package business

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// PublishStatusRequest 获取 TikTok 帖子的发布状态 API Request
type PublishStatusRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// PublishID 帖子发布任务的唯一标识符。
	// 您需将/business/video/publish/或/business/photo/publish接口返回的share_id字段值传给本字段。
	PublishID string `json:"publish_id,omitempty"`
}

// Encode implements GetRequest
func (r *PublishStatusRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	values.Set("publish_id", r.PublishID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// PublishStatusResponse 获取 TikTok 帖子的发布状态 API Response
type PublishStatusResponse struct {
	model.BaseResponse
	Data *PublishStatus `json:"data,omitempty"`
}

// PublishStatus TikTok 帖子的发布状态
type PublishStatus struct {
	// PostIDs 仅当 status 为 PUBLISH_COMPLETE 且帖子公开可见时返回本字段。
	// 发布的 TikTok 帖子的 ID 列表。
	// 注意：
	// 帖子 ID 的数据处理可能长达 3 分钟。若未直接返回帖子 ID ，需等待 3 分钟后重试。
	// 您可将本字段返回的 ID 列表传入/business/video/list/中的筛选字段video_ids，从而获取这些帖子的指标数据
	PostIDs []string `json:"post_ids,omitempty"`
	// Status TikTok 帖子的发布状态。
	// 枚举值：
	// PROCESSING_DOWNLOAD：正在从 URL 中提取内容。
	// PUBLISH_COMPLETE：帖子已通过审核且发布成功。
	// FAILED：帖子发布出现错误，帖子发布失败。
	// SEND_TO_USER_INBOX：帖子草稿上传成功，系统已将相应通知发至创作者的收件箱中。
	Status enum.PublishStatus `json:"status,omitempty"`
	// Reason 仅当 status 为 FAILED 时返回本字段。
	// 发布失败的原因。
	// 示例： frame_rate_check_failed。
	// 若想获取失败原因的具体描述，可查看失败原因。
	Reason string `json:"reason,omitempty"`
}

func (s PublishStatus) IsError() bool {
	return s.Status == enum.PublishStatus_FAILED
}

func (s PublishStatus) Error() string {
	return s.Reason
}
