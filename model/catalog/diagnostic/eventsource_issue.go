package diagnostic

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// EventSourceIssueRequest 获取商品库事件源诊断信息
type EventSourceIssueRequest struct {
	// BcID 商务中心ID。该商务中心需为所指定商品库的所属商务中心，或所指定商品库已作为资产分享给该商务中心
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID。您需有访问该商品库的权限。
	CatalogID string `json:"catalog_id,omitempty"`
	// EventSourceType 事件源类型。
	// 枚举值：
	// APP: App。需同时传入参数 app_id。
	// PIXEL: Pixel。需同时传入参数 pixel_code。
	EventSourceType enum.EventSourceType `json:"event_source_type,omitempty"`
	// AppID event_source_type 为 APP 时必填。
	// App ID。
	// 若想获取与商品库（catalog_id）绑定的 App ID，可使用 /catalog/eventsource_bind/get/。
	AppID string `json:"app_id,omitempty"`
	// PixelCode event_source_type 为 PIXEL 时必填。
	// Pixel 代码。
	// 若想获取与商品库（catalog_id）绑定的 Pixel 代码，可使用 /catalog/eventsource_bind/get/。
	PixelCode string `json:"pixel_code,omitempty"`
	// EventType 您想要获取对应数据的事件类型。
	// 枚举值：
	// VIEW_CONTENT：查看详情。
	// ADD_TO_CART：加入购物车。
	// PURCHASE：支付完成。
	// 默认值： VIEW_CONTENT。
	EventType enum.EventType `json:"event_type,omitempty"`
	// TimeRange 您想要获取对应数据的时间范围。
	// 枚举值：
	// YESTERDAY：近1天。
	// LAST_7_DAYS：近7天。
	// LAST_30_DAYS：近30天。
	// 默认值：LAST_7_DAYS
	TimeRange string `json:"time_range,omitempty"`
}

// Encode implements GetRequest interface
func (r *EventSourceIssueRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("catalog_id", r.CatalogID)
	values.Set("bc_id", r.BcID)
	values.Set("event_source_type", string(r.EventSourceType))
	if r.AppID != "" {
		values.Set("app_id", r.AppID)
	}
	if r.PixelCode != "" {
		values.Set("pixel_code", r.PixelCode)
	}
	if r.EventType != "" {
		values.Set("event_type", string(r.EventType))
	}
	if r.TimeRange != "" {
		values.Set("time_range", r.TimeRange)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// EventSourceIssueResponse 获取商品库事件源诊断信息
type EventSourceIssueResponse struct {
	model.BaseResponse
	Data struct {
		// List 诊断信息列表
		List []EventSourceIssue `json:"list,omitempty"`
	} `json:"data"`
}

// EventSourceIssue 诊断信息
type EventSourceIssue struct {
	// DiagnosticResult 诊断出的问题。
	// 若未检测到问题，本字段的值将为"No Issues Found"。
	// 示例："You have fewer View Content events than Add to Cart events."
	// 若想了解可检测的问题，请查看可检测的商品库事件源问题列表。
	DiagnosticResult string `json:"diagnostic_result,omitempty"`
	// Level 问题级别。
	// 枚举值：
	// ERROR：错误。此问题影响广告投放，请解决。
	// WARNING：提示。此问题可能会影响广告投放，请检查。
	// INFO：无问题。
	Level enum.CatalogDiagnosticIssueLevel `json:"level,omitempty"`
	// DiagnosticSolution 问题的推荐解决办法。
	// 若未检测到问题，本字段将为空值。
	DiagnosticSolution string `json:"diagnostic_solution,omitempty"`
}
