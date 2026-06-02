package domain

// CiticProductNavRepository defines the interface for fetching CITIC product NAV data.
type CiticProductNavRepository interface {
	GetProductNav(req CiticProductNavRequest) ([]byte, error)
}
