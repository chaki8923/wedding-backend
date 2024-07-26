ALTER TABLE invitations ADD COLUMN uu_id CHAR(36) AFTER id;
UPDATE invitations SET uu_id = (SELECT UUID());
ALTER TABLE invitations ADD UNIQUE (uu_id);