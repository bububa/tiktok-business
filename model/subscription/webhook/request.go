package webhook

import (
	"encoding/json"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Request Webhook Request
type Request struct {
	BaseRequest
	// Entry An array containing an object describing the details of changes.
	Entry []Entity `json:"entry,omitempty"`
}

type BaseRequest struct {
	// RequestID Webhook 请求的唯一 ID
	RequestID string `json:"request_id,omitempty"`
	// Object 订阅对象类型。
	// 有效值：
	// 1：线索
	// 2：广告组的审核状态。
	// 3 ：广告的审核状态。
	// 8：单个广告、广告组中所有广告或广告账户下所有广告的疲劳状态。
	// 10 ：the linking of a TikTok video to a TTO campaign.
	// 11 = the change in any of the reporting metrics spend, impressions, clicks, and video_play_actions for one or more ad accounts.
	Object enum.WebhookObject `json:"object,omitempty"`
	// Time 发送事件通知时的 Unix 时间戳（非触发广告审核状态变化事件的时间）。
	Time model.DateTime `json:"time,omitzero"`
	AdAccountSuspensionEntity
}

type RequestTemplate[T Entity] struct {
	BaseRequest
	// Entry An array containing an object describing the details of changes.
	Entry []T `json:"entry,omitempty"`
}

type Entity interface {
	Entity() enum.SubscribeEntity
}

// UnmarshalJSON implement json Unmarshal interface
func (r *Request) UnmarshalJSON(b []byte) (err error) {
	if err := json.Unmarshal(b, &r.BaseRequest); err != nil {
		return err
	}
	if r.SubscriptionID != "" {
		return nil
	}
	switch r.Object {
	case 1:
		var tmp RequestTemplate[LeadEntity]
		if err := json.Unmarshal(b, &tmp); err != nil {
			return err
		}
		r.Entry = make([]Entity, 0, len(tmp.Entry))
		for _, v := range tmp.Entry {
			r.Entry = append(r.Entry, v)
		}
	case 2:
		var tmp RequestTemplate[AdgroupEntity]
		if err := json.Unmarshal(b, &tmp); err != nil {
			return err
		}
		r.Entry = make([]Entity, 0, len(tmp.Entry))
		for _, v := range tmp.Entry {
			r.Entry = append(r.Entry, v)
		}
	case 3:
		var tmp RequestTemplate[AdEntity]
		if err := json.Unmarshal(b, &tmp); err != nil {
			return err
		}
		r.Entry = make([]Entity, 0, len(tmp.Entry))
		for _, v := range tmp.Entry {
			r.Entry = append(r.Entry, v)
		}
	case 8:
		var tmp RequestTemplate[CreativeFatigureEntity]
		if err := json.Unmarshal(b, &tmp); err != nil {
			return err
		}
		r.Entry = make([]Entity, 0, len(tmp.Entry))
		for _, v := range tmp.Entry {
			r.Entry = append(r.Entry, v)
		}
	case 10:
		var tmp RequestTemplate[TCMVideosEntity]
		if err := json.Unmarshal(b, &tmp); err != nil {
			return err
		}
		r.Entry = make([]Entity, 0, len(tmp.Entry))
		for _, v := range tmp.Entry {
			r.Entry = append(r.Entry, v)
		}
	case 11:
		var tmp RequestTemplate[ReportDataChangeEntity]
		if err := json.Unmarshal(b, &tmp); err != nil {
			return err
		}
		r.Entry = make([]Entity, 0, len(tmp.Entry))
		for _, v := range tmp.Entry {
			r.Entry = append(r.Entry, v)
		}
	case 14:
		var tmp RequestTemplate[APIServiceStatusEntity]
		if err := json.Unmarshal(b, &tmp); err != nil {
			return err
		}
		r.Entry = make([]Entity, 0, len(tmp.Entry))
		for _, v := range tmp.Entry {
			r.Entry = append(r.Entry, v)
		}
	}
	return err
}
