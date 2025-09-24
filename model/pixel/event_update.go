package pixel

import "github.com/bububa/tiktok-business/util"

// EventUpdateRequest 更新Pixel事件 API Request
type EventUpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CurrencyValue 转化价值
	CurrencyValue string `json:"currency_value,omitempty"`
	// Currency 转化价值对应币种，枚举值: INR(印度卢比),JPY(日元),USD(美元)，默认值 USD
	Currency string `json:"currency,omitempty"`
	// EventID 事件 ID。您可以从/pixel/list/接口返回的event_id字段获得事件ID
	EventID string `json:"event_id,omitempty"`
	// EventName 事件名称
	// 名称不能超过40字符（不区分语言），不能包含emoji。
	EventName string `json:"event_name,omitempty"`
}

// Encode implements PostRequest interface
func (r *EventUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
