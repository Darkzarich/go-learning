package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	di "api/internal/domain/intraday"
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

func (c *Client) FetchIntradays(ctx context.Context, filter di.ListFilter) ([]di.Intraday, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/storage/intraday", nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	q := req.URL.Query()
	q.Add("start_date", filter.From.Format(time.RFC3339))
	q.Add("end_date", filter.To.Format(time.RFC3339))

	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var respBody getIntradaysResp
		if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			return nil, fmt.Errorf("decode response: %w", err)
		}

		intradays := make([]di.Intraday, len(respBody))
		for idx, i := range respBody {
			intradays[idx] = di.Intraday{
				TickerID:  i.TickerID,
				Name:      i.Name,
				Price:     i.Price,
				Timestamp: i.Time,
			}
		}
		return intradays, nil
	}

	msg, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return nil, fmt.Errorf("%w: status %d: %s", di.ErrRejected, resp.StatusCode, bytes.TrimSpace(msg))
	}

	return nil, fmt.Errorf("storage internal error status %d: %s", resp.StatusCode, bytes.TrimSpace(msg))
}
