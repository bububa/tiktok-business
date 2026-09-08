package report

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// BidProtectionDetailGetRequest Get bid protection history API Request
type BidProtectionDetailGetRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// DataLevel The data level
	// Enum values:
	// CAMPAIGN: Campaign level.
	// ADGROUP: Ad group level.
	DataLevel string `json:"data_level,omitempty"`
	// QueryIDs The list of IDs to query.
	// The maximum allowed size is 200 divided by the time range you define using start_date and end_date. For example, if you set start_date to 2026-08-01 and end_date to 2026-08-11 (a 10-day range), the maximum number of IDs you can specify in query_ids is 20 (200 divided by 10).
	// When data_level is CAMPAIGN, specify a list of Upgraded Smart+ Campaign IDs.
	// To retrieve Upgraded Smart+ Campaign IDs, use /smart_plus/campaign/get/.
	// When data_level is ADGROUP, specify a list of Upgraded Smart+ Ad Group IDs.
	// To retrieve Upgraded Smart+ Ad Group IDs, use /smart_plus/adgroup/get/.
	// All IDs must belong to the same advertiser.
	QueryIDs []string `json:"query_ids,omitempty"`
	// StartDate The start date for your bid protection history query, in the format of YYYY-MM-DD (ad account timezone).
	// Ensure the start date is within the past 60 days and is earlier than or equal to your end_date.
	StartDate string `json:"start_date,omitempty"`
	// EndDate The end date for your bid protection history query, in the format of YYYY-MM-DD (ad account timezone).
	// Ensure the end date is within 60 days after your start_date and is later than or equal to the start_date.
	EndDate string `json:"end_date,omitempty"`
}

// Encode implements GetRequest interface
func (r *BidProtectionDetailGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("data_level", r.DataLevel)
	values.Set("query_ids", string(util.JSONMarshal(r.QueryIDs)))
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// BidProtectionDetailGetResponse Get bid protection history API Response
type BidProtectionDetailGetResponse struct {
	model.BaseResponse
	Data struct {
		// BidProtectionRecords The details of the bid protection records you queried.
		BidProtectionRecords []BidProtectionRecord `json:"bid_protection_records,omitempty"`
	} `json:"data,omitempty"`
}

// BidProtectionRecord The details of the bid protection records you queried.
type BidProtectionRecord struct {
	// DataLevel The data level.
	// Enum values:
	// CAMPAIGN: Campaign level.
	// ADGROUP: Ad group level.
	DataLevel string `json:"data_level,omitempty"`
	// QueryID The specific ID you queried.
	// When data_level is CAMPAIGN, this field represents an Upgraded Smart+ Campaign ID.
	// When data_level is ADGROUP, this field represents an Upgraded Smart+ Ad Group ID.
	QueryID string `json:"query_id,omitempty"`
	// BidProtectionDailyStatus The bid protection status on the date.
	// Enum values:
	// UNDER_PROTECTION: Active. The campaign is active today; results are calculated on a daily basis.
	// INELIGIBLE: Ineligible. Bid protection is ineligible for this date because this campaign or ad group status was paused or deleted.
	// CONFIRMING: Pending. Bid protection results are pending and most ad credits will be updated in 8 to 10 days.
	// PAYMENT_COMPLETE: Credit issued. Ad credits have been issued.
	// TARGET_MET: Target met. The bid target was met, so no ad credits were issued.
	BidProtectionDailyStatus enum.BidProtectionDailyStatus `json:"bid_protection_daily_status,omitempty"`
	// StatusDetail The specific details of the bid protection status on the date.
	StatusDetail string `json:"status_detail,omitempty"`
	// CreditAmount The ad credit amount for the date, scaled by a factor of 100,000.
	// The value does not represent the exact monetary amount directly. To determine the actual amount, divide this value by 100,000.
	// For example, if you receive a credit_amount of 3000000 in USD, you can calculate the true value by dividing it by 100,000, which gives you an actual ad credit of 30 USD.
	CreditAmount model.Float64 `json:"credit_amount,omitempty"`
	// Currency The currency for the ad credits amount, in the format of ISO 4217 currency code.
	// Example: USD.
	Currency string `json:"currency,omitempty"`
}
