-- +goose Up


ALTER TABLE users
ALTER api_key SET DEFAULT CONCAT(MD5(RAND()),MD5(RAND()));


ALTER TABLE users
ADD CONSTRAINT unique_key UNIQUE(api_key);

ALTER TABLE users
MODIFY COLUMN api_key VARCHAR(100) NOT NULL;


-- +goose Down

ALTER TABLE users
ALTER COLUMN api_key DROP DEFAULT;

ALTER TABLE users
DROP CONSTRAINT unique_key;