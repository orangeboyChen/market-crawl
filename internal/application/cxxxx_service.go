package application

import (
	"strings"

	"market-crawl/internal/domain"
)

// CxxxxProductNavService is the application service that handles product NAV business logic.
type CxxxxProductNavService struct {
	repo domain.CxxxxProductNavRepository
}

// NewCxxxxProductNavService creates a new CxxxxProductNavService with the given repository.
func NewCxxxxProductNavService(repo domain.CxxxxProductNavRepository) *CxxxxProductNavService {
	return &CxxxxProductNavService{repo: repo}
}

// GetProductNav fetches the product NAV from the configured placeholder API and returns raw response bytes.
// startDate and endDate are expected in "2006-01-02" format and will be converted to "20060102".
func (s *CxxxxProductNavService) GetProductNav(prodCode, startDate, endDate string) ([]byte, error) {
	req := domain.CxxxxProductNavRequest{
		ProdCode:  prodCode,
		StartDate: strings.ReplaceAll(startDate, "-", ""),
		EndDate:   strings.ReplaceAll(endDate, "-", ""),
	}
	return s.repo.GetProductNav(req)
}
