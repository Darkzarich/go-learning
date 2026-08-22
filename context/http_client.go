package main

import (
    "context"
    "fmt"
    "io"
    "net/http"
    "time"
)

func fetchURL(ctx context.Context, url string) error {
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return err
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    fmt.Printf("Fetched %d bytes\n", len(body))
    return nil
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    if err := fetchURL(ctx, "https://example.com"); err != nil {
        fmt.Println("Error:", err)
    }
}