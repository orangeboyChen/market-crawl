package infrastructure

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"market-crawl/internal/domain"
)

const cxxxxProductNavURL = "https://example.com/"

// CxxxxClient implements domain.CxxxxProductNavRepository by calling a placeholder API.
type CxxxxClient struct {
	httpClient *http.Client
}

// NewCxxxxClient creates a new CxxxxClient with a default timeout.
func NewCxxxxClient() *CxxxxClient {
	return &CxxxxClient{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetProductNav sends a GET request to the placeholder API and returns the raw response bytes.
func (c *CxxxxClient) GetProductNav(req domain.CxxxxProductNavRequest) ([]byte, error) {
	httpReq, err := http.NewRequest(http.MethodGet, cxxxxProductNavURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	q := httpReq.URL.Query()
	q.Set("prodCode", req.ProdCode)
	q.Set("startDate", req.StartDate)
	q.Set("endDate", req.EndDate)
	httpReq.URL.RawQuery = q.Encode()

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
