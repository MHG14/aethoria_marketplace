-- name: CreateTrade :one
INSERT INTO trades (item_id, seller_id, buyer_id, price, type)
VALUES ($1, $2, $3, $4, $5)
    RETURNING *;

-- name: ListTradesByGuild :many
SELECT * FROM trades
WHERE seller_id = $1 OR buyer_id = $1
ORDER BY created_at DESC;