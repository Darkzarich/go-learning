package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	di "consumer/internal/domain/intraday"
)

type ClientImpl struct {
	baseURL string
}

func NewClient(url string) *ClientImpl {
	return &ClientImpl{baseURL: url}
}

func (c *ClientImpl) SaveIntraday(ctx context.Context, intraday di.Intraday) error {
	data, err := json.Marshal(intraday)
	if err != nil {
		return fmt.Errorf("marshal intraday: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/storage/intraday", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return fmt.Errorf("storage status %d: %s", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()

	return nil
}
