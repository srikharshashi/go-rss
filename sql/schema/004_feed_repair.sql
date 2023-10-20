-- +goose Up
ALTER TABLE feeds
MODIFY url VARCHAR(128);


-- +goose Down

DROP TABLE IF EXISTS feeds;