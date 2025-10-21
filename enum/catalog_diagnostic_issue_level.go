package enum

// CatalogDiagnosticIssueLevel 用于筛选结果的问题级别。
type CatalogDiagnosticIssueLevel string

const (
	// CatalogDiagnosticIssueLevel_CRITICAL ：筛选需要立刻关注的严重问题。
	CatalogDiagnosticIssueLevel_CRITICAL CatalogDiagnosticIssueLevel = "CRITICAL"
	// CatalogDiagnosticIssueLevel_WARNING ：筛选警告型问题。这些问题有对应的建议，可用于优化商品设置。
	CatalogDiagnosticIssueLevel_WARNING CatalogDiagnosticIssueLevel = "WARNING"
	// CatalogDiagnosticIssueLevel_ERROR ：错误。此问题影响广告投放，请解决
	CatalogDiagnosticIssueLevel_ERROR CatalogDiagnosticIssueLevel = "ERROR"
	// CatalogDiagnosticIssueLevel_INFO：无问题
	CatalogDiagnosticIssueLevel_INFO CatalogDiagnosticIssueLevel = "INFO"
)
