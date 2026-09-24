Notes to self:

```sh
# build and run everything
docker compose up --build

# init database
cat initdb/init.sql | docker exec -i some-postgres psql -U postgres -d postgres

# export db connection string to run storage service
export DATABASE_URL=postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable
```

API of the storage service:

```
GET localhost:8081/storage/intraday?start_date=2026-08-20T12:30:35Z&end_date=2026-09-25T12:30:35Z

GET localhost:8081/storage/intraday/1?start_date=2026-08-20T12:30:35Z&end_date=2026-09-25T12:30:35Z

POST localhost:8081/storage/intraday

{
  "ticker": "BTC",
  "price": 7172162,
  "timestamp": "2026-09-24T00:48:01Z"
}
```
