-- name: CreateWalletTransaction :one
INSERT INTO wallet_transactions (guild_id, type, amount, ref_type, ref_id)
VALUES ($1, $2, $3, $4, $5)
    RETURNING *;

-- name: ListWalletTransactions :many
SELECT * FROM wallet_transactions
WHERE guild_id = $1
ORDER BY created_at DESC;

-- name: GetDailySpent :one
SELECT COALESCE(SUM(amount), 0)::BIGINT FROM wallet_transactions
WHERE guild_id = $1
  AND type = 'reserve'
  AND created_at >= NOW()::DATE;