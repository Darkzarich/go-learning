package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrAlreadyPaid   = errors.New("order already paid")
)

type ValidationError struct {
	Field   string
	Message string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field %q: %s", ve.Field, ve.Message)
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
	Field   string `json:"field,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		// Fallback in case encoding fails
		log.Printf("writeJSON: encode error: %v", err)
	}
}

func writeError(w http.ResponseWriter, err error) {
	// Errors carrying their own fields via errors.As
	var ve *ValidationError
	if errors.As(err, &ve) {
		writeJSON(w, http.StatusBadRequest, apiError{
			Code:    "validation",
			Field:   ve.Field,
			Message: ve.Message,
		})
		return
	}

	// Sentinel errors via errors.Is
	switch {
	case errors.Is(err, ErrOrderNotFound):
		writeJSON(w, http.StatusNotFound, apiError{Code: "not_found"})
	case errors.Is(err, ErrAlreadyPaid):
		writeJSON(w, http.StatusConflict, apiError{Code: "already_paid"})
	default:
		log.Printf("unhandled error: %v", err)
		writeJSON(w, http.StatusInternalServerError, apiError{Code: "internal"})
	}
}

// payOrder simulates a payment operation that may fail.
func payOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		writeError(w, &ValidationError{
			Field:   "id",
			Message: "order id is required",
		})
		return
	}

	// Imagine you look up the order and find it does not exist.
	if orderID == "999" {
		writeError(w, ErrOrderNotFound)
		return
	}

	// Imagine you find the order is already paid.
	if orderID == "111" {
		writeError(w, ErrAlreadyPaid)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "paid"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/pay", payOrder)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// Test running and sending requests:
//
//   $ go run main.go
//   $ curl -s http://localhost:8080/pay?id=999
//   $ curl -s http://localhost:8080/pay?id=111
//   $ curl -s http://localhost:8080/pay?id=123
