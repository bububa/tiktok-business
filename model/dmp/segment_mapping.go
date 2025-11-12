package dmp

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// SegmentMappingRequest 添加或删除受众细分群映射 API Request
type SegmentMappingRequest struct {
	// AdvertiserIDs 广告主ID
	// 用于添加或删除受众细分群映射的广告主ID列表。所有广告主 ID 必须属于同一个企业号
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
	// Action 要进行的操作。
	// 枚举值 ：
	// add：为单一或多个受众添加 UID 映射。
	// delete：为单一或多个受众删除 UID 映射。
	// 默认值： add。
	// 注意: UID 指：
	// 加密的设备 ID（对应id_type为IDFA_MD5，AAID_MD5，IDFA_SHA256或 AAID_SHA256)
	// 外部用户 ID，例如加密邮件地址或加密电话号码（对应id_type为EMAIL_SHA256或PHONE_SHA256）。
	Action string `json:"action,omitempty"`
	// IDSchema 传送受众时所用的 ID 类型。
	// 枚举值：
	// IDFA_MD5: 以 MD5 加密 iOS IDFA。
	// AAID_MD5: 以 MD5 加密 Android AAID 或 GAID。
	// IDFA_SHA256: 以 SHA256 加密 iOS IDFA。
	// AAID_SHA256: 以 SHA256 加密 Android AAID 或 GAID。
	// EMAIL_SHA256: 以 SHA256 加密电子邮件地址。请参阅流式 API 相关提示 - 电子邮件地址标准化部分。
	// PHONE_SHA256: 以 SHA256 加密 E.164格式 的电话号码。举例：+1231234567。
	// 注意:
	// 不允许传入重复的ID类型。
	// ID 类型可按任意顺序传入，且支持将枚举值按小写形式传入。
	IDSchema []enum.IDSchema `json:"id_schema,omitempty"`
	// BatchÎata 本字段包含一个 JSON 格式的数组，其中有一组或多组以下字段。本对象用于添加或删除多 ID 映射（每个用户或设备对应多个受众 ID ）
	BatchData [][]SegmentMappingData `json:"batch_data,omitempty"`
}

type SegmentMappingData struct {
	// ID 用户或设备 ID
	ID string `json:"id,omitempty"`
	// AudienceIDs 用于添加或删除多 ID 映射的受众 ID（audience_id）数组，受众ID必须归属于传入的广告主 ID 列表（advertiser_ids）中的至少一个广告主。
	// 注意： 您可通过/segment/audience/获取audience_id
	AudienceIDs []string `json:"audience_ids,omitempty"`
}

// Encode implements PostRequest
func (r *SegmentMappingRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
