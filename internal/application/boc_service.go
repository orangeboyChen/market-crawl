package application

import (
	"market-crawl/internal/domain"
)

// BocRevenueService is the application service that handles BOC revenue list business logic.
type BocRevenueService struct {
	repo domain.BocRevenueRepository
}

// NewBocRevenueService creates a new BocRevenueService with the given repository.
func NewBocRevenueService(repo domain.BocRevenueRepository) *BocRevenueService {
	return &BocRevenueService{repo: repo}
}

// GetRevenueList fetches the revenue list from the BOC API and returns raw response bytes.
func (s *BocRevenueService) GetRevenueList(strBakCode, fundCycle string) ([]byte, error) {
	req := domain.BocRevenueRequest{
		StrBakCode: strBakCode,
		FundCycle:  fundCycle,
	}
	return s.repo.GetRevenueList(req)
}
