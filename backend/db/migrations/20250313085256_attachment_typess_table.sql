-- +goose Up
-- +goose StatementBegin
INSERT INTO attachment_types (name, description) VALUES
    ('BeforePhoto', 'Фотография объекта или ситуации до выполнения работ.'),
    ('AfterPhoto', 'Фотография объекта или ситуации после выполнения работ.');
-- +goose StatementEnd
