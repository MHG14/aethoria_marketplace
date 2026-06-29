-- name: CreateBid :one
INSERT INTO bids (auction_id, guild_id, amount)
VALUES ($1, $2, $3)
    RETURNING *;

-- name: GetBid :one
SELECT * FROM bids WHERE id = $1;

-- name: CancelBid :one
UPDATE bids SET is_cancelled = TRUE WHERE id = $1 RETURNING *;

-- name: ListBidsByAuction :many
SELECT * FROM bids
WHERE auction_id = $1 AND is_cancelled = FALSE
ORDER BY amount DESC;

-- name: ListActiveBidsByGuildAndAuction :many
SELECT * FROM bids
WHERE auction_id = $1 AND guild_id = $2 AND is_cancelled = FALSE
ORDER BY amount DESC;