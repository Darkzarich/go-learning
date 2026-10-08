package intraday

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	di "api/internal/domain/intraday"
)

type IntradayService interface {
	ListIntradays(ctx context.Context, filter di.ListFilter) ([]di.Intraday, error)
}

type Handler struct {
	svc IntradayService
}

func NewHandler(svc IntradayService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api", h.list)
	mux.HandleFunc("GET /api/{id}", h.list)
}

func validateAndParseFromTo(from, to string) (time.Time, time.Time, error) {
	var parsedFrom, parsedTo time.Time
	var err error

	if from != "" {
		parsedFrom, err = time.Parse(time.RFC3339, from)
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("invalid start_date: %w", err)
		}
	}

	if to != "" {
		parsedTo, err = time.Parse(time.RFC3339, to)
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("invalid end_date: %w", err)
		}
	}

	if !parsedFrom.IsZero() && !parsedTo.IsZero() && parsedFrom.After(parsedTo) {
		return time.Time{}, time.Time{}, errors.New("start_date must be before end_date")
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

	var tickerID int64
	if s := r.PathValue("id"); s != "" {
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil || id <= 0 {
			writeErr(w, http.StatusBadRequest, "invalid ticker id")
			return
		}
		tickerID = id
	}

	filter := di.ListFilter{
		TickerID: tickerID,
		From:     parsedFrom,
		To:       parsedTo,
	}

	slog.Debug("list intradays", "filter", filter)

	intradaysFromDB, err := h.svc.ListIntradays(r.Context(), filter)
	if err != nil {
		if errors.Is(err, di.ErrRejected) {
			writeErr(w, http.StatusBadRequest, err.Error())
			return
		}

		slog.Error("list intradays", "error", err)
		writeErr(w, http.StatusInternalServerError, "internal error")
		return
	}

	intradays := make([]getIntradaysResp, 0, len(intradaysFromDB))

	for _, i := range intradaysFromDB {
		intradays = append(intradays, getIntradaysResp{
			TickerID: i.TickerID,
			Name:     i.Name,
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
