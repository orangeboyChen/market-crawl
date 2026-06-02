package domain

// BocRevenueRequest represents the request parameters for fetching BOC revenue list.
type BocRevenueRequest struct {
	StrBakCode string `json:"strBakCode"`
	FundCycle  string `json:"fundCycle"`
}
