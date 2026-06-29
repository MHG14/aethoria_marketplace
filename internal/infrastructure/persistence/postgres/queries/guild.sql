-- name: GetGuild :one
SELECT * FROM guilds WHERE id = $1;

-- name: GetGuildForUpdate :one
SELECT * FROM guilds WHERE id = $1 FOR UPDATE;

-- name: UpdateGuildWallet :one
UPDATE guilds
SET total_money    = $2,
    reserved_money = $3,
    daily_spent    = $4
WHERE id = $1
    RETURNING *;

-- name: ResetDailySpent :exec
UPDATE guilds SET daily_spent = 0;