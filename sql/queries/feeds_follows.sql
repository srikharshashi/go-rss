-- name: CreateFeedFollows :exec
INSERT INTO feeds_follows(id,created_at,updated_at,user_id,feed_id)
VALUES (?,?,?,?,?);

-- name: GetFeedFollowsByUser :many
SELECT * from feeds_follows where user_id=?;

-- name: DeleteFeedFollowByUserandId :exec
DELETE FROM feeds_follows where id=? AND user_id=?;