package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"market-crawl/internal/domain"
)

const bocRevenueURL = "https://ebsnew.boc.cn/SAP/bocop/unlogin/ezdb/app/ten_thou_tevenue_list_info"

// BocClient implements domain.BocRevenueRepository by calling the BOC API.
type BocClient struct {
	httpClient *http.Client
}

// NewBocClient creates a new BocClient with a default timeout.
func NewBocClient() *BocClient {
	return &BocClient{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetRevenueList sends a POST request to the BOC API and returns the raw response bytes.
func (c *BocClient) GetRevenueList(req domain.BocRevenueRequest) ([]byte, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, bocRevenueURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("clentid", "553")
	httpReq.Header.Set("chnflg", "6")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBody, nil
}
