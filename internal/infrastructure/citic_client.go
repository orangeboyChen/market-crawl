package infrastructure

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"market-crawl/internal/domain"
)

const citicProductNavURL = "https://wechat.citic-wealth.com/cms.product/api/custom/productInfo/getTAProductNav"

// CiticClient implements domain.CiticProductNavRepository by calling the CITIC Wealth API.
type CiticClient struct {
	httpClient *http.Client
}

// NewCiticClient creates a new CiticClient with a default timeout.
func NewCiticClient() *CiticClient {
	return &CiticClient{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetProductNav sends a GET request to the CITIC Wealth API and returns the raw response bytes.
func (c *CiticClient) GetProductNav(req domain.CiticProductNavRequest) ([]byte, error) {
	httpReq, err := http.NewRequest(http.MethodGet, citicProductNavURL, nil)
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
