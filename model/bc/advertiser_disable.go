package bc

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdvertiserDisableRequest 永久停用广告账户 API Request
type AdvertiserDisableRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// AdvertiserIDs 您想要永久停用的商务中心内广告账户的ID。最大数量为1。
	// 重要提示: 商务中心内广告账户停用后，广告账户的账户余额会转入所归属的商务中心中
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
}

// Encode implements PostRequest
func (r *AdvertiserDisableRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdvertiserDisableResponse 永久停用广告账户 API Response
type AdvertiserDisableResponse struct {
	model.BaseResponse
	Data *AdvertiserDisableResult `json:"data,omitempty"`
}

type AdvertiserDisableResult struct {
	// DisabledAdvertiserIDs 已成功永久停用的广告账户的ID
	DisabledAdvertiserIDs []string `json:"disabled_advertiser_ids,omitempty"`
	// FailedInfos 未成功永久停用的广告账户的相关信息。
	// 键：未成功永久停用的广告账户的ID（advertiser_id）。
	// 值：该广告账户未成功永久停用的原因。枚举值：
	// DELIVERING：广告账户下三天内有在投广告。
	// UNPAID_BILL：广告账户下有未支付账单。
	// SUSPENDED：广告账户处于暂时禁用状态。
	// UNFINISHED_TRANSFER：广告账户下有未完成的转账单。
	// AUTOPAY_UNBILLED：该广告账户的付款类型为自动付款，且该账户下有未入账的消耗。
	FailedInfos map[string]string `json:"failed_infos,omitempty"`
}
