package main

import (
	"context"
	"fmt"
)

type ctxKey string

const (
	userIDKey  ctxKey = "userID"
	traceIDKey ctxKey = "traceID"
)

func handler(ctx context.Context) {
	userID, _ := ctx.Value(userIDKey).(string)
	traceID, _ := ctx.Value(traceIDKey).(string)
	fmt.Printf("Handling request for user %s, trace %s\n", userID, traceID)
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIDKey, "alice")
	ctx = context.WithValue(ctx, traceIDKey, "abc-123")

	handler(ctx)
}
