package enum

	// ProductUpdateMode 更新模式。
type ProductUpdateMode string

const (
// ProductUpdateMode_INCREASE：增量。
ProductUpdateMode_INCREASE ProductUpdateMode = "INCREASE"
// ProductUpdateMode_REPLACE：全量。
ProductUpdateMode_REPLACE ProductUpdateMode = "REPLACE"
)
