package report

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// BidProtectionStatusGetRequest Get bid protection statuses API Request
type BidProtectionStatusGetRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// DataLevel The data level you want to query.
	// Enum values:
	// CAMPAIGN: Campaign level.
	// ADGROUP: Ad group level.
	DataLevel string `json:"data_level,omitempty"`
	// QueryIDs The list of IDs you want to query.
	// Max size: 200.
	// When data_level is CAMPAIGN, specify a list of Upgraded Smart+ Campaign IDs.
	// To retrieve these IDs, use /smart_plus/campaign/get/.
	// When data_level is ADGROUP, specify a list of Upgraded Smart+ Ad Group IDs.
	// To retrieve these IDs, use /smart_plus/adgroup/get/.
	// All IDs must belong to the same advertiser.
	QueryIDs []string `json:"query_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r *BidProtectionStatusGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("data_level", r.DataLevel)
	values.Set("query_ids", string(util.JSONMarshal(r.QueryIDs)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// BidProtectionStatusGetResponse Get bid protection statuses API Response
type BidProtectionStatusGetResponse struct {
	model.BaseResponse
	Data struct {
		// List The bid protection status information for your queried campaigns or ad groups.
		List []BidProtectionStatus `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// BidProtectionStatus The bid protection status information for your queried campaigns or ad groups.
type BidProtectionStatus struct {
	// DataLevel The data level.
	// Enum values:
	// CAMPAIGN: Campaign level.
	// ADGROUP: Ad group level.
	DataLevel string `json:"data_level,omitempty"`
	// QueryID The specific ID you queried.
	// When data_level is CAMPAIGN, this field represents an Upgraded Smart+ Campaign ID.
	// When data_level is ADGROUP, this field represents an Upgraded Smart+ Ad Group ID.
	QueryID string `json:"query_id,omitempty"`
	// BidProtectionStatus The current bid protection status for the campaign or ad group.
	// Enum values:
	// ACTIVE: Bid protection is active for the campaign or ad group. You’re eligible for ad credits when costs exceed your bid. You can query historical records through /report/bid_protection/detail/get/.
	// INVALID: Bid protection is temporarily ineligible for the campaign or ad group because the campaign or ad group was paused or deleted within the first 3 days after creation.
	// INACTIVE: Bid protection is permanently ineligible for the campaign or ad group.
	BidProtectionStatus enum.BidProtectionStatus `json:"bid_protection_status,omitempty"`
	// CompensationCategory he bid protection compensation category.
	// Enum values:
	// FULL_LIFE_CYCLE: Full lifecycle compensation.Bid Protection (Full Lifecycle) is the default solution, covering the campaign's entire duration.
	// THREE_DAY: Three-day compensation. Only bidding strategies that are currently unsupported in the Full Lifecycle version will use Bid Protection (Learning Phase). This covers the first three full consecutive days from the day you activate the campaign.
	CompensationCategory string `json:"compensation_category,omitempty"`
}
