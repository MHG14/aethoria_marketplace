#!/bin/bash
set -e

DB_URL="${DATABASE_URL:-postgres://postgres:postgres@localhost:5432/aethoria?sslmode=disable}"

echo "Running migrations..."
for f in internal/infrastructure/persistence/postgres/*.sql; do
    echo "Applying $f..."
    psql "$DB_URL" -f "$f"
done

echo "Seeding data..."
psql "$DB_URL" -f scripts/seed.sql

echo "Done."