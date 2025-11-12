package enum

// ShareType 分享类型
type ShareType string

const (
	// SHARED: 合作伙伴分享给您的资产。
	SHARED ShareType = "SHARED"
	// SHARING: 您分享给合作伙伴的资产。
	SHARING ShareType = "SHARING"
)
