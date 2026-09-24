package intraday

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	di "storage/internal/domain/intraday"
)

type IntradayService interface {
	Create(ctx context.Context, p *di.CreatePayload) error
	List(ctx context.Context, filter di.ListFilter) ([]di.Intraday, error)
}

type Handler struct {
	svc IntradayService
}

func NewHandler(svc IntradayService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /storage/intraday", h.create)
	mux.HandleFunc("GET /storage/intraday/{id}", h.list)
	mux.HandleFunc("GET /storage/intraday", h.list)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req createIntradayReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid json")
		return
	}

	err := h.svc.Create(r.Context(), &di.CreatePayload{
		Ticker:    req.Ticker,
		Price:     req.Price,
		Timestamp: req.Timestamp,
	})
	if err != nil {
		if errors.Is(err, di.ErrInvalidInput) {
			writeErr(w, http.StatusBadRequest, err.Error())
			return
		}

		slog.Error("create intraday", "error", err)
		writeErr(w, http.StatusInternalServerError, "internal error")
		return
	}

	writeJSON(w, http.StatusCreated, createIntradayResp{Message: "OK"})
}

func validateAndParseFromTo(from, to string) (time.Time, time.Time, error) {
	if from == "" || to == "" {
		return time.Time{}, time.Time{}, errors.New("start_date and end_date are required")
	}

	parsedFrom, err := time.Parse(time.RFC3339, from)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid start_date: %w", err)
	}

	parsedTo, err := time.Parse(time.RFC3339, to)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid end_date: %w", err)
	}

	return parsedFrom, parsedTo, nil
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("start_date")
	to := r.URL.Query().Get("end_date")

	parsedFrom, parsedTo, err := validateAndParseFromTo(from, to)
	if err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}

	filter := di.ListFilter{
		TickerID: r.PathValue("id"),
		From:     parsedFrom,
		To:       parsedTo,
	}

	slog.Debug("list intradays", "filter", filter)

	intradaysFromDB, err := h.svc.List(r.Context(), filter)
	if err != nil {
		slog.Error("list intradays", "error", err)
		writeErr(w, http.StatusInternalServerError, "internal error")
		return
	}

	intradays := make([]getIntradaysResp, 0, len(intradaysFromDB))

	for _, i := range intradaysFromDB {
		intradays = append(intradays, getIntradaysResp{
			TickerID: i.TickerID,
			Name:     i.Ticker,
			Price:    i.Price,
			Time:     i.Timestamp.Format(time.RFC3339),
		})
	}

	writeJSON(w, http.StatusOK, intradays)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
