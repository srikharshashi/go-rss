-- +goose Up
ALTER TABLE users 
ADD COLUMN api_key VARCHAR(100);

UPDATE users SET api_key= CONCAT(MD5(RAND()),MD5(RAND()));


-- +goose Down
ALTER TABLE USERS DROP COLUMN api_key;