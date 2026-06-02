package domain

// BocRevenueRepository defines the interface for fetching BOC revenue list data.
type BocRevenueRepository interface {
	GetRevenueList(req BocRevenueRequest) ([]byte, error)
}
