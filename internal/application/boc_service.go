package application

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

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

// GetRevenueList fetches the revenue list from the BOC API, calculates NAV for each item
// based on the provided baseDate and baseNav, and returns the enriched JSON response.
// baseDate format: "2006-01-02", baseNav is the known NAV on that date.
func (s *BocRevenueService) GetRevenueList(strBakCode, fundCycle, baseDate string, baseNav float64) ([]byte, error) {
	req := domain.BocRevenueRequest{
		StrBakCode: strBakCode,
		FundCycle:  fundCycle,
	}

	rawResp, err := s.repo.GetRevenueList(req)
	if err != nil {
		return nil, err
	}

	var resp domain.BocRevenueResponse
	if err := json.Unmarshal(rawResp, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse BOC response: %w", err)
	}

	// Sort items by date ascending for calculation
	sort.Slice(resp.Items, func(i, j int) bool {
		return resp.Items[i].SdtPeriod < resp.Items[j].SdtPeriod
	})

	// Find the base date index
	baseIdx := -1
	for i, item := range resp.Items {
		if item.SdtPeriod == baseDate {
			baseIdx = i
			break
		}
	}

	if baseIdx == -1 {
		return nil, fmt.Errorf("baseDate %s not found in response items", baseDate)
	}

	// Set NAV for the base date
	resp.Items[baseIdx].Nav = strconv.FormatFloat(baseNav, 'f', 8, 64)

	// Calculate NAV forward from base date (baseIdx+1 to end)
	// nav[i] = nav[i-1] * (1 + tenThouRet[i] / 10000)
	for i := baseIdx + 1; i < len(resp.Items); i++ {
		prevNav, _ := strconv.ParseFloat(resp.Items[i-1].Nav, 64)
		tenThouRet, err := strconv.ParseFloat(resp.Items[i].TenThouRet, 64)
		if err != nil {
			tenThouRet = 0
		}
		nav := prevNav * (1 + tenThouRet/10000)
		resp.Items[i].Nav = strconv.FormatFloat(nav, 'f', 8, 64)
	}

	// Calculate NAV backward from base date (baseIdx-1 to 0)
	// nav[i] = nav[i+1] / (1 + tenThouRet[i+1] / 10000)
	for i := baseIdx - 1; i >= 0; i-- {
		nextNav, _ := strconv.ParseFloat(resp.Items[i+1].Nav, 64)
		tenThouRet, err := strconv.ParseFloat(resp.Items[i+1].TenThouRet, 64)
		if err != nil {
			tenThouRet = 0
		}
		nav := nextNav / (1 + tenThouRet/10000)
		resp.Items[i].Nav = strconv.FormatFloat(nav, 'f', 8, 64)
	}

	// Sort items back to descending order (latest first, matching upstream format)
	sort.Slice(resp.Items, func(i, j int) bool {
		return resp.Items[i].SdtPeriod > resp.Items[j].SdtPeriod
	})

	result, err := json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal enriched response: %w", err)
	}

	return result, nil
}
