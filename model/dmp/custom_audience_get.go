package dmp

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CustomAudienceGetRequest 获取受众详情 API Request
type CustomAudienceGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceIDs 受众ID列表。长度范围：[1, 100]
	CustomAudienceIDs []string `json:"custom_audience_ids,omitempty"`
	// HistorySize 用户请求的历史数据的大小
	HistorySize int `json:"history_size,omitempty"`
}

// Encode implements GetRequest
func (r *CustomAudienceGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("custom_audience_ids", string(util.JSONMarshal(r.CustomAudienceIDs)))
	if r.HistorySize > 0 {
		values.Set("history_size", strconv.Itoa(r.HistorySize))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// CustomAudienceGetResponse 获取受众详情 API Response
type CustomAudienceGetResponse struct {
	model.BaseResponse
	Data struct {
		// List 受众群体详情列表
		List []CustomAudienceWithHistory `json:"list,omitempty"`
	} `json:"data"`
}

type CustomAudienceWithHistory struct {
	// AudienceDetails 受众群体数据
	AudienceDetails *CustomAudience `json:"audience_details,omitempty"`
	// AudienceHistory 受众群体变更历史列表
	AudienceHistory []AudienceHistory `json:"audience_history,omitempty"`
}

// AudienceHistory 受众群体变更历史
type AudienceHistory struct {
	// Action 变更行为。欲了解枚举值含义，请查看下文"action 值与通过/dmp/custom_audience/update/的操作的映射关系" 小节
	Action string `json:"action,omitempty"`
	// ActionDetail 操作详情
	ActionDetail string `json:"action_detail,omitempty"`
	// Editor 操作者。如果是通过TikTok API for Business上传的，那么值是"open_api"
	Editor string `json:"editor,omitempty"`
	// Msg 部分成功时的失败信息。比如：文件中部分数据解析失败。
	Msg string `json:"msg,omitempty"`
	// OptTime 操作时间。格式如：2019-04-14 06:52:56
	OptTime model.DateTime `json:"opt_time,omitzero"`
}
