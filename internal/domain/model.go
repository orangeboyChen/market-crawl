package domain

// NetValueRequest represents the request parameters for fetching net value list.
type NetValueRequest struct {
	ProdId    string `json:"prodId"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

// NetValueResponse represents the raw response from the ICBC API.
type NetValueResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
