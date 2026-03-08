package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdvertiserBalanceGetRequest Get the balance and budget of ad accounts API Request
type AdvertiserBalanceGetRequest struct {
	// BcID 商务中心 ID
	BcID string `json:"bc_id,omitempty"`
	// Fields A list of additional fields to return in the response.
	// Supported values:
	// budget_remaining: The remaining budget.
	// budget_frequency_restriction: Restrictions on the number of budget changes allowed for the current day.
	// budget_amount_restriction: Restrictions on the minimum amount that can be changed for the budget.
	// min_transferable_amount: Details of the minimal amount that you can transfer to the ad account.
	// max_transferable_amount: Details of the maximum amounts that you can transfer from the ad account to the Business Center.
	// balance_info: Details of the balances for the ad account.
	// If this field is not specified, all information will be returned by default excluding budget_remaining, budget_frequency_restriction, min_transferable_amount, max_transferable_amount, and balance_info.
	// If you want to retrieve the additional fields (budget_remaining, budget_frequency_restriction, budget_amount_restriction, min_transferable_amount, max_transferable_amount, and balance_info) together with other fields, set fields to ["budget_remaining", "budget_frequency_restriction", "budget_amount_restriction", "min_transferable_amount", "max_transferable_amount", "balance_info"] and page_size to 1.
	Fields []string `json:"fields,omitempty"`
	// Filtering Filtering conditions
	Filtering *AdvertiserBalanceGetFilter `json:"filtering,omitempty"`
	// Page Current number of pages.
	// Default value: 1. Value range : ≥ 1.
	Page int `json:"page,omitempty"`
	// PageSize Page size.
	// Default value: 10. Value range: 1-50.
	PageSize int `json:"page_size,omitempty"`
}

type AdvertiserBalanceGetFilter struct {
	// Keyword Keywords, you can search for ad account name or ad account ID.
	Keyword string `json:"keyword,omitempty"`
	// AdvertiserStatus Ad Account display status.
	// Enum values：
	// SHOW_ACCOUNT_STATUS_NOT_APPROVED: failed.
	// SHOW_ACCOUNT_STATUS_APPROVED: passed.
	// SHOW_ACCOUNT_STATUS_IN_REVIEW: under review.
	// SHOW_ACCOUNT_STATUS_PUNISHED: punishment.
	AdvertiserStatus []enum.AdvertiserStatus `json:"advertiser_status,omitempty"`
	// PaymentPortfolioID Valid when you have enabled the multiple Payment Portfolios feature.
	// The ID of the Payment Portfolio to filter the results by.
	// Note:
	// Multiple Payment Portfolios for one client is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative.
	// If this field is specified, the balance information in the response will be for the ad account linked to the Payment Portfolio (payment_portfolio_id).
	PaymentPortfolioID string `json:"payment_portfolio_id,omitempty"`
}

// Encode implements GetRequest
func (r *AdvertiserBalanceGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
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

// AdvertiserBalanceGetResponse Get the balance and budget of ad accounts API Response
type AdvertiserBalanceGetResponse struct {
	model.BaseResponse
	Data *AdvertiserBalanceGetResult `json:"data,omitempty"`
}

type AdvertiserBalanceGetResult struct {
	// PageInfo Pagination information.
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AdvertiserAccountList Advertiser account balance list.
	// Note: If the ad account belongs to a Business Center that has enabled Payment Portfolio, the following parameters will be for the Payment Portfolio rather than the ad account:
	// account_balance
	// valid_account_balance
	// frozen_balance
	// tax
	// cash_balance
	// valid_cash_balance
	// transferable_amount
	AdvertiserAccountList []AdvertiserAccount `json:"advertiser_account_list,omitempty"`
}

type AdvertiserAccount struct {
	// AdvertiserID Advertiser account ID.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserName Advertiser account name.
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// AdvertiserStatus Advertiser account display status.
	// Enum values：
	// SHOW_ACCOUNT_STATUS_NOT_APPROVED: failed.
	// SHOW_ACCOUNT_STATUS_APPROVED: passed.
	// SHOW_ACCOUNT_STATUS_IN_REVIEW: under review.
	// SHOW_ACCOUNT_STATUS_PUNISHED: punishment.
	AdvertiserStatus enum.AdvertiserStatus `json:"advertiser_status,omitempty"`
	// AdvertiserType Advertiser account type.
	// Enum values：AUCTION(auction), RESERVATION(reservation).
	AdvertiserType enum.ServiceType `json:"advertiser_type,omitempty"`
	// Timezone Advertiser account time zone. For enum values, see Appendix-Time Zone.
	Timezone string `json:"timezone,omitempty"`
	// Currency Advertiser account currency. For enum values, see Appendix-Currency.
	Currency string `json:"currency,omitempty"`
	// AccountOpenDays Advertiser account opening days
	AccountOpenDays int `json:"account_open_days,omitempty"`
	// BalanceReminder Balance hit the line, when the balance hit the line is yes, it means that you should add money to the main account of the advertisement.
	BalanceReminder bool `json:"balance_reminder,omitempty"`
	// Company Advertiser account company name
	Company string `json:"company,omitempty"`
	// ContactName Advertiser account contact name
	ContactName string `json:"contact_name,omitempty"`
	// ContactEmail Advertiser account contact email.
	ContactEmail string `json:"contact_email,omitempty"`
	// CreateTime Advertiser account opening time (UTC+0).
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// AccountBalance Advertiser account total balance, kept to two decimal places.
	AccountBalance float64 `json:"account_balance,omitempty"`
	// ValidAccountBalance Advertiser account valid account balance, kept to two decimal places.
	ValidAccountBalance float64 `json:"valid_account_balance,omitempty"`
	// FrozenBalance Advertiser account frozen balance, kept to two decimal places.
	FrozenBalance float64 `json:"frozen_balance,omitempty"`
	// Tax Advertiser account taxes, kept to two decimal places. The default value is 0
	Tax float64 `json:"tax,omitempty"`
	// CashBalance Advertiser account cash balance, kept to two decimal places.
	CashBalance float64 `json:"cash_balance,omitempty"`
	// ValidCashBalance Advertiser account valid cash balance, kept to two decimal places
	ValidCashBalance float64 `json:"valid_cash_balance,omitempty"`
	// GrantBalance Advertiser account coupon/voucher balance, kept to two decimal places
	GrantBalance float64 `json:"grant_balance,omitempty"`
	// ValidGrantBalance Advertiser account valid coupon/voucher balance, kept to two decimal places.
	ValidGrantBalance float64 `json:"valid_grant_balance,omitempty"`
	// PaymentPortfolioID Returned only for a Business Center that has enabled Payment Portfolio.
	// The ID of the Payment Portfolio linked to the ad account.
	PaymentPortfolioID string `json:"payment_portfolio_id,omitempty"`
	// PaymentPortfolioName Returned only for a Business Center that has enabled Payment Portfolio.
	// The name of the Payment Portfolio linked to the ad account.
	PaymentPortfolioName string `json:"payment_portfolio_name,omitempty"`
	// TransferableAmount Amount that you can transfer from the advertiser account, based on the currency used in your advertiser account, and kept to two decimal places.
	// Note:
	// The transferable amount excludes the frozen balance (frozen_balance) and the deposit for ad delivery. Therefore, the transferable_amount might be lower than both valid_cash_balance and valid_account_balance. To ensure uninterrupted ad delivery, we recommend transferring an amount not greater than the transferable_amount.
	// This field will only be returned when total_number is 1.
	TransferableAmount float64 `json:"transferable_amount,omitempty"`
	// MaxTransferableAmount Details of the maximum amount that you can transfer from the ad account to the Business Center.
	MaxTransferableAmount *TransferAmountInfo `json:"max_transferable_amount,omitempty"`
	// BalanceInfo Details of the balances for the ad account
	BalanceInfo *AdvertiserAccountBalanceInfo `json:"balance_info,omitempty"`
	// BudgetMode Budget mode of the ad account.
	// Enum values:
	// UNLIMITED: No budget limit for the ad account.
	// MONTHLY_BUDGET: The ad account consumes BC credit balance within the monthly budget.
	// DAILY_BUDGET: The ad account consumes BC credit balance within the daily budget.
	// CUSTOM_BUDGET: The ad account consumes BC credit balance within the one-time custom budget.
	// Note:
	// If the BC is in manual allocation mode, the returned value is always UNLIMITED.
	// The default budget mode for each ad account is UNLIMITED unless you customize your budget mode.
	BudgetMode enum.AdvertiserBudgetMode `json:"budget_mode,omitempty"`
	// Budget If budget_mode=MONTHLY_BUDGET, the monthly budget of the advertiser account will be returned.
	// If budget_mode=DAILY_BUDGET, the daily budget of the advertiser account will be returned.
	// If budget_mode=CUSTOM_BUDGET, the one-time custom budget of the advertiser account will be returned.
	// If budget_mode=UNLIMITED, returns 0.
	Budget float64 `json:"budget,omitempty"`
	// BudgetCost If budget_mode is MONTHLY_BUDGET, DAILY_BUDGET, or CUSTOM_BUDGET, spent budget of the ad account will be returned.
	// If budget_mode=UNLIMITED, returns 0.
	BudgetCost float64 `json:"budget_cost,omitempty"`
	// BudgetRemaining The remaining budget.
	// Formula: budget_remaining = budget - budget_cost.
	BudgetRemaining float64 `json:"budget_remaining,omitempty"`
	// BudgetFrequencyRestriction Restrictions on the number of budget changes allowed for the current day.
	// You can make up to 10 updates per day, including both budget mode and budget amount changes or only budget amount changes for the ad account.
	BudgetFrequencyRestriction *BudgetFrequencyRestriction `json:"budget_frequency_restriction,omitempty"`
	// BudgetAmountRestriction Restrictions on the minimum amount that can be changed for the budget.
	BudgetAmountRestriction *BudgetAmountRestriction `json:"budget_amount_restriction,omitempty"`
	// MinTransferableAmount Returned only when page_size is set to 1 in the request.
	// Details of the minimal amount that you can transfer to the ad account.
	MinTransferableAmount *TransferAmountInfo `json:"min_transferable_amount,omitempty"`
}

type AdvertiserAccountBalanceInfo struct {
	// AccountBalance The total balance of the ad account, kept to two decimal places.
	// Formula: account_balance = valid_account_balance + frozen_balance + tax.
	AccountBalance float64 `json:"account_balance,omitempty"`
	// ValidAccountBalance Advertiser account valid account balance, kept to two decimal places.
	ValidAccountBalance float64 `json:"valid_account_balance,omitempty"`
	// FrozenBalance Advertiser account frozen balance, kept to two decimal places.
	FrozenBalance float64 `json:"frozen_balance,omitempty"`
	// Tax Advertiser account taxes, kept to two decimal places. The default value is 0
	Tax float64 `json:"tax,omitempty"`
	// CashBalance Advertiser account cash balance, kept to two decimal places.
	CashBalance float64 `json:"cash_balance,omitempty"`
	// FrozenCashBalance The frozen cash balance of the ad account, kept to two decimal places.
	FrozenCashBalance float64 `json:"frozen_cash_balance,omitempty"`
	// CashTax The cash taxes of the ad account, kept to two decimal places
	CashTax float64 `json:"cash_tax,omitempty"`
	// ValidCashBalance Advertiser account valid cash balance, kept to two decimal places
	ValidCashBalance float64 `json:"valid_cash_balance,omitempty"`
	// GrantBalance Advertiser account coupon/voucher balance, kept to two decimal places
	GrantBalance float64 `json:"grant_balance,omitempty"`
	// ValidGrantBalance Advertiser account valid coupon/voucher balance, kept to two decimal places.
	ValidGrantBalance float64 `json:"valid_grant_balance,omitempty"`
	// FrozenGrantBalance The frozen voucher balance of the ad account, kept to two decimal places.
	FrozenGrantBalance float64 `json:"frozen_grant_balance,omitempty"`
	// CreditBalance The total credit balance of the ad account, kept to two decimal places.
	CreditBalance float64 `json:"credit_balance,omitempty"`
	// ValidCreditBalance The available credit balance of the ad account, kept to two decimal places.
	ValidCreditBalance float64 `json:"valid_credit_balance,omitempty"`
	// FrozenCreditBalance The frozen credit balance of the ad account, kept to two decimal places.
	FrozenCreditBalance float64 `json:"frozen_credit_balance,omitempty"`
}

// BudgetFrequencyRestriction Restrictions on the number of budget changes allowed for the current day.
type BudgetFrequencyRestriction struct {
	// TotalCount The maximum number of updates allowed.
	TotalCount int `json:"total_count,omitempty"`
	// UsedCount The number of updates that have been made
	UsedCount int `json:"used_count,omitempty"`
	// RemainingCount The remaining number of updates allowed
	RemainingCount int `json:"remaining_count,omitempty"`
	// EffectiveStartTime The time when the restriction starts, in the format of YYYY-MM-DD HH:MM:SS (UTC+0 time).
	EffectiveStartTime model.DateTime `json:"effective_start_time,omitzero"`
	// EffectiveEndTime The time when the restriction ends, in the format of YYYY-MM-DD HH:MM:SS (UTC+0 time).
	EffectiveEndTime model.DateTime `json:"effective_end_time,omitzero"`
}

// BudgetAmountRestriction Restrictions on the minimum amount that can be changed for the budget.
type BudgetAmountRestriction struct {
	// MinimumAmount The minimum amount that you can change for the budget.
	// The amount is based on the budget that has been consumed during the budget period. The change amount should be equal to or greater than this value.
	MinimumAmount float64 `json:"minimum_amount,omitempty"`
}
