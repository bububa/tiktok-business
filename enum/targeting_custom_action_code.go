package enum

// TargetingCustomActionCode The custom action used to filter out the audiences.
type TargetingCustomActionCode string

const (
	// TargetingCustomActionCode_VIEW_PRODUCT: The audience viewed the product.
	TargetingCustomActionCode_VIEW_PRODUCT TargetingCustomActionCode = "VIEW_PRODUCT"
	// TargetingCustomActionCode_ADD_TO_CART: The audience added the product to the cart.
	TargetingCustomActionCode_ADD_TO_CART TargetingCustomActionCode = "ADD_TO_CART"
	// TargetingCustomActionCode_PURCHASE: The audience purchased the product.
	TargetingCustomActionCode_PURCHASE TargetingCustomActionCode = "PURCHASE"
)
