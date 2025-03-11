-- +goose Up
-- +goose StatementBegin
CREATE TABLE priorities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    level INT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE priorities;
-- +goose StatementEnd