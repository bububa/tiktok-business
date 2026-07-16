package enum

type TTOAnchorType string

const TTOAnchorType_WEB_ANCHOR TTOAnchorType = "WEB_ANCHOR"

type TTOAnchorSubType string

const (
	TTOAnchorSubType_DESTINATION   TTOAnchorSubType = "DESTINATION"
	TTOAnchorSubType_ECOMMERCE     TTOAnchorSubType = "ECOMMERCE"
	TTOAnchorSubType_ENTERTAINMENT TTOAnchorSubType = "ENTERTAINMENT"
)

type TTOAnchorStatus string

const (
	TTOAnchorStatus_DRAFT             TTOAnchorStatus = "DRAFT"
	TTOAnchorStatus_CREATED           TTOAnchorStatus = "CREATED"
	TTOAnchorStatus_IN_REVIEW         TTOAnchorStatus = "IN_REVIEW"
	TTOAnchorStatus_APPROVED          TTOAnchorStatus = "APPROVED"
	TTOAnchorStatus_REJECTED_BY_AUDIT TTOAnchorStatus = "REJECTED_BY_AUDIT"
)
