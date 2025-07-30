package enum

// ExcludeFieldType 想要从返回中移除的字段种类。
type ExcludeFieldType string

const (
	// ExcludeFieldType_NULL_FIELD 值为null的字段
	ExcludeFieldType_NULL_FIELD ExcludeFieldType = "NULL_FIELD"
)
