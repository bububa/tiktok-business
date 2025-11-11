package enum

// BcAssetRelationStatus 商务中心与资产的关系状态。
type BcAssetRelationStatus string

const (
	// BcAssetRelationStatus_BOUND: 管理资产的申请已获拥有者批准。
	BcAssetRelationStatus_BOUND BcAssetRelationStatus = "BOUND"
	// BcAssetRelationStatus_UNBOUND: 管理资产的申请获拥有者批准后已被取消。
	BcAssetRelationStatus_UNBOUND BcAssetRelationStatus = "UNBOUND"
	// BcAssetRelationStatus_PENDING: 管理资产的申请正在等待拥有者批准。
	BcAssetRelationStatus_PENDING BcAssetRelationStatus = "PENDING"
	// BcAssetRelationStatus_REJECTED: 管理资产的申请已被拥有者拒绝。
	BcAssetRelationStatus_REJECTED BcAssetRelationStatus = "REJECTED"
)
