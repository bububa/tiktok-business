package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TransactionGetRequest Get the transaction records of a BC API Request
type TransactionGetRequest struct {
	// BcID Business Center ID
	BcID string `json:"bc_id,omitempty"`
	// Filtering Filtering conditions
	Filtering *TransactionGetFilter `json:"filtering,omitempty"`
	// StartDate Start date of transaction records that you want to get, in the format of "YYYY-MM-DD". The default value is the date 90 days earlier.
	StartDate string `json:"start_date,omitempty"`
	// EndDate End date of transaction records that you want to get, in the format of "YYYY-MM-DD". The default value is the date today.
	EndDate string `json:"end_date,omitempty"`
	// Page Current number of pages. Default value: 1. Value range : ≥ 1.
	Page int `json:"page,omitempty"`
	// PageSize Page size. Default value: 10. Value range: 1-50.
	PageSize int `json:"page_size,omitempty"`
}

type TransactionGetFilter struct {
	// FundsType Fund type. If not specified, both types of fund will be returned. Enum values: FUNDS_TYPE_CASH (cash), FUNDS_TYPE_GRANT (non-cash).
	FundsType enum.FundsType `json:"funds_type,omitempty"`
}

// Encode implements GetRequest
func (r *TransactionGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.StartDate != "" {
		values.Set("start_date", r.StartDate)
	}
	if r.EndDate != "" {
		values.Set("end_date", r.EndDate)
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

// TransactionGetResponse Get the transaction records of a BC API Response
type TransactionGetResponse struct {
	model.BaseResponse
	Data *TransactionGetResult `json:"data,omitempty"`
}

type TransactionGetResult struct {
	// PageInfo Pagination information
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List Total pages of results.
	List []Transaction `json:"list,omitempty"`
	// TransactionSummary Sum of transactions
	TransactionSummary *TransactionSummary `json:"transaction_summary,omitempty"`
}

type Transaction struct {
	// Date Transaction date
	Date int64 `json:"date,omitempty"`
	// Amount Transaction amount.
	Amount float64 `json:"amount,omitempty"`
	// Timezone Timezone
	Timezone *model.TimeLocation `json:"timezone,omitempty"`
	// Currency Currency used in the transaction.
	Currency string `json:"currency,omitempty"`
	// FundsType Fund type
	FundsType enum.FundsType `json:"funds_type,omitempty"`
	// InvoiceID Invoice ID
	InvoiceID string `json:"invoice_id,omitempty"`
	// InvoiceSerialNumber Invoice serial number
	InvoiceSerialNumber string `json:"invoice_serial_number,omitempty"`
}

// TransactionSummary Sum of transactions
type TransactionSummary struct {
	// AmountCharged Total amount of money charged to the account.
	AmountCharged float64 `json:"amount_charged,omitempty"`
	// AmountPaid Total amount of money transferred out.
	AmountPaid float64 `json:"amount_paid,omitempty"`
	// Currency currency
	Currency string `json:"currency,omitempty"`
}
