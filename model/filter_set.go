package model

import "github.com/bububa/tiktok-business/enum"

// FilterSet 筛选集合
type FilterSet struct {
	// Operator 筛选条件间的操作符，AND 代表同时满足筛选关系， OR 代表任意一项满足，目前仅支持 OR
	Operator enum.FilterSetOperator `json:"operator,omitempty"`
	// Filters 筛选条件列表。目前仅支持传入单个值，格式见下方的“筛选格式”
	Filters []Filter `json:"filters,omitempty"`
}

// Filter 筛选条件
type Filter struct {
	// Field 筛选字段。目前仅支持EVENT
	Field string `json:"field,omitempty"`
	// Operator 连接筛选字段和筛选字段值的筛选操作符。枚举值: EQ（等于），GT（大于），LT（小于）。目前仅支持EQ
	Operatior enum.FilterOperator `json:"operator,omitempty"`
	// Value 筛选字段值。获得枚举值，可参阅枚举值-筛选字段值
	Value enum.FilterValue `json:"value,omitempty"`
}
