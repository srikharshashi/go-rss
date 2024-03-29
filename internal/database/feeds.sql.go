// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: feeds.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createFeed = `-- name: CreateFeed :exec
INSERT INTO feeds(id,created_at,updated_at,name,url,user_id)
VALUES(?,?,?,?,?,?)
`

type CreateFeedParams struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       sql.NullString
	UserID    sql.NullString
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) error {
	_, err := q.db.ExecContext(ctx, createFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	return err
}

const getFeedByID = `-- name: GetFeedByID :one
SELECT id, created_at, updated_at, name, user_id, url FROM feeds WHERE id=?
`

func (q *Queries) GetFeedByID(ctx context.Context, id string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeedByID, id)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.UserID,
		&i.Url,
	)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
SELECT id, created_at, updated_at, name, user_id, url FROM feeds
`

func (q *Queries) GetFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.UserID,
			&i.Url,
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
