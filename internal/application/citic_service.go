package application

import (
	"strings"

	"market-crawl/internal/domain"
)

// CiticProductNavService is the application service that handles CITIC product NAV business logic.
type CiticProductNavService struct {
	repo domain.CiticProductNavRepository
}

// NewCiticProductNavService creates a new CiticProductNavService with the given repository.
func NewCiticProductNavService(repo domain.CiticProductNavRepository) *CiticProductNavService {
	return &CiticProductNavService{repo: repo}
}

// GetProductNav fetches the product NAV from the CITIC Wealth API and returns raw response bytes.
// startDate and endDate are expected in "2006-01-02" format and will be converted to "20060102".
func (s *CiticProductNavService) GetProductNav(prodCode, startDate, endDate string) ([]byte, error) {
	req := domain.CiticProductNavRequest{
		ProdCode:  prodCode,
		StartDate: strings.ReplaceAll(startDate, "-", ""),
		EndDate:   strings.ReplaceAll(endDate, "-", ""),
	}
	return s.repo.GetProductNav(req)
}
