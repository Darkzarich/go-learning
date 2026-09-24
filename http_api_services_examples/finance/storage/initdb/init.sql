CREATE TABLE
  IF NOT EXISTS tickers (
    id BIGSERIAL PRIMARY KEY,
    ticker VARCHAR(20) NOT NULL UNIQUE
  );

CREATE TABLE
  IF NOT EXISTS intradays (
    id BIGSERIAL PRIMARY KEY,
    ticker_id BIGINT NOT NULL REFERENCES tickers (id) ON DELETE CASCADE,
    price NUMERIC(20, 8) NOT NULL,
    timestamp TIMESTAMP NOT NULL
  );

CREATE INDEX IF NOT EXISTS idx_intraday_ticker_ts ON intradays (ticker_id, timestamp DESC);