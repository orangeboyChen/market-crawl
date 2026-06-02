package application

import (
	"market-crawl/internal/domain"
)

// NetValueService is the application service that handles net value business logic.
type NetValueService struct {
	repo domain.NetValueRepository
}

// NewNetValueService creates a new NetValueService with the given repository.
func NewNetValueService(repo domain.NetValueRepository) *NetValueService {
	return &NetValueService{repo: repo}
}

// GetNetValueList fetches the net value list from the external API.
func (s *NetValueService) GetNetValueList(prodId string, pageIndex, pageSize int) (*domain.NetValueResponse, error) {
	req := domain.NetValueRequest{
		ProdId:    prodId,
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}
	return s.repo.GetNetValueList(req)
}
