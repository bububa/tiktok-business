package enum

// AuthorizationType 授权类型
type AuthorizationType string

const (
	// NOT_AUTHORIZED
	NOT_AUTHORIZED AuthorizationType = "NOT_AUTHORIZED"
	// AUTHORIZED
	AUTHORIZED AuthorizationType = "AUTHORIZED"
)
