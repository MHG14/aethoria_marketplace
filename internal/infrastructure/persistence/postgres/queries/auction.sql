-- name: CreateAuction :one
INSERT INTO auctions (item_id, seller_id, starting_price, end_time, original_end_time, status)
VALUES ($1, $2, $3, $4, $4, 'active')
    RETURNING *;

-- name: GetAuction :one
SELECT * FROM auctions WHERE id = $1;

-- name: GetAuctionForUpdate :one
SELECT * FROM auctions WHERE id = $1 FOR UPDATE;

-- name: GetActiveAuctionByItem :one
SELECT * FROM auctions WHERE item_id = $1 AND status = 'active' LIMIT 1;

-- name: UpdateAuctionBid :one
UPDATE auctions
SET highest_bid       = $2,
    highest_bidder_id = $3,
    end_time          = $4
WHERE id = $1
    RETURNING *;

-- name: UpdateAuctionStatus :one
UPDATE auctions SET status = $2 WHERE id = $1 RETURNING *;

-- name: ListExpiredAuctions :many
SELECT * FROM auctions
WHERE status = 'active' AND end_time <= NOW();

-- name: ListActiveAuctions :many
SELECT * FROM auctions WHERE status = 'active' ORDER BY created_at DESC;