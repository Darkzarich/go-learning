# Users HTTP API service with SQLite database.

## Usage

### Run API server

```bash
make run-api
```

### Run cron jobs

```bash
make run-cron
```

### Running both API and cron jobs

```bash
make -j2 run-api run-cron
```

### Create a user

```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'
```

### Get all users

```bash
curl http://localhost:3000/users
```

### Get a user

```bash
curl http://localhost:3000/users/1
```

### Update a user

```bash
curl -X PUT http://localhost:3000/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'
```

### Delete a user

```bash
curl -X DELETE http://localhost:3000/users/1
```

### Set last_login to 30 days ago

This script will update user with ID 1 to have last_login set to 30 days ago.

The cron job will remove users who haven't logged in since 30 days every 30 seconds so this script is useful for testing.

```bash
go run scripts/make_inactive_last_login.go
```
