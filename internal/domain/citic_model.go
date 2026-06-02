package domain

// CiticProductNavRequest represents the request parameters for fetching CITIC product NAV.
type CiticProductNavRequest struct {
	ProdCode  string `json:"prodCode"`
	StartDate string `json:"startDate"` // Format: yyyyMMdd
	EndDate   string `json:"endDate"`   // Format: yyyyMMdd
}

// CiticProductNavResponse represents the raw response from the CITIC Wealth API.
type CiticProductNavResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
