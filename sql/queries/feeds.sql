-- name: CreateFeed :exec
INSERT INTO feeds(id,created_at,updated_at,name,url,user_id)
VALUES(?,?,?,?,?,?);

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByID :one
SELECT * FROM feeds WHERE id=?;