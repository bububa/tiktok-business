package enum

// BusinessType 页面类型。
type BusinessType string

const (
	// BusinessType_LEAD_GEN(即时表单（线索表单）)
	BusinessType_LEAD_GEN BusinessType = "LEAD_GEN"
	// BusinessType_STORE_FRONT(商品橱窗页)
	BusinessType_STORE_FRONT BusinessType = "STORE_FRONT"
	// BusinessType_APP_PROFILE_PAGE (应用介绍页)
	BusinessType_APP_PROFILE_PAGE BusinessType = "APP_PROFILE_PAGE"
	// BusinessType_TIKTOK_INSTANT_PAGE ( 自定义页面 (自定义TikTok即时体验页面 ) )
	BusinessType_TIKTOK_INSTANT_PAGE BusinessType = "TIKTOK_INSTANT_PAGE"
	// BusinessType_SHOP_ADS_PLP (店铺广告产品列表页)
	BusinessType_SHOP_ADS_PLP BusinessType = "SHOP_ADS_PLP"
	// BusinessType_SHOP_ADS_PDP (店铺广告产品详情页)
	BusinessType_SHOP_ADS_PDP BusinessType = "SHOP_ADS_PDP"
	// BusinessType_POP_UP_FORM(私信页面) 。
	BusinessType_POP_UP_FORM BusinessType = "POP_UP_FORM"
)
