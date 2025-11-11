package bc

import (
	"github.com/bububa/tiktok-business/enum"
)

type BusinessCenter struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// Name 商务中心名称
	Name string `json:"name,omitempty"`
	// Company 商务中心公司名称
	Company string `json:"company,omitempty"`
	// Currency 商务中心支付币种，取值范围见附录-币种
	Currency string `json:"currency,omitempty"`
	// RegisteredArea 商务中心注册地代码。取值范围见附录-地区代码
	RegisteredArea string `json:"registered_area,omitempty"`
	// Status 商务中心状态，枚举值：REVIEWING(审核中)，DENY(已拒绝)，ENABLE(已通过)，PUNISH(已禁用)。见 枚举值-商务中心状态
	Status enum.BcStatus `json:"status,omitempty"`
	// Timezone 商务中心所在时区，取值范围见附录-时区信息
	Timezone string `json:"timezone,omitempty"`
	// Type 商务中心类型。枚举值：
	// NORMAL：普通商务中心。
	// DIRECT：非自助直客商务中心。
	// AGENCY：非自助代理商商务中心。
	// SELF_SERVICE：自助直客商务中心。
	// SELF_SERVICE_AGENCY：自助代理商商务中心
	Type enum.BcType `json:"type,omitempty"`
}
