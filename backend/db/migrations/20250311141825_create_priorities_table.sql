-- +goose Up
-- +goose StatementBegin
CREATE TABLE priorities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE priorities;
-- +goose StatementEnd