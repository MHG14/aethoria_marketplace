-- name: CreateListing :one
INSERT INTO listings (item_id, seller_id, price, status, updated_at)
VALUES ($1, $2, $3, 'active', NOW())
    RETURNING *;

-- name: GetListing :one
SELECT * FROM listings WHERE id = $1;

-- name: GetListingForUpdate :one
SELECT * FROM listings WHERE id = $1 FOR UPDATE;

-- name: UpdateListingStatus :one
UPDATE listings
SET status = $2, buyer_id = $3, updated_at = NOW()
WHERE id = $1
    RETURNING *;

-- name: ListActiveListings :many
SELECT * FROM listings WHERE status = 'active' ORDER BY created_at DESC;