package enum

// OperationStatus 操作状态。
type OperationStatus string

const (
	// OperationStatus_ENABLE：推广系列处于启用（“开”）状态。
	OperationStatus_ENABLE OperationStatus = "ENABLE"
	// OperationStatus_DISABLE：推广系列处于未启用（“关”）状态。
	OperationStatus_DISABLE OperationStatus = "DISABLE"
)
