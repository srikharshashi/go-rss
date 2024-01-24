// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: feeds_follows.sql

package database

import (
	"context"
	"time"
)

const createFeedFollows = `-- name: CreateFeedFollows :exec
INSERT INTO feeds_follows(id,created_at,updated_at,user_id,feed_id)
VALUES (?,?,?,?,?)
`

type CreateFeedFollowsParams struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
	FeedID    string
}

func (q *Queries) CreateFeedFollows(ctx context.Context, arg CreateFeedFollowsParams) error {
	_, err := q.db.ExecContext(ctx, createFeedFollows,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	return err
}

const deleteFeedFollowByUserandId = `-- name: DeleteFeedFollowByUserandId :exec
DELETE FROM feeds_follows where id=? AND user_id=?
`

type DeleteFeedFollowByUserandIdParams struct {
	ID     string
	UserID string
}

func (q *Queries) DeleteFeedFollowByUserandId(ctx context.Context, arg DeleteFeedFollowByUserandIdParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollowByUserandId, arg.ID, arg.UserID)
	return err
}

const getFeedFollowsByUser = `-- name: GetFeedFollowsByUser :many
SELECT id, created_at, updated_at, user_id, feed_id from feeds_follows where user_id=?
`

func (q *Queries) GetFeedFollowsByUser(ctx context.Context, userID string) ([]FeedsFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedsFollow
	for rows.Next() {
		var i FeedsFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
