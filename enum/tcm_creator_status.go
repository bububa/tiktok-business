package enum

	// TCMCreatorStatus 创作者入驻状态
type TCMCreatorStatus string

const (
	// TCMCreatorStatus_IN
	TCMCreatorStatus_IN TCMCreatorStatus = "IN"
	// TCMCreatorStatus_NOT_IN
	TCMCreatorStatus_NOT_IN TCMCreatorStatus = "NOT_IN"
	// TCMCreatorStatus_INVALID
	TCMCreatorStatus_INVALID TCMCreatorStatus = "INVALID"
	// TCMCreatorStatus_INVITED
	TCMCreatorStatus_INVITED TCMCreatorStatus = "INVITED"
)
