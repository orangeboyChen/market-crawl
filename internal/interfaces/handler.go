package interfaces

import (
	"net/http"
	"strconv"

	"market-crawl/internal/application"
)

// NetValueHandler handles HTTP requests for net value data.
type NetValueHandler struct {
	service *application.NetValueService
}

// NewNetValueHandler creates a new NetValueHandler.
func NewNetValueHandler(service *application.NetValueService) *NetValueHandler {
	return &NetValueHandler{service: service}
}

// GetNetValueList handles GET /api/net-value-list?prodId=xxx&pageIndex=1&pageSize=10
func (h *NetValueHandler) GetNetValueList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	prodId := r.URL.Query().Get("prodId")
	if prodId == "" {
		http.Error(w, "prodId is required", http.StatusBadRequest)
		return
	}

	pageIndex, err := strconv.Atoi(r.URL.Query().Get("pageIndex"))
	if err != nil {
		pageIndex = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		pageSize = 10
	}

	result, err := h.service.GetNetValueList(prodId, pageIndex, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
