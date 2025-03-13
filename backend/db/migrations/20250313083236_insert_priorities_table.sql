-- +goose Up
-- +goose StatementBegin
INSERT INTO priorities (name, description) VALUES
    ('Низкий', 'Приоритет низкого уровня.'),
    ('Средний', 'Приоритет среднего уровня.'),
    ('Высокий', 'Приоритет высокого уровня.');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM priorities WHERE name IN ('Низкий', 'Средний', 'Высокий');
-- +goose StatementEnd