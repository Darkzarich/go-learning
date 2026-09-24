package intraday

import (
	"context"
	"fmt"
	"log/slog"
	"storage/internal/domain/intraday"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	pool          *pgxpool.Pool
	mu            sync.RWMutex
	tickerIDCache map[string]int64
}

var _ intraday.Repository = (*PostgresRepo)(nil)

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{pool: pool, tickerIDCache: make(map[string]int64)}
}

func (r *PostgresRepo) cachedTickerID(ticker string) (int64, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	id, ok := r.tickerIDCache[ticker]

	return id, ok
}

func (r *PostgresRepo) putTickerID(ticker string, id int64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tickerIDCache[ticker] = id
}

func (r *PostgresRepo) Create(ctx context.Context, p *intraday.CreatePayload) error {
	tickerID, cached := r.cachedTickerID(p.Ticker)

	if cached {
		_, err := r.pool.Exec(ctx, `
			INSERT INTO intradays (ticker_id, price, timestamp)
			VALUES ($1, $2, $3)
		`, tickerID, p.Price, p.Timestamp)
		if err != nil {
			return fmt.Errorf("repo create intraday for ticker %q: %w", p.Ticker, err)
		}
	} else {
		err := r.pool.QueryRow(ctx, `
	WITH t AS (
		INSERT INTO tickers (ticker)
		VALUES ($1)
		ON CONFLICT (ticker) DO UPDATE SET ticker = EXCLUDED.ticker
		RETURNING id
	)
	INSERT INTO intradays (ticker_id, price, timestamp)
	SELECT id, $2, $3 FROM t
	RETURNING ticker_id
`, p.Ticker, p.Price, p.Timestamp).Scan(&tickerID)
		if err != nil {
			return fmt.Errorf("repo create intraday for ticker %q: %w", p.Ticker, err)
		}

		r.putTickerID(p.Ticker, tickerID)
	}

	slog.Debug("Create", "ticker", p.Ticker, "ticker_id", tickerID, "price", p.Price, "timestamp", p.Timestamp)

	return nil
}

func (r *PostgresRepo) List(ctx context.Context, filter intraday.ListFilter) ([]intraday.Intraday, error) {
	baseQuery := `SELECT i.id as id, t.id as ticker_id, ticker, price, timestamp FROM intradays i
	JOIN tickers t ON t.id = i.ticker_id
	WHERE timestamp >= $1 AND timestamp <= $2 `

	// 3rd placeholder because from and to take 1 and 2 respectively
	placeholderIdx := 3
	var args []any = []any{filter.From, filter.To}

	if filter.TickerID != "" {
		intTickerID, err := strconv.Atoi(filter.TickerID)
		if err != nil {
			return nil, fmt.Errorf("repo list intradays %+v: %w", filter, err)
		}
		baseQuery += fmt.Sprintf(`AND ticker_id = $%d `, placeholderIdx)
		args = append(args, intTickerID)
		// placeholderIdx++
	}

	finalQuery := baseQuery + `ORDER BY ticker, timestamp;`

	slog.Debug("List", "query", finalQuery, "args", args)

	rows, err := r.pool.Query(ctx, finalQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("repo list intradays %+v: %w", filter, err)
	}
	defer rows.Close()

	var intradays []intraday.Intraday

	for rows.Next() {
		var i intraday.Intraday
		if err := rows.Scan(&i.ID, &i.TickerID, &i.Ticker, &i.Price, &i.Timestamp); err != nil {
			return nil, fmt.Errorf("repo list intradays %+v: %w", filter, err)
		}
		intradays = append(intradays, i)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo list intradays %+v: %w", filter, err)
	}

	return intradays, nil
}
