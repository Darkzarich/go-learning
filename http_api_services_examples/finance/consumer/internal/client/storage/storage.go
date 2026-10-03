package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	di "consumer/internal/domain/intraday"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient(httpClient *http.Client, url string) *Client {
	return &Client{
		httpClient: httpClient,
		baseURL:    url,
	}
}

func (c *Client) SaveIntraday(ctx context.Context, intraday di.Intraday) error {
	data, err := json.Marshal(intraday)
	if err != nil {
		return fmt.Errorf("marshal intraday: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/storage/intraday", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("storage status %d: %s", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()

	return nil
}
