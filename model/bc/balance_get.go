package bc

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// BalanceGetRequest Get the balance of a BC API Request
type BalanceGetRequest struct {
	// BcID Business Center ID
	BcID string `json:"bc_id,omitempty"`
	// PaymentPortfolioID The ID of the Payment Portfolio.
	// To retrieve the list of Payment Portfolios that belong to the same client as the Business Center, use /payment_portfolio/get/.
	PaymentPortfolioID string `json:"payment_portfolio_id,omitempty"`
}

// Encode implements GetRequest
func (r *BalanceGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	if r.PaymentPortfolioID != "" {
		values.Set("payment_portfolio_id", r.PaymentPortfolioID)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// BalanceGetResponse Get the balance of a BC API Response
type BalanceGetResponse struct {
	model.BaseResponse
	Data *BalanceGetResult `json:"data,omitempty"`
}

type BalanceGetResult struct {
	// BcID Business Center ID
	BcID string `json:"bc_id,omitempty"`
	// Currency The currency of the Business Center. For supported currencies, see Appendix-Currency
	Currency string `json:"currency,omitempty"`
	// AccountBalance Business Center total balance, rounded to two decimal places
	AccountBalance float64 `json:"account_balance,omitempty"`
	// ValidAccountBalance Business Center valid account balance, rounded to two decimal places
	ValidAccountBalance float64 `json:"valid_account_balance,omitempty"`
	// FrozenBalance Business Center freeze balance, rounded to two decimal places
	FrozenBalance float64 `json:"frozen_balance,omitempty"`
	// Tax Business Center tax, rounded to two decimal places
	Tax float64 `json:"tax,omitempty"`
	// CashBalance Business Center cash balance, rounded to two decimal places
	CashBalance float64 `json:"cash_balance,omitempty"`
	// ValidCashBalance Business Center valid cash balance, rounded to two decimal places
	ValidCashBalance float64 `json:"valid_cash_balance,omitempty"`
	// GrantBalance Business Center coupon/voucher balance, rounded to two decimal places
	GrantBalance float64 `json:"grant_balance,omitempty"`
	// ValidGrantBalance Business Center valid coupon/voucher balance, rounded to two decimal places
	ValidGrantBalance float64 `json:"valid_grant_balance,omitempty"`
	// PaymentPortfolioType The type of the Payment Portfolio.
	// Enum values:
	// SHARED: Advanced Payment Portfolio. For Advanced Payment Portfolios, funds are centrally managed and jointly used by all linked ad accounts.
	// NON_SHARED: Standard Payment Portfolio. For Standard Payment Portfolios, funds are individually managed for each linked ad account.
	PaymentPortfolioType string `json:"payment_portfolio_type,omitempty"`
}
