package domain

// BocRevenueRequest represents the request parameters for fetching BOC revenue list.
type BocRevenueRequest struct {
	StrBakCode string `json:"strBakCode"`
	FundCycle  string `json:"fundCycle"`
}

// BocRevenueResponse represents the parsed response from the BOC API.
type BocRevenueResponse struct {
	ServiceResponse BocServiceResponse `json:"serviceResponse"`
	StrBakCode      string             `json:"strBakCode"`
	Items           []BocRevenueItem   `json:"items"`
}

// BocServiceResponse represents the service response metadata.
type BocServiceResponse struct {
	ResponseCode string `json:"responseCode"`
	ResponseMsg  string `json:"responseMsg"`
}

// BocRevenueItem represents a single revenue item from the BOC API.
type BocRevenueItem struct {
	SdtPeriod             string `json:"sdtPeriod"`
	TenThouRet            string `json:"tenThouRet"`
	SevenDayAnn           string `json:"sevenDayAnn"`
	FundManagerChangeFlag string `json:"fundManagerChangeFlag"`
	Nav                   string `json:"nav,omitempty"`
}
