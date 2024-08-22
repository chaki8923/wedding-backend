-- +goose Up
ALTER TABLE invitees ADD COLUMN join_flag BOOLEAN NOT NULL DEFAULT TRUE;

-- +goose Down
ALTER TABLE invitees DROP COLUMN join_flag;