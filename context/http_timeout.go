package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ctxKey string

const requestIDKey ctxKey = "requestID"

// Middleware adds a request ID to the context and sets a timeout
func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = "generated-id"
		}
		ctx := context.WithValue(r.Context(), requestIDKey, reqID)
		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqID, _ := ctx.Value(requestIDKey).(string)

	// Simulate a long-running operation that respects context
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "Request %s completed\n", reqID)
	case <-ctx.Done():
		http.Error(w, "request canceled or timed out", http.StatusGatewayTimeout)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: withRequestID(mux),
	}
	server.ListenAndServe()
}
