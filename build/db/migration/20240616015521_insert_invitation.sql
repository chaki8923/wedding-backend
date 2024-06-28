-- +goose Up
CREATE TABLE invitations (
   id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
   title VARCHAR(255) DEFAULT NULL,
   event_date VARCHAR(255) DEFAULT NULL,
   place VARCHAR(255) DEFAULT NULL,
   comment VARCHAR(255) DEFAULT NULL,
   file_url VARCHAR(255) DEFAULT NULL,
   user_id INT UNSIGNED NOT NULL,
   created_at TIMESTAMP DEFAULT NULL,
   updated_at TIMESTAMP DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE INDEX invitation_id on invitations (id);

-- +goose Down
DROP INDEX invitation_id;
DROP TABLE invitations;