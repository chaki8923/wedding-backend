ALTER TABLE invitees ADD COLUMN uuid CHAR(36) AFTER id;
UPDATE invitees SET uuid = (SELECT UUID());
ALTER TABLE invitees ADD UNIQUE (uuid);