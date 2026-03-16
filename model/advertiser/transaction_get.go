package advertiser

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TransactionGetRequest Get the transaction records of ad accounts API Request
type TransactionGetRequest struct {
	// BcID Business Center ID.
	// Note: You can only get the transaction records of ad accounts under a Business Center whose type is not NORMAL. You can use /bc/get/ to get the Business Center type via type in the response.
	BcID string `json:"bc_id,omitempty"`
	// Filtering Filtering conditions
	Filtering *TransactionGetFilter `json:"filtering,omitempty"`
	// Page Current number of pages. Default value: 1. Value range : ≥ 1.
	Page int `json:"page,omitempty"`
	// PageSize Page size. Default value: 10. Value range: 1-50.
	PageSize int `json:"page_size,omitempty"`
}

type TransactionGetFilter struct {
	// TransferType Billing type. Enum values: TRANS_TYPE_TRANSFER(transfer), TRANS_TYPE_TAX(tax), TRANS_TYPE_COST(consumption), TRANS_TYPE_PAYMENT (payment). Default value: TRANS_TYPE_TRANSFER.
	// Note: funds_type returned in response will be null when you set transfer_type as TRANS_TYPE_TAX or TRANS_TYPE_COST, because these billing types don't involve funds type settings.
	TransferType enum.AdvertiserTransferType `json:"transfer_type,omitempty"`
	// FundsType Fund type. Allowed types: FUNDS_TYPE_CASH(cash), FUNDS_TYPE_GRANT(coupon/voucher). The default is to select both types.
	FundsType enum.FundsType `json:"funds_type,omitempty"`
	// SummaryByAccount Whether to summarize by account. false means no summary; true means summary. Default value: false.
	SummaryByAccount bool `json:"summary_by_account,omitempty"`
	// Keyword Search keywords, you can search for ad account name or ad account ID.
	Keyword string `json:"keyword,omitempty"`
	// StartDate Transaction record search start time, (UTC+0) format：2020-10-12. The default date is 90 days ago.
	StartDate string `json:"start_date,omitempty"`
	// EndDate Transaction record search end time, (UTC+0) format：2020-11-12. The default date is the same day.
	// Note: The largest time gap allowed is one year, so please specify a date within one year after the start_date as end_date.
	EndDate string `json:"end_date,omitempty"`
}

// Encode implements GetRequest
func (r *TransactionGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
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

// TransactionGetResponse Get the transaction records of ad accounts API Response
type TransactionGetResponse struct {
	model.BaseResponse
	Data *TransactionGetResult `json:"data,omitempty"`
}

type TransactionGetResult struct {
	// PageInfo Pagination information
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// TransactionList Ad Account transaction record.
	TransactionList []Transaction `json:"transaction_list,omitempty"`
}

// Transaction Ad Account transaction record.
type Transaction struct {
	// AdvertiserID Ad Account ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserName Ad Account Name
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// Amount Transaction amount, kept to two decimal places.
	Amount float64 `json:"amount,omitempty"`
	// Currency Transaction currency, the value range is shown in Appendix-Currency
	Currency string `json:"currency,omitempty"`
	// Date Trading time, (UTC+0) format：2020-10-12 00:00:00
	Date model.DateTime `json:"date,omitzero"`
	// FundsType Type of funding. Enum values：FUNDS_TYPE_CASH(cash), FUNDS_TYPE_GRANT(coupon/voucher) .
	FundsType enum.FundsType `json:"funds_type,omitempty"`
	// TransferType Billing type. Enum values: TRANS_TYPE_TRANSFER(transfer), TRANS_TYPE_TAX(consumption), TRANS_TYPE_COST(tax).
	TransferType enum.AdvertiserTransferType `json:"transfer_type,omitempty"`
	// Timezone Trading time zone. For enum values, see Appendix-Time Zone
	Timezone *model.TimeLocation `json:"timezone,omitempty"`
}
