package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type StorageClient struct {
	baseURL string
}

func NewStorageClient(url string) IStorageClient {
	return &StorageClient{baseURL: url}
}

func (c *StorageClient) SaveIntraday(ctx context.Context, intraday Intraday) error {
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

	defer resp.Body.Close()

	return nil
}
