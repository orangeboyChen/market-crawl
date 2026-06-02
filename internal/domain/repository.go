package domain

// NetValueRepository defines the interface for fetching net value data from external sources.
type NetValueRepository interface {
	GetNetValueList(req NetValueRequest) (*NetValueResponse, error)
}
