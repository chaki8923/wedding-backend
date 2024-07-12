-- +goose Up
CREATE TABLE uploads (
   id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
   comment VARCHAR(255) DEFAULT NULL,
   file_url VARCHAR(255) DEFAULT NULL,
   created_at TIMESTAMP DEFAULT NULL,
   updated_at TIMESTAMP DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE INDEX upload_id on uploads (id);

-- +goose Down
DROP INDEX invitee_id;
DROP TABLE invitees;