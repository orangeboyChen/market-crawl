package domain

// CxxxxProductNavRepository defines the interface for fetching product NAV data.
type CxxxxProductNavRepository interface {
	GetProductNav(req CxxxxProductNavRequest) ([]byte, error)
}
