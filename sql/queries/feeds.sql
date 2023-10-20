-- name: CreateFeed :exec
INSERT INTO feeds(id,created_at,updated_at,name,url,user_id)
VALUES(?,?,?,?,?,?);

-- name: GetFeeds :many
select * from feeds;