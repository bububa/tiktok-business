package enum

// PixelAssetRelationStatus 资产关系状态。
type PixelAssetRelationStatus string

const (
	// PixelAssetRelationStatus_NULL: 该 Pixel 为当前广告主（advertiser_id）所有，且尚未作为资产转移至商务中心中。
	PixelAssetRelationStatus_NULL PixelAssetRelationStatus = "null"
	// PixelAssetRelationStatus_TRANSFERRED: 该 Pixel 已作为资产转移至商务中心中，但仍为当前广告主（advertiser_id）所有。若想转移 Pixel，可使用/bc/pixel/transfer/。
	PixelAssetRelationStatus_TRANSFERRED PixelAssetRelationStatus = "TRANSFERRED"
	// PixelAssetRelationStatus_SHARED: 该 Pixel 已作为资产分享给当前广告主（advertiser_id）。对于由他人分享的资产 Pixel ，你可以修改，但是不能删除。
	PixelAssetRelationStatus_SHARED PixelAssetRelationStatus = "SHARED"
	// PixelAssetRelationStatus_UNBOUND: 该 Pixel 之前作为资产分享给当前广告主（advertiser_id），但已与当前广告主解绑。在统计数据中，状态为UNBOUND 的资产不会被计入。
	PixelAssetRelationStatus_UNBOUND PixelAssetRelationStatus = "UNBOUND"
)
