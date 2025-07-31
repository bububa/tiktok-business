package enum

// OperationStatus 操作状态。
type OperationStatus string

const (
	// OperationStatus_ENABLE 处于启用（“开”）状态。
	OperationStatus_ENABLE OperationStatus = "ENABLE"
	// OperationStatus_DISABLE 处于未启用（“关”）状态。
	OperationStatus_DISABLE OperationStatus = "DISABLE"
	// OperationStatus_DELETE：删除状态。
	OperationStatus_DELETE OperationStatus = "DELETE"
	// OperationStatus_FROZEN：处于冻结状态，无法再次启用。
	OperationStatus_FROZEN OperationStatus = "FROZEN"
)
