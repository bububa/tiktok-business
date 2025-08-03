package enum

// PageTemplateType 落地页模板类型
type PageTemplateType string

const (
	// PageTemplateType_CUSTOMIZED（定制）
	PageTemplateType_CUSTOMIZED PageTemplateType = "CUSTOMIZED"
	// PageTemplateType_INTRODUCTION（介绍）
	PageTemplateType_INTRODUCTION PageTemplateType = "INTRODUCTION"
	// PageTemplateType_BRAND_SAFETY（品牌安全）
	PageTemplateType_BRAND_SAFETY PageTemplateType = "BRAND_SAFETY"
	// PageTemplateType_SALES_PRODUCT（产品销售）
	PageTemplateType_SALES_PRODUCT PageTemplateType = "SALES_PRODUCT"
	// PageTemplateType_MOVIE_TRAILER（电影预告）
	PageTemplateType_MOVIE_TRAILER PageTemplateType = "MOVIE_TRAILER"
)
