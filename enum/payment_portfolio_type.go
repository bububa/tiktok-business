package enum

// PaymentPortfolioType The type of the Payment Portfolio.
type PaymentPortfolioType string

const (
	// PaymentPortfolioType_SHARED: Advanced Payment Portfolio. For Advanced Payment Portfolios, funds are centrally managed and jointly used by all linked ad accounts.
	PaymentPortfolioType_SHARED PaymentPortfolioType = "SHARED"
	// PaymentPortfolioType_NON_SHARED: Standard Payment Portfolio. For Standard Payment Portfolios, funds are individually managed for each linked ad account.
	PaymentPortfolioType_NON_SHARED PaymentPortfolioType = "NON_SHARED"
)
