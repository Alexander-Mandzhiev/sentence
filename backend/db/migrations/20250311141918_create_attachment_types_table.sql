-- +goose Up
-- +goose StatementBegin
CREATE TABLE attachment_types (
                                  id SERIAL PRIMARY KEY,
                                  name VARCHAR(50) NOT NULL UNIQUE,
                                  description TEXT,
                                  is_active BOOLEAN DEFAULT TRUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE attachment_types;
-- +goose StatementEnd