-- +goose Up
-- +goose StatementBegin
INSERT INTO implementors (name, email, is_active) VALUES ('Нет исполнителя', 'no_executor@mail.com', true);
-- +goose StatementEnd