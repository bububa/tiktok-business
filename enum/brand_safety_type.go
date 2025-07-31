package enum

// BrandSafetyType 品牌安全类型。枚举值：
type BrandSafetyType string

const (
	// BrandSafetyType_NO_BRAND_SAFETY：不使用任何品牌安全解决方案，即使用全部库存。您的广告可能会在某些包含成人主题的内容周围展示。
	BrandSafetyType_NO_BRAND_SAFETY BrandSafetyType = "NO_BRAND_SAFETY"
	// BrandSafetyType_EXPANDED_INVENTORY：使用TikTok的品牌安全解决方案，即使用扩展库存。您的广告将不会出现在不当内容或包含成人主题的内容后面。在下个API版本中EXPANDED_INVENTORY将替代NO_BRAND_SAFETY，并成为默认的品牌安全选项。
	BrandSafetyType_EXPANDED_INVENTORY BrandSafetyType = "EXPANDED_INVENTORY"
	// BrandSafetyType_STANDARD_INVENTORY：使用TikTok的品牌安全解决方案中的标准库存。您的广告将展示在适合大多数品牌的内容周围。
	BrandSafetyType_STANDARD_INVENTORY BrandSafetyType = "STANDARD_INVENTORY"
	// BrandSafetyType_LIMITED_INVENTORY：使用TikTok的品牌安全解决方案中的有限库存。您的广告将展示在不包含成人主题的内容周围。
	BrandSafetyType_LIMITED_INVENTORY BrandSafetyType = "LIMITED_INVENTORY"
	// BrandSafetyType_THIRD_PARTY：使用第三方品牌安全解决方案。
	BrandSafetyType_THIRD_PARTY BrandSafetyType = "THIRD_PARTY"
)
