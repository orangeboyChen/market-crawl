package interfaces

import (
	"net/http"

	"market-crawl/internal/application"
)

// BocRevenueHandler handles HTTP requests for BOC revenue list data.
type BocRevenueHandler struct {
	service *application.BocRevenueService
}

// NewBocRevenueHandler creates a new BocRevenueHandler.
func NewBocRevenueHandler(service *application.BocRevenueService) *BocRevenueHandler {
	return &BocRevenueHandler{service: service}
}

// GetRevenueList handles GET /api/boc-revenue-list?strBakCode=xxx&fundCycle=xxx
func (h *BocRevenueHandler) GetRevenueList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	strBakCode := r.URL.Query().Get("strBakCode")
	if strBakCode == "" {
		http.Error(w, "strBakCode is required", http.StatusBadRequest)
		return
	}

	fundCycle := r.URL.Query().Get("fundCycle")
	if fundCycle == "" {
		http.Error(w, "fundCycle is required", http.StatusBadRequest)
		return
	}

	result, err := h.service.GetRevenueList(strBakCode, fundCycle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
