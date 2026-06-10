package domain

// CxxxxProductNavRequest represents the request parameters for fetching product NAV.
type CxxxxProductNavRequest struct {
	ProdCode  string `json:"prodCode"`
	StartDate string `json:"startDate"` // Format: yyyyMMdd
	EndDate   string `json:"endDate"`   // Format: yyyyMMdd
}

// CxxxxProductNavResponse represents the raw response from the placeholder API.
type CxxxxProductNavResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
