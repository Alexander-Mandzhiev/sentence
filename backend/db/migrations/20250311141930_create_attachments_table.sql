-- +goose Up
-- +goose StatementBegin
CREATE TABLE attachments (
                             id SERIAL PRIMARY KEY,
                             attachment_type_id INT NOT NULL REFERENCES attachment_types(id) ON DELETE RESTRICT,
                             file_name VARCHAR(255) NOT NULL,
                             file_path TEXT NOT NULL,
                             mime_type VARCHAR(100),
                             file_size BIGINT,
                             uploaded_by INT REFERENCES users(id) ON DELETE SET NULL,
                             created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE attachments;
-- +goose StatementEnd