package interfaces

import (
	"net/http"

	"market-crawl/internal/application"
)

// CxxxxProductNavHandler handles HTTP requests for product NAV data.
type CxxxxProductNavHandler struct {
	service *application.CxxxxProductNavService
}

// NewCxxxxProductNavHandler creates a new CxxxxProductNavHandler.
func NewCxxxxProductNavHandler(service *application.CxxxxProductNavService) *CxxxxProductNavHandler {
	return &CxxxxProductNavHandler{service: service}
}

// GetProductNav handles GET /api/citic-product-nav?prodCode=xxx&startDate=2026-01-01&endDate=2026-06-01
func (h *CxxxxProductNavHandler) GetProductNav(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	prodCode := r.URL.Query().Get("prodCode")
	if prodCode == "" {
		http.Error(w, "prodCode is required", http.StatusBadRequest)
		return
	}

	startDate := r.URL.Query().Get("startDate")
	if startDate == "" {
		http.Error(w, "startDate is required", http.StatusBadRequest)
		return
	}

	endDate := r.URL.Query().Get("endDate")
	if endDate == "" {
		http.Error(w, "endDate is required", http.StatusBadRequest)
		return
	}

	result, err := h.service.GetProductNav(prodCode, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
