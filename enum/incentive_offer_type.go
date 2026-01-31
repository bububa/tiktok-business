package enum

// IncentiveOfferType The type of incentive offer applicable to the Upgraded Smart+ Ad Group.
// If your ad group meets the eligibility criteria and exceeds certain CPA or Minimum ROAS/Target ROAS thresholds, you will be incentivized with ad credits to use within the same ad account.
// To learn more about the incentive offer eligibility criteria and the calculation of incentive amount, see Smart+ Platform Incentive Offer (Cost Cap/Minimum ROAS/Target ROAS).
type IncentiveOfferType string

const (
	// IncentiveOfferType_INELIGIBLE: The ad group is ineligible for any incentive offer.
	IncentiveOfferType_INELIGIBLE IncentiveOfferType = "INELIGIBLE"
	// IncentiveOfferType_COST_CAP_AND_MIN_ROAS: The ad group uses the Cost Cap or Minimum ROAS/Target ROAS bidding strategy and is eligible for the incentive offer.
	IncentiveOfferType_COST_CAP_AND_MIN_ROAS IncentiveOfferType = "COST_CAP_AND_MIN_ROAS"
)
