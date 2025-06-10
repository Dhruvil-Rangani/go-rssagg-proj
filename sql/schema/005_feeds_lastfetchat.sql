-- +goose UP
CREATE EXTENSION IF NOT EXISTS pgcrypto;

ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose DOWN
ALTER TABLE feeds DROP COLUMN last_fetched_at;