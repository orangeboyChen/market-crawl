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

const icbcNetValueURL = "https://papi.icbc.com.cn/finance/deposit/consignment/getNetValueList"

// ICBCClient implements domain.NetValueRepository by calling the ICBC external API.
type ICBCClient struct {
	httpClient *http.Client
}

// NewICBCClient creates a new ICBCClient with a default timeout.
func NewICBCClient() *ICBCClient {
	return &ICBCClient{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetNetValueList sends a POST request to the ICBC API and returns the raw response bytes.
func (c *ICBCClient) GetNetValueList(req domain.NetValueRequest) ([]byte, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, icbcNetValueURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

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
