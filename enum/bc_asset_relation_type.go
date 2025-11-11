package enum

// BcAssetRelationType 商务中心与资产的关系类型。
type BcAssetRelationType string

const (
	// BcAssetRelationType_OWNER_BC: 该资产由当前商务中心创建并拥有。
	BcAssetRelationType_OWNER_BC BcAssetRelationType = "OWNER_BC"
	// BcAssetRelationType_OWNER_PARTNER: 该资产由合作伙伴商务中心拥有。
	BcAssetRelationType_OWNER_PARTNER BcAssetRelationType = "OWNER_PARTNER"
	// BcAssetRelationType_OWNER_INDIVIDUAL: 该资产由一个个人账户拥有。
	BcAssetRelationType_OWNER_INDIVIDUAL BcAssetRelationType = "OWNER_INDIVIDUAL"
)
