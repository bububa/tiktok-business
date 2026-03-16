package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AccountCostGetRequest Get the cost records of a BC and ad accounts API Request
type AccountCostGetRequest struct {
	// BcID Business Center ID
	BcID string `json:"bc_id,omitempty"`
	// Filtering Filtering conditions
	Filtering *AccountCostGetFilter `json:"filtering,omitempty"`
	// Page Current number of pages. Default value: 1. Value range : ≥ 1.
	Page int `json:"page,omitempty"`
	// PageSize Page size. Default value: 10. Value range: 1-50.
	PageSize int `json:"page_size,omitempty"`
}

type AccountCostGetFilter struct {
	// Keyword Search keyword.
	// The fuzzy match is supported.
	// You can specify an advertiser ID or name as search keyword.
	Keyword string `json:"keyword,omitempty"`
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
func (r *AccountCostGetRequest) Encode() string {
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

// AccountCostGetResponse Get the cost records of a BC and ad accounts API Response
type AccountCostGetResponse struct {
	model.BaseResponse
	Data *AccountCostGetResult `json:"data,omitempty"`
}

type AccountCostGetResult struct {
	// PageInfo Pagination information
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// CostList The list of cost records for each advertiser
	CostList []AccountCost `json:"cost_list,omitempty"`
	// TransactionSummary A summary of cost records for the Business Center.
	TransactionSummary *AccountCostSummary `json:"transaction_summary,omitempty"`
}

type AccountCost struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserName Advertiser name
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// Amount Total cost amount for the advertiser
	Amount float64 `json:"amount,omitempty"`
	// Cash cost amount for the advertiser.
	CashAmount float64 `json:"cash_amount,omitempty"`
	// GrantAmount Ad credit cost amount for the advertiser.
	GrantAmount float64 `json:"grant_amount,omitempty"`
	// TaxAmount Estimated tax amount for the advertiser
	TaxAmount float64 `json:"tax_amount,omitempty"`
	// Currency
	Currency string `json:"currency,omitempty"`
}

type AccountCostSummary struct {
	// Amount Total cost amount for the Business Center
	Amount float64 `json:"amount,omitempty"`
	// Cash cost amount for the Business Center.
	CashAmount float64 `json:"cash_amount,omitempty"`
	// GrantAmount Ad credit cost amount for the Business Center.
	GrantAmount float64 `json:"grant_amount,omitempty"`
	// TaxAmount Estimated tax amount for the Business Center
	TaxAmount float64 `json:"tax_amount,omitempty"`
	// Currency
	Currency string `json:"currency,omitempty"`
}
