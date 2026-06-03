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

	// Filter out items where TenThouRet or SevenDayAnn is empty
	filtered := make([]domain.BocRevenueItem, 0, len(resp.Items))
	for _, item := range resp.Items {
		if item.TenThouRet == "" || item.SevenDayAnn == "" {
			continue
		}
		filtered = append(filtered, item)
	}
	resp.Items = filtered

	// If no valid items remain, return early without items
	if len(resp.Items) == 0 {
		result, err := json.Marshal(resp)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return result, nil
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
		// Base date not in items; find the nearest insertion point and use baseNav as a virtual anchor
		insertIdx := sort.Search(len(resp.Items), func(i int) bool {
			return resp.Items[i].SdtPeriod >= baseDate
		})
		// Calculate NAV forward from insertIdx using baseNav as the virtual previous NAV
		prevNav := baseNav
		for i := insertIdx; i < len(resp.Items); i++ {
			tenThouRet, err := strconv.ParseFloat(resp.Items[i].TenThouRet, 64)
			if err != nil {
				tenThouRet = 0
			}
			nav := prevNav * (1 + tenThouRet/10000)
			resp.Items[i].Nav = strconv.FormatFloat(nav, 'f', 8, 64)
			prevNav = nav
		}
		// Calculate NAV backward from insertIdx-1 using baseNav as the virtual next NAV
		nextNav := baseNav
		for i := insertIdx - 1; i >= 0; i-- {
			tenThouRet, err := strconv.ParseFloat(resp.Items[i+1].TenThouRet, 64)
			if err != nil {
				tenThouRet = 0
			}
			nav := nextNav / (1 + tenThouRet/10000)
			resp.Items[i].Nav = strconv.FormatFloat(nav, 'f', 8, 64)
			nextNav = nav
		}
	} else {
		// Base date exists in items; set its NAV and propagate
		resp.Items[baseIdx].Nav = strconv.FormatFloat(baseNav, 'f', 8, 64)

		// Calculate NAV forward from base date (baseIdx+1 to end)
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
		for i := baseIdx - 1; i >= 0; i-- {
			nextNav, _ := strconv.ParseFloat(resp.Items[i+1].Nav, 64)
			tenThouRet, err := strconv.ParseFloat(resp.Items[i+1].TenThouRet, 64)
			if err != nil {
				tenThouRet = 0
			}
			nav := nextNav / (1 + tenThouRet/10000)
			resp.Items[i].Nav = strconv.FormatFloat(nav, 'f', 8, 64)
		}
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
