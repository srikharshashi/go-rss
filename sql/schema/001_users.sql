-- +goose Up

CREATE TABLE users(
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(36) NOT NULL,
    api_key VARCHAR(64) NOT NULL 
);

-- +goose Down

DROP TABLE IF EXISTS users;
