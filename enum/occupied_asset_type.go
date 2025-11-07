package enum

// OccupiedAssetType 想要检查是否被处于启用状态的视频购物广告，商品购物广告或直播购物广告占用的资产类型。
type OccupiedAssetType string

const (
	// OccupiedAssetType_IDENTITY_TT_USER: TikTok 用户（TT_USER）认证身份。
	OccupiedAssetType_IDENTITY_TT_USER OccupiedAssetType = "IDENTITY_TT_USER"
	// OccupiedAssetType_IDENTITY_BC_AUTH_TT: 添加到商务中心的 TikTok 用户（BC_AUTH_TT）认证身份。
	OccupiedAssetType_IDENTITY_BC_AUTH_TT OccupiedAssetType = "IDENTITY_BC_AUTH_TT"
	// OccupiedAssetType_IDENTITY_TTS_TT: TikTok Shop 关联的 TikTok 用户（TTS_TT）认证身份。
	OccupiedAssetType_IDENTITY_TTS_TT OccupiedAssetType = "IDENTITY_TTS_TT"
	// OccupiedAssetType_SPU：SPU
	OccupiedAssetType_SPU OccupiedAssetType = "SPU"
)
