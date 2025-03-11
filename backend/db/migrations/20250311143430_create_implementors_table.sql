-- +goose Up
-- +goose StatementBegin
CREATE TABLE implementors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) UNIQUE,
    is_active BOOLEAN DEFAULT TRUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE implementors;
-- +goose StatementEnd