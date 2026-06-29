# Aethoria Marketplace
A mystical bazaar where seekers buy and bid on legendary artifacts of Aethoria.

## Quick start

```bash
docker compose up --build
```

This starts postgres (with schema + seed data) and the app on port 8080.

## Run without Docker

```bash
# start postgres however you prefer, then:
export DATABASE_URL=postgres://postgres:postgres@localhost:5432/aethoria?sslmode=disable
go run ./cmd/server
```

## Test the API

Follow `docs/test_flow.md` for a complete step-by-step test procedure
covering listings, auctions, bidding, wallet operations, and edge cases.

