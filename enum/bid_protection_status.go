package enum

// BidProtectionStatus The current bid protection status for the campaign or ad group.
type BidProtectionStatus string

const (
	// BidProtectionStatus_ACTIVE: Bid protection is active for the campaign or ad group. You’re eligible for ad credits when costs exceed your bid. You can query historical records through /report/bid_protection/detail/get/.
	BidProtectionStatus_ACTIVE BidProtectionStatus = "ACTIVE"
	// BidProtectionStatus_INVALID: Bid protection is temporarily ineligible for the campaign or ad group because the campaign or ad group was paused or deleted within the first 3 days after creation.
	BidProtectionStatus_INVALID BidProtectionStatus = "INVALID"
	// BidProtectionStatus_INACTIVE: Bid protection is permanently ineligible for the campaign or ad group.
	BidProtectionStatus_INACTIVE BidProtectionStatus = "INACTIVE"
)
