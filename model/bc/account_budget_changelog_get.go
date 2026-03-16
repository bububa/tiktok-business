package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AccountBudgetChangelogGetRequest Get the budget change history of an ad account API Request
type AccountBudgetChangelogGetRequest struct {
	// BcID Business Center ID
	BcID string `json:"bc_id,omitempty"`
	// AdvertiserID Ad account ID.
	// To obtain the list of ad accounts that you have access to, use /bc/asset/get/. Set asset_type to ADVERTISER and select ad accounts with the returned advertiser_role as ADMIN, OPERATOR, or ANALYST.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Filtering Filtering conditions
	Filtering *AccountBudgetChangelogGetFilter `json:"filtering,omitempty"`
	// Page Current number of pages. Default value: 1. Value range : ≥ 1.
	Page int `json:"page,omitempty"`
	// PageSize Page size. Default value: 10. Value range: 1-50.
	PageSize int `json:"page_size,omitempty"`
}

type AccountBudgetChangelogGetFilter struct {
	// StartDate Query start date, in the format of YYYY-MM-DD (UTC+0 time).
	// You can either specify start_date and end_date simultaneously or leave both start_date and end_date unspecified.
	// If you specify start_date and end_date simultaneously, the maximum time range is 365 days.
	// If you leave both start_date and end_date unspecified, the results for the last seven days will be returned by default.
	StartDate string `json:"start_date,omitempty"`
	// EndDate Query end date, in the format of YYYY-MM-DD (UTC+0 time).
	// You can either specify start_date and end_date simultaneously or leave both start_date and end_date unspecified.
	// If you specify start_date and end_date simultaneously, the maximum time range is 365 days.
	// If you leave both start_date and end_date unspecified, the results for the last seven days will be returned by default.
	EndDate string `json:"end_date,omitempty"`
}

// Encode implements GetRequest
func (r *AccountBudgetChangelogGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("advertiser_id", r.AdvertiserID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AccountBudgetChangelogGetResponse Get the budget change history of an ad account API Response
type AccountBudgetChangelogGetResponse struct {
	model.BaseResponse
	Data *AccountBudgetChangelogGetResult `json:"data,omitempty"`
}

type AccountBudgetChangelogGetResult struct {
	// PageInfo Pagination information
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ChangelogList The budget change history details.
	ChangelogList []AccountBudgetChangelog `json:"changelog_list,omitempty"`
}

type AccountBudgetChangelog struct {
	// OperationTime The time when the budget change occurred, in the format of YYYY-MM-DD HH:MM:SS (UTC+0 time).
	OperationTime model.DateTime `json:"operation_time,omitzero"`
	// ActivityType Change type.
	// Enum values:
	// RESET: Reset the one-time budget consumption of the ad account to 0.
	// BUDGET_MODE_UPDATE: Change the budget mode to monthly, daily, or one-time.
	// INCREASE_BUDGET: Increase budget without changing budget mode.
	// DECREASE_BUDGET: Decrease budget without changing budget mode.
	// REMOVE_BUDGET: Change the budget mode to unlimited.
	ActivityType string `json:"activity_type,omitempty"`
	// PreviousBudget The previous budget amount of the ad account before the change, kept to two decimal places
	PreviousBudget float64 `json:"previous_budget,omitempty"`
	// PreviousBudgetMode The previous budget mode of the ad account.
	// Enum values:
	// UNLIMITED: Unlimited budget. No budget limit for the ad account.
	// MONTHLY_BUDGET: Monthly budget. The ad account consumes Business Center credit line within the monthly budget.
	// DAILY_BUDGET: Daily budget. The ad account consumes Business Center credit line within the daily budget.
	// CUSTOM_BUDGET: One-time budget (custom budget). The ad account consumes Business Center credit line within the one-time custom budget.
	PreviousBudgetMode enum.AdvertiserBudgetMode `json:"previous_budget_mode,omitempty"`
	// CurrentBudget The current budget amount of the ad account after the change, kept to two decimal places.
	CurrentBudget float64 `json:"current_budget,omitempty"`
	// CurrentBudgetMode The current budget mode of the ad account.
	// Enum values:
	// UNLIMITED: Unlimited budget. No budget limit for the ad account.
	// MONTHLY_BUDGET: Monthly budget. The ad account consumes Business Center credit line within the monthly budget.
	// DAILY_BUDGET: Daily budget. The ad account consumes Business Center credit line within the daily budget.
	// CUSTOM_BUDGET: One-time budget (custom budget). The ad account consumes Business Center credit line within the one-time custom budget.
	CurrentBudgetMode enum.AdvertiserBudgetMode `json:"current_budget_mode,omitempty"`
	// Currency The currency for the current budget, in ISO 4217 code format.
	Currency string `json:"currency,omitempty"`
	// OperatorID The user ID of the operator.
	OperatorID string `json:"operator_id,omitempty"`
	// OperatiorName The user name of the operator.
	OperatorName string `json:"operator_name,omitempty"`
}
