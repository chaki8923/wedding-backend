-- +goose Up
CREATE TABLE IF NOT EXISTS upload_images (
   id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
   comment VARCHAR(255) DEFAULT NULL,
   file_url VARCHAR(255) DEFAULT NULL,
   created_at TIMESTAMP DEFAULT NULL,
   updated_at TIMESTAMP DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE INDEX upload_images_id ON upload_images (id);

-- +goose Down
DROP INDEX IF EXISTS upload_images_id ON upload_images;
DROP TABLE IF EXISTS upload_images;