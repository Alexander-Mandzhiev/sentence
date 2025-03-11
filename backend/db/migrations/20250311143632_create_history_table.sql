-- +goose Up
-- +goose StatementBegin
CREATE TABLE history (
    id SERIAL PRIMARY KEY,
    sentence_id BIGINT NOT NULL REFERENCES sentences(id) ON DELETE CASCADE,
    status_id INT NOT NULL REFERENCES statuses(id) ON DELETE SET NULL,
    changed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    reason TEXT,
    details TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE history;
-- +goose StatementEnd