package enum

// OutputFormat 输出格式。
type OutputFormat string

const (
	// CSV_STRING
	CSV_STRING OutputFormat = "CSV_STRING"
	// CSV_DOWNLOAD
	CSV_DOWNLOAD OutputFormat = "CSV_DOWNLOAD"
	// XLSX_DOWNLOAD
	XLSX_DOWNLOAD OutputFormat = "XLSX_DOWNLOAD"
)
