package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	data, err := json.Marshal(newSaveIntradayReq(intraday))
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
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		_, _ = io.Copy(io.Discard, resp.Body)
		return nil
	}

	msg, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return fmt.Errorf("%w: status %d: %s", di.ErrRejected, resp.StatusCode, bytes.TrimSpace(msg))
	}

	return fmt.Errorf("storage internal error status %d: %s", resp.StatusCode, bytes.TrimSpace(msg))
}
