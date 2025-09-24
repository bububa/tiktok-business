package enum

// PixelEventRuleOperator 操作符
type PixelEventRuleOperator string

const (
	// PixelEventRuleOperator_OPERATORTYPE_CONTAINS
	PixelEventRuleOperator_OPERATORTYPE_CONTAINS PixelEventRuleOperator = "OPERATORTYPE_CONTAINS"
	// PixelEventRuleOperator_OPERATORTYPE_DOES_NOT_EQUAL
	PixelEventRuleOperator_OPERATORTYPE_DOES_NOT_EQUAL PixelEventRuleOperator = "OPERATORTYPE_DOES_NOT_EQUAL"
	// PixelEventRuleOperator_OPERATORTYPE_EQUALS
	PixelEventRuleOperator_OPERATORTYPE_EQUALS PixelEventRuleOperator = "OPERATORTYPE_EQUALS"
)
