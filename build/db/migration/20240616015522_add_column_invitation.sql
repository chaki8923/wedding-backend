-- +goose Up
ALTER TABLE invitations ADD COLUMN uuid CHAR(36) AFTER id;
UPDATE invitations SET uuid = (SELECT UUID());
ALTER TABLE invitations ADD UNIQUE (uuid);

-- +goose Down
ALTER TABLE invitations DROP COLUMN uuid;