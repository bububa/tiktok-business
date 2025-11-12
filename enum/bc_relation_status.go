package enum

// BcRelationStatus 商务中心与资产的关系状态。
type BcRelationStatus string

const (
	// BcRelationStatus_BOUND: 管理资产的申请已获拥有者批准。
	BcRelationStatus_BOUND BcRelationStatus = "BOUND"
	// BcRelationStatus_UNBOUND: 管理资产的申请获拥有者批准后已被取消。
	BcRelationStatus_UNBOUND BcRelationStatus = "UNBOUND"
	// BcRelationStatus_PENDING: 管理资产的申请正在等待拥有者批准。
	BcRelationStatus_PENDING BcRelationStatus = "PENDING"
	// BcRelationStatus_REJECTED: 管理资产的申请已被拥有者拒绝。
	BcRelationStatus_REJECTED BcRelationStatus = "REJECTED"
)
