-- name: CreateFeedsFollows :one
INSERT INTO feeds_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedsFollows :many
SELECT * FROM feeds_follows WHERE user_id = $1;

-- name: DeleteFeedsFollows :exec
DELETE FROM feeds_follows WHERE id = $1 AND user_id = $2;


