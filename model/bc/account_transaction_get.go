package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AccountTransactionGetRequest Get the transaction records of a BC or ad accounts API Request
type AccountTransactionGetRequest struct {
	// BcID 商务中心 ID
	BcID string `json:"bc_id,omitempty"`
	// TransactionLevel The transaction level that the transaction records will be retrieved at.
	// Enum values:
	// BC: The Business Center (bc_id).
	// ADVERTISER: The ad accounts within the Business Center (bc_id).
	// PAYMENT_PORTFOLIO: The Payment Portfolio.
	// Default value: ADVERTISER.
	// Note: If a Business Center has enabled Payment Portfolio, the transactions of ad accounts shared by other Business Centers won't be returned.
	TransactionLevel enum.BcTransferLevel `json:"transaction_level,omitempty"`
	// Filtering Filtering conditions.
	Filtering *AccountTransactionGetFilter `json:"filtering,omitempty"`
	// Page Current number of pages.
	// Default value: 1.
	// Value range: ≥ 1.
	Page int `json:"page,omitempty"`
	// PageSize Page size.
	// Default value: 10.
	// Value range: 1-50.
	PageSize int `json:"page_size,omitempty"`
}

type AccountTransactionGetFilter struct {
	// TransactionTypes Transaction types.
	// BILL_PAYMENT: Bill payment.
	// ADD_BALANCE: Add balance.
	// CANCELLATION: Cancellation.
	// PROMOTION_ISSUED: Issued promotion, which indicates the issuance of ad credit.
	// PROMOTION_EXPIRED: Expired promotion, which indicates the expiration of ad credit.
	// INCREASE_BALANCE: Increase of balance.
	// DECREASE_BALANCE: Decrease of balance.
	// CARD_VERIFICATION: Card verification.
	// CARD_VERIFICATION_REFUND: Card verification refund.
	// ORDER_CREATION_PAYMENT: Order creation payment.
	// ORDER_EDIT_PAYMENT: Order edit payment.
	// RETURN_OF_REMAINING_BUDGET: Return of remaining budget.
	// REFUND: Refund.
	TransactionTypes []string `json:"transaction_types,omitempty"`
	// OwnerOfAccount Valid only when transaction_level is ADVERTISER or not passed.
	// Owner of ad accounts.
	// Enum values:
	// OWNED_BY_CURRENT_BUSINESS_CENTER: Owned by the current Business Center.
	// SHARED_BY_OTHER_BUSINESS_CENTER: Shared by another Business Center.
	// Default value: OWNED_BY_CURRENT_BUSINESS_CENTER.
	// If you want to obtain the transaction records of ad accounts shared by another Business Center, you need to have Admin access to these accounts. To confirm if you have Admin access to these accounts, use /bc/asset/get/. The returned advertiser_role for the asset_id with asset_type as ADVERTISER should be ADMIN.
	// Note: If your Business Center has more than 200 ad accounts shared by another Business Center, you cannot set this field to SHARED_BY_OTHER_BUSINESS_CENTER to filter shared ad accounts. In such cases, we recommend that you specify the account_ids filter or both account_ids and account_name filters simultaneously to narrow down the scope of your query.
	OwnerOfAccount []string `json:"owner_of_account,omitempty"`
	// AccountIDs Valid only when transaction_level is ADVERTISER or not passed.
	// A list of ad account IDs.
	// Max size: 100.
	AccountIDs []string `json:"account_ids,omitempty"`
	// AccountName Valid only when transaction_level is ADVERTISER or not passed.
	// Ad account name.
	// Fuzzy match is supported.
	AccountName string `json:"account_name,omitempty"`
	// StartTime Query start time, closed interval, in the format of YYYY-MM-DD HH:MM:SS (ad account time zone or Business Center time zone).
	// Example: 2024-01-01 00:00:00.
	// The maximum query time range that you can define through start_date and end_date is 365 days.
	// If start_time and end_time are not passed, the transaction records of the last 90 days will be returned by default.
	StartTime string `json:"start_time,omitempty"`
	// EndTime Query end time, closed interval, in the format of YYYY-MM-DD HH:MM:SS (ad account time zone or Business Center time zone).
	// Example: 2024-01-01 00:00:00.
	// The maximum query time range that you can define through start_time and end_time is 365 days.
	// If start_time and end_time are not passed, the transaction records of the last 90 days will be returned by default.
	EndTime string `json:"end_time,omitempty"`
	// BillingTypes Billing types.
	// Enum values:
	// CASH: cash.
	// CREDIT: ad credit.
	BillingTypes []enum.BillingType `json:"billing_types,omitempty"`
}

// Encode implements GetRequest
func (r *AccountTransactionGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc", r.BcID)
	if r.TransactionLevel != "" {
		values.Set("transaction_level", string(r.TransactionLevel))
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

// AccountTransactionGetResponse Get the transaction records of a BC or ad accounts API Response
type AccountTransactionGetResponse struct {
	model.BaseResponse
	Data *AccountTransactionGetResult `json:"data,omitempty"`
}

type AccountTransactionGetResult struct {
	// PageInfo Pagination information
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// TransactionList The list of transaction records.
	TransactionList []Transaction `json:"transaction,omitempty"`
}

type Transaction struct {
	// TransactionID Transaction ID, the unique identifier of the translation
	TransactionID string `json:"transaction_id,omitempty"`
	// PaymentPortfolioID The ID of the Payment Portfolio associated with the transaction.
	// This field returns a non-null value only when transaction_level is PAYMENT_PORTFOLIO.
	// Note: To manage the budget of a Payment Portfolio, use advertiser_budgets in /advertiser/update/. The default budget_mode within advertiser_budgets will be CUSTOM.
	PaymentPortfolioID string `json:"payment_portfolio_id,omitempty"`
	// PaymentPortfolioName The name of the Payment Portfolio associated with the transaction.
	// This field returns a non-null value only when transaction_level is PAYMENT_PORTFOLIO.
	PaymentPortfolioName string `json:"payment_portfolio_name,omitempty"`
	// AccountID The ID of the ad account associated with the transaction.
	// When transaction_level is set to BC in the request, the value of this field will be null.
	AccountID string `json:"account_id,omitempty"`
	// AccountName The name of the ad account associated with the transaction.
	// When transaction_level is set to BC in the request, the value of this field will be null.
	AccountName string `json:"account_name,omitempty"`
	// BcID Business Center ID.
	BcID string `json:"bc_id,omitempty"`
	// BcName Business Center name.
	BcName string `json:"bc_name,omitempty"`
	// Amount The total amount of the transaction, which includes the subtotal and tax amount.
	// For example, if the amount excluding tax is 50 USD and the tax is 5 USD, the total amount will be 55 USD.
	Amount float64 `json:"amount,omitempty"`
	// Subtotal The subtotal amount (the amount excluding taxes) of the transaction.
	Subtotal float64 `json:"subtotal,omitempty"`
	// TaxAmount The amount of tax included in the transaction.
	TaxAmount float64 `json:"tax_amount,omitempty"`
	// Currency The currency code in which the amounts (amount, subtotal, and tax_amount) are calculated.
	Currency string `json:"currency,omitempty"`
	// AmountType Amount type.
	// Enum values:
	// POSITIVE: This type represents that a positive amount is associated with the transaction, such as a balance addition.
	// NEGATIVE: This type represents that a negative amount is associated with the transaction, such as an order creation payment.
	// OTHER: This type represents that neither a positive amount nor a negative amount is associated with the transaction, such as a card verification refund.
	AmountType string `json:"amount_type,omitempty"`
	// TransactionType Transaction type.
	// BILL_PAYMENT: Bill payment.
	// ADD_BALANCE: Add balance.
	// CANCELLATION: Cancellation.
	// PROMOTION_ISSUED: Issued promotion, which indicates the issuance of ad credit.
	// PROMOTION_EXPIRED: Expired promotion, which indicates the expiration of ad credit.
	// INCREASE_BALANCE: Increase of balance.
	// DECREASE_BALANCE: Decrease of balance.
	// CARD_VERIFICATION: Card verification.
	// CARD_VERIFICATION_REFUND: Card verification refund.
	// ORDER_CREATION_PAYMENT: Order creation payment.
	// ORDER_EDIT_PAYMENT: Order edit payment.
	// RETURN_OF_REMAINING_BUDGET: Return of remaining budget.
	// REFUND: Refund.
	TransactionType string `json:"transaction_type,omitempty"`
	// BillingType Billing type.
	// Enum values:
	// CASH: cash.
	// CREDIT: ad credit.
	BillingType enum.BillingType `json:"billing_type,omitempty"`
	// Timezone The time zone where the transaction occurred, in the format of UTC±HH:MM.
	Timezone string `json:"timezone,omitempty"`
	// CreateTime The time when the transaction record was generated, in the format of "YYYY-MM-DD HH:MM:SS"
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// InvoiceID Invoice ID.
	// To download the invoice synchronously, pass the invoice ID to /bc/invoice/download/.
	// To download the invoice asynchronously, pass the invoice ID to /bc/invoice/task/create/.
	InvoiceID string `json:"invoice_id,omitempty"`
	// SerialNumber Invoice serial number.
	SerialNumber string `json:"serial_number,omitempty"`
}
