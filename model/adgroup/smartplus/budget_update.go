package smartplus

import (
	"github.com/bububa/tiktok-business/model/adgroup"
	"github.com/bububa/tiktok-business/util"
)

// BudgetUpdateRequest Update the budgets of Upgraded Smart+ Ad Groups API Request
type BudgetUpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Budget Information about the new budgets that you want to set for one or more ad groups.
	// Max size: 20.
	Budget []adgroup.Budget `json:"budget,omitempty"`
}

// Encode implements PostRequest interface
func (r BudgetUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
