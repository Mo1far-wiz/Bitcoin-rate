# bitcoin-rate

This is simple Backend application that takes BTC rate from public API and sent out Emails.

## Technologies
- Gin
- Postgres
- Docker-compose

## To run

You ned to create ```.env``` file with such keys:
```.env
SMTP_HOST=***
SMTP_PORT=***
SMTP_USER=***
SMTP_PASS=***

DB_USER=user
DB_PASSWORD=password
DB_NAME=btc-rate
DB_HOST=localhost
DB_PORT=5432
```

Than run following commands:
```cmd
docker-compose up
```
```cmd
go run ./cmd/web
```

### Tests

There are no UNIT or Integration tests yet, but i have added Postman Collection.
