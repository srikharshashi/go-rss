-- +goose Up
CREATE TABLE feeds(
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(32) NOT NULL,
    url VARCHAR(32) NOT NULL,
    user_id VARCHAR(36) references users(id) ON DELETE CASCADE 
);



-- +goose Down

DROP TABLE IF EXISTS feeds;