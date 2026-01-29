package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest Update an Upgraded Smart+ Campaign API Request
type UpdateRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID Campaign ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName Campaign name.
	// Length limit: 512 characters. Emoji is not supported.
	// Each word in Chinese or Japanese counts as two characters, while each letter in English counts as one character
	CampaignName string `json:"campaign_name,omitempty"`
	// BudgetAutoAdjustStrategy Valid only when the following conditions are all met:
	// budget_optimize_on is true or not specified.
	// budget_mode is BUDGET_MODE_DYNAMIC_DAILY_BUDGET.
	// app_promotion_type is not MINIS.
	// The campaign budget strategy for automatic daily campaign budget.
	// Enum value:
	// AUTO_BUDGET_INCREASE: To enable automatic budget increase. Allow your budget to automatically increase when your ads are performing well and target CPA, Day 0 target ROAS, and budget requirements are met.
	// When budget_auto_adjust_strategy is AUTO_BUDGET_INCREASE, the specified budget will be the initial daily budget. Your daily budget will be allowed to automatically increase by 20%, up to 10 times per day, when your budget utilization reaches 90% or more. Your daily budget will reset to your original daily budget each day.
	// Note: Enabling automatic budget increase is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative.
	BudgetAutoAdjustStrategy enum.BudgetAutoAdjustStrategy `json:"budget_auto_adjust_strategy,omitempty"`
	// Buget Required when budget_mode is specified.
	// Fixed campaign budget or initial campaign budget.
	// When budget_auto_adjust_strategy is UNSET, this field represents your fixed campaign budget.
	// When budget_auto_adjust_strategy is AUTO_BUDGET_INCREASE, this field represents your initial campaign budget. To retrieve the current campaign budget, check the returned current_budget.
	// To learn about the minimum budget and how to set budget modes, see Budget.
	// To directly see the daily budget value range for a currency, see Currency-Daily budget value range.
	Budget float64 `json:"budget,omitempty"`
	// PONumber PO (purchase order) number.
	// PO numbers are useful for invoice tracking and reconciliation.
	// Learn more about PO numbers on monthly invoices.
	PONumber string `json:"po_number,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse Update an Upgraded Smart+ Campaign API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *Campaign `json:"data,omitempty"`
}
