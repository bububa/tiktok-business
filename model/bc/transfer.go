package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TransferRequest Process a payment API Request
type TransferRequest struct {
	// BcID Business Center ID.
	BcID string `json:"bc_id,omitempty"`
	// TransferLevel Transfer level.
	// Enum values:
	//   - ADVERTISER: Recharge money to or deduct money from an ad account.
	//   - BC: Increase or decrease the credit balance of the Business Center.
	// When transfer_level is set to BC, you need to specify the ID of a Business Center that has enabled Monthly Invoicing via bc_id.
	// Default value: ADVERTISER.
	TransferLevel enum.BcTransferLevel `json:"transfer_level,omitempty"`
	// AdvertiserID Required when transfer_level is set to ADVERTISER or not specified.
	// Ad account ID.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// TransferType How you'd like to process payments for an ad account or the Business Center.
	// Enum values:
	// RECHARGE: Recharge money to an ad account or increase the credit balance of a Business Center.
	// REFUND: Deduct money from an ad account or decrease the credit balance of a Business Center.
	TransferType enum.BcTransferType `json:"transfer_type,omitempty"`
	// AmountInfo Valid only when transfer_level is ADVERTISER.
	// You need to specify one of amount_info, cash_amount, or grant_amount.
	// Details of the amount to process.
	// When you specify amount_info, cash_amount and grant_amount outside of amoumt_info will be ignored.
	AmountInfo *TransferAmountInfo `json:"amount_info,omitempty"`
	// CashAmount You need to specify one of amount_info, cash_amount, or grant_amount.
	// If Balance Sharing is enabled and transfer_level is ADVERTISER, this parameter is not supported.
	// The cash or credit amount to process, rounded to two decimal places.
	// When transfer_level is set to ADVERTISER or not specified, this field represents the cash amount to process.
	// When transfer_level is set to BC or not specified, this field represents the credit amount to process.
	// Value range: > 0.
	// Note: If you choose RECHARGE as the transfer type and then specify a cash amount, there will be a limit on the minimum amount that you can transfer from a Business Center to the ad account. To find the minimum amount limit , use /advertiser/balance/get/ and check the returned cash_amount within the object min_transferable_amount.
	CashAmount float64 `json:"cash_amount,omitempty"`
	// GrantAmount You need to specify one of amount_info, cash_amount, or grant_amount.
	// When transfer_level is BC, this parameter is not supported.
	// Voucher amount to process, rounded to two decimal places.
	// Value range: > 0.
	GrantAmount float64 `json:"grant_amount,omitempty"`
	// RequestID Valid only when transfer_level is BC.
	// Request ID.
	// This field supports idempotency to prevent you from sending the same request twice. If you retry requests with the same request ID multiple times, then only one will succeed.
	// It is different from the request_id returned in the response parameters, which is used to uniquely identify an HTTP request.
	// The value should be a string representation of a 64-bit integer number.
	RequestID string `json:"request_id,omitempty"`
}

// TransferAmountInfo details of the amount to process in transfer.
type TransferAmountInfo struct {
	// CashAmount The cash amount to process.
	CashAmount float64 `json:"cash_amount,omitempty"`
	// GrantAmount The voucher amount to process
	GrantAmount float64 `json:"grant_amount,omitempty"`
	// CreditAmount The credit amount to process
	CreditAmount float64 `json:"credit_amount,omitempty"`
}

// Encode implements PostRequest
func (r *TransferRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TransferResponse Process a payment API Response
type TransferResponse struct {
	model.BaseResponse
	Data *TransferResult `json:"data,omitempty"`
}

type TransferResult struct {
	// BcID Business Center ID.
	// When transfer_level is BC, this field will not be returned.
	BcID model.Uint64 `json:"bc_id,omitempty"`
	// AdvertiserID Ad account ID.
	// When transfer_level is BC, this field will not be returned.
	AdvertiserID model.Uint64 `json:"advertiser_id,omitempty"`
	// TransactionInfos Details of the transfers.
	// Note: When transfer_level is ADVERTISER or not specified, since you are transferring an amount between the ad account and a Business Center, you will receive two transaction records, one at the ad account level and one at the Business Center level. In such cases, you can map the transactions to the transaction records returned from /bc/account/transaction/get/ by sending two calls to /bc/account/transaction/get/, one with transaction_level set to BC and one with transaction_level set to ADVERTISER.
	TransactionInfos []TransactionInfo `json:"transaction_infos,omitempty"`
}

// TransactionInfo Details of the transfers.
type TransactionInfo struct {
	// TransactionID Transaction ID of the transfer.
	TransactionID string `json:"transaction_id,omitempty"`
	// TransactionType Transaction type of the transfer.
	// Enum values:
	// CASH: cash.
	// GRANT: voucher.
	// CREDIT: credit.
	TransactionType enum.BillingType `json:"transaction_type,omitempty"`
}
