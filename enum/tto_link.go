package enum

type TTOLinkAction string

const (
	TTOLinkAction_LINK   TTOLinkAction = "LINK"
	TTOLinkAction_REVOKE TTOLinkAction = "REVOKE"
)

type TTOLinkConfirmAction string

const (
	TTOLinkConfirmAction_APPROVE TTOLinkConfirmAction = "APPROVE"
	TTOLinkConfirmAction_REJECT  TTOLinkConfirmAction = "REJECT"
)

type TTOLinkStatus string

const (
	TTOLinkStatus_IN_REVIEW TTOLinkStatus = "IN_REVIEW"
	TTOLinkStatus_APPROVE   TTOLinkStatus = "APPROVE"
	TTOLinkStatus_REJECTED  TTOLinkStatus = "REJECTED"
	TTOLinkStatus_REVOKE    TTOLinkStatus = "REVOKE"
)
