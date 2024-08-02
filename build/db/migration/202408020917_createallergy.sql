-- +goose Up
CREATE TABLE allergies (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(255) DEFAULT NULL,
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE INDEX allergy_id on allergies (id);

INSERT INTO allergies (name)
VALUES
    ('えび'),
    ('かに'),
    ('くるみ'),
    ('小麦'),
    ('そば'),
    ('卵'),
    ('乳'),
    ('落花生'),
    ('アーモンド'),
    ('あわび'),
    ('いか'),
    ('いくら'),
    ('オレンジ'),
    ('カシューナッツ'),
    ('キウイ'),
    ('牛肉'),
    ('ごま'),
    ('さけ'),
    ('さば'),
    ('大豆'),
    ('鶏肉'),
    ('バナナ'),
    ('豚肉'),
    ('マカダミアナッツ'),
    ('もも'),
    ('やまいも'),
    ('りんご'),
    ('ゼラチン');

-- +goose Down
DROP INDEX allergy_id;
DROP TABLE allergies;
