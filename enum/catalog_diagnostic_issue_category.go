package enum

// CatalogDiagnosticIssueCategory 用于筛选结果的问题类型。
type CatalogDiagnosticIssueCategory string

const (
	// CatalogDiagnosticIssueCategory_PRODUCT_ATTRIBUTES ：商品属性问题。
	CatalogDiagnosticIssueCategory_PRODUCT_ATTRIBUTES CatalogDiagnosticIssueCategory = "PRODUCT_ATTRIBUTES"
	// CatalogDiagnosticIssueCategory_PRODUCT_REVIEW ：商品审核问题。
	CatalogDiagnosticIssueCategory_PRODUCT_REVIEW CatalogDiagnosticIssueCategory = "PRODUCT_REVIEW"
	// CatalogDiagnosticIssueCategory_CATALOG ：商品库问题。
	CatalogDiagnosticIssueCategory_CATALOG CatalogDiagnosticIssueCategory = "CATALOG"
	// CatalogDiagnosticIssueCategory_PIXEL_OR_EVENT ：Pixel 或事件问题。
	CatalogDiagnosticIssueCategory_PIXEL_OR_EVENT CatalogDiagnosticIssueCategory = "PIXEL_OR_EVENT"
	// CatalogDiagnosticIssueCategory_FILE_UPLOAD_OR_FEED ：文件上传或更新源问题。
	CatalogDiagnosticIssueCategory_FILE_UPLOAD_OR_FEED CatalogDiagnosticIssueCategory = "FILE_UPLOAD_OR_FEED"
)
